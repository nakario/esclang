package esc

type Token interface {
	BeginAt() int
	EndAt() int
}

type baseToken struct {
	beginAt int
	endAt int
}

func (t baseToken) BeginAt() int {
	return t.beginAt
}

func (t baseToken) EndAt() int {
	return t.endAt
}

type NormalToken interface {
	Token
	Char() rune
}

type normalToken struct {
	baseToken
	char rune
}

func (t normalToken) Char() rune {
	return t.char
}

type ControlSequenceToken interface {
	Token
	ParameterRunes() []rune
	IntermediateRunes() []rune
	FinalRune() rune
}

type controlSequenceToken struct {
	baseToken
	parameterRunes []rune
	intermediateRunes []rune
	finalRune rune
}

func (t controlSequenceToken) ParameterRunes() []rune {
	return t.parameterRunes[:]
}

func (t controlSequenceToken) IntermediateRunes() []rune {
	return t.intermediateRunes[:]
}

func (t controlSequenceToken) FinalRune() rune {
	return t.finalRune
}

type _SGRToken struct {
	controlSequenceToken
	param int
}

type EOFToken struct {
	baseToken
}
