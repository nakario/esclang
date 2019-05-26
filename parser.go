package esc

import (
	"errors"
	"io"
)

type Parser struct {
	l *Lexer
	labels map[rune]int
	jumpTo map[rune]struct{}
	instructions []Instruction
}

func NewParser(l *Lexer) *Parser {
	Logger.Println("NewParser")
	return &Parser{l, make(map[rune]int), make(map[rune]struct{}), []Instruction{}}
}

func (p *Parser) Parse() error {
	Logger.Println("*Parser.Parse")
	var t Token
	var err error
	var labeling = false
	var jumping = false
	var bgInstruction func(rune) Instruction
	var fgInstruction func(rune) Instruction
	Logger.Println("Start reading tokens")
	for t, err = p.l.NextToken(); err == nil; t, err = p.l.NextToken() {
		Logger.Println("Reading token", t)
		switch token := t.(type) {
		case NormalToken:
			Logger.Println("normal token")
			if labeling {
				_, ok := p.labels[token.Char()]
				if ok {
					return errors.New("duplicated label")
				}
				p.labels[token.Char()] = len(p.instructions) - 1
			} else if bgInstruction != nil {
				p.instructions = append(p.instructions, bgInstruction(token.Char()))
			}
			if jumping {
				p.jumpTo[token.Char()] = struct{}{}
			}
			if fgInstruction != nil {
				p.instructions = append(p.instructions, fgInstruction(token.Char()))
			}
		case _SGRToken:
			Logger.Println("SGR token")
			labeling = false
			jumping = false
			switch token.param {
			case 30:
				fgInstruction = copyInstruction
			case 31:
				fgInstruction = incrementInstruction
			case 32:
				fgInstruction = inputInstruction
			case 33:
				fgInstruction = rotateLeftInstruction
			case 34:
				fgInstruction = rotateRightInstruction
			case 35:
				fgInstruction = outputInstructioin
			case 36:
				fgInstruction = decrementInstruction
			case 37:
				fgInstruction = swapInstruction
			case 39:
				fgInstruction = nopInstruction
			case 40:
				bgInstruction = callInstruction
			case 41:
				bgInstruction = incrementPtrInstruction
			case 42:
				jumping = true
				bgInstruction = jumpIfZeroInstruction
			case 43:
				bgInstruction = incrementPrimalPtrInstruction
			case 44:
				bgInstruction = decrementPrimalPtrInstruction
			case 45:
				labeling = true
			case 46:
				bgInstruction = decrementPtrInstruction
			case 47:
				bgInstruction = exitInstruction
			case 49:
				bgInstruction = nopInstruction
			default:
				return errors.New("unexpected SGR parameter")
			}
		default:
			err = errors.New("unexpected token")
		}
	}
	if err == io.EOF {
		err = nil
	}
	Logger.Println(p.labels)
	Logger.Println(p.jumpTo)
	for k, _ := range p.jumpTo {
		if _, ok := p.labels[k]; !ok {
			return errors.New("unexpected jump destination")
		}
	}
	return err
}
