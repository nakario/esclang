package esc

import (
	"github.com/pkg/errors"
	"math/bits"
)

type Instruction func(p *Program) error

func (i Instruction) Exec(p *Program) error {
	return i(p)
}

// foreground

func copyInstruction(r rune) Instruction {
	Logger.Println("create copyInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call copyInstruction")
		s.data[s.pointers[s.primalPtr]] = uint32(r)
		return nil
	}
}

func incrementInstruction(_ rune) Instruction {
	Logger.Println("create incrementInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call incrementInstruction")
		s.data[s.pointers[s.primalPtr]] += 1
		return nil
	}
}

func inputInstruction(_ rune) Instruction {
	Logger.Println("create inputInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call inputInstruction")
		r, _, err := p.ioStream.input.ReadRune()
		if err != nil {
			return err
		}
		s.data[s.pointers[s.primalPtr]] = uint32(r)
		return nil
	}
}

func rotateLeftInstruction(_ rune) Instruction {
	Logger.Println("create rotateLeftInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call rotateLeftInstruction")
		d := s.data[s.pointers[s.primalPtr]]
		s.data[s.pointers[s.primalPtr]] = bits.RotateLeft32(d, 1)
		return nil
	}
}

func rotateRightInstruction(_ rune) Instruction {
	Logger.Println("create rotateRightInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call rotateRightInstruction")
		d := s.data[s.pointers[s.primalPtr]]
		s.data[s.pointers[s.primalPtr]] = bits.RotateLeft32(d, -1)
		return nil
	}
}

func outputInstructioin(_ rune) Instruction {
	Logger.Println("create outputInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call outputInstruction")
		d := s.data[s.pointers[s.primalPtr]]
		_, err := p.ioStream.output.WriteString(string([]rune{rune(d)}))
		return err
	}
}

func decrementInstruction(_ rune) Instruction {
	Logger.Println("create decrementInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call decrementInstruction")
		s.data[s.pointers[s.primalPtr]] -= 1
		return nil
	}
}

func swapInstruction(_ rune) Instruction {
	Logger.Println("create swapInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call swapInstruction")
		a := s.data[s.pointers[s.primalPtr]]
		b := s.pointers[s.primalPtr]
		s.data[s.pointers[s.primalPtr]] = b
		s.pointers[s.primalPtr] = a
		return nil
	}
}

// background

func callInstruction(name rune) Instruction {
	Logger.Println("create callInstruction")
	return func(p *Program) error {
		Logger.Println("call callInstruction")
		if name == '.' {
			return nil
		}
		// TODO: call p.modules[name]
		return nil
	}
}

func incrementPtrInstruction(_ rune) Instruction {
	Logger.Println("create incrementInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call incrementInstruction")
		s.pointers[s.primalPtr] += 1
		return nil
	}
}

func jumpIfZeroInstruction(label rune) Instruction {
	Logger.Println("create jumpIfZeroInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call jumpIfZeroInstruction")
		if s.data[s.pointers[s.primalPtr]] == 0 {
			p.programCounter = p.modules[p.name].labels[label]
		}
		return nil
	}
}

func incrementPrimalPtrInstruction(_ rune) Instruction {
	Logger.Println("create incrementPrimalPtrInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call incrementPrimalPtrInstruction")
		s.primalPtr += 1
		return nil
	}
}

func decrementPrimalPtrInstruction(_ rune) Instruction {
	Logger.Println("create decrementPrimalPtrInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call decrementPrimalPtrInstruction")
		s.primalPtr -= 1
		return nil
	}
}

func decrementPtrInstruction(_ rune) Instruction {
	Logger.Println("create decrementPtrInstruction")
	return func(p *Program) error {
		s := p.state
		Logger.Println("call decrementPtrInstruction")
		s.pointers[s.primalPtr] -= 1
		return nil
	}
}

var Exit = errors.New("exit this module")

func exitInstruction(_ rune) Instruction {
	Logger.Println("create exitInstruction")
	return func(p *Program) error {
		Logger.Println("call exitInstruction")
		return Exit
	}
}

// nop

func nopInstruction(_ rune) Instruction {
	return func(_ *Program) error {
		return nil
	}
}
