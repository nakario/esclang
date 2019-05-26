package esc

import (
	"github.com/pkg/errors"
	"io"
	"strconv"
)

type Lexer struct {
	source    []rune
	size      int
	nextIndex int
}

func NewLexer(source string) *Lexer {
	Logger.Println("NewLexer")
	s := []rune(source)
	return &Lexer{s, len(s), 0}
}

func (l *Lexer) NextToken() (Token, error) {
	if l.nextIndex == l.size {
		return nil, io.EOF
	}
	if l.isNextTokenControlSequence() {
		return l.nextControlSequenceToken()
	}
	nt := normalToken{
		baseToken{l.nextIndex, l.nextIndex + 1},
		l.source[l.nextIndex],
	}
	l.nextIndex++
	return nt, nil
}

func (l *Lexer) isNextTokenControlSequence() bool {
	return l.source[l.nextIndex] == '\033'
}

var InvalidControlSequenceError = errors.New("failed to parse a control sequence")

// TODO: reduce the number of checks of `l.nextIndex == l.size`
func (l *Lexer) nextControlSequenceToken() (ControlSequenceToken, error) {
	if err := l.validateNextControlSequenceToken(); err != nil {
		return nil, InvalidControlSequenceError
	}
	beginAt := l.nextIndex
	l.nextIndex++
	if l.nextIndex == l.size {
		return nil, InvalidControlSequenceError
	}

	if l.source[l.nextIndex] != '[' {
		return nil, InvalidControlSequenceError
	}
	l.nextIndex++
	if l.nextIndex == l.size {
		return nil, InvalidControlSequenceError
	}

	j := l.parameterEndAt(l.nextIndex)
	parameter := l.source[l.nextIndex:j]
	l.nextIndex = j
	if l.nextIndex == l.size {
		return nil, InvalidControlSequenceError
	}

	j = l.intermediateEndAt(l.nextIndex)
	intermediate := l.source[l.nextIndex:j]
	l.nextIndex = j
	if l.nextIndex == l.size {
		return nil, InvalidControlSequenceError
	}

	final := l.source[l.nextIndex]
	l.nextIndex++
	if l.nextIndex == l.size {
		return nil, InvalidControlSequenceError
	}

	cst := controlSequenceToken{
		baseToken{beginAt, l.nextIndex},
		parameter,
		intermediate,
		final,
	}

	if final == 'm' {
		paramValue := 0
		if len(parameter) != 0 {
			var err error
			paramValue, err = strconv.Atoi(string(parameter))
			if err != nil {
				return nil, InvalidControlSequenceError
			}
		}
		return _SGRToken{
			cst,
			paramValue,
		}, nil
	}

	return cst, nil
}

// TODO: Implement validation
func (l *Lexer) validateNextControlSequenceToken() error {
	return nil
}

func (l *Lexer) parameterEndAt(beginAt int) int {
	i := beginAt
	for i < l.size && 0x30 <= l.source[i] && l.source[i] <= 0x3f {
		i++
	}
	return i
}

func (l *Lexer) intermediateEndAt(beginAt int) int {
	i := beginAt
	for i < l.size && 0x20 <= l.source[i] && l.source[i] <= 0x2f {
		i++
	}
	return i
}
