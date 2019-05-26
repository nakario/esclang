package esc

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

type Program struct {
	ioStream		ioStream
	modules			map[rune]*module
	name			rune
	programCounter	int
	state			*state
}

func Load(source io.Reader) (*Program, error) {
	Logger.Println("Load")
	b, err := ioutil.ReadAll(source)
	if err != nil {
		panic(err)
	}
	s := string(b)
	l := NewLexer(s)
	p := NewParser(l)
	if err := p.Parse(); err != nil {
		return nil, err
	}
	name := '.'
	m := &module{
		name,
		p.labels,
		p.instructions,
	}
	modules := map[rune]*module{name: m}
	prog := &Program{
		ioStream{
			bufio.NewReader(os.Stdin),
			os.Stdout,
		},
		modules,
		name,
		0,
		&state{
			memory{
				make([]uint32, 4096),
				make([]uint32, 4096),
			},
			0,
		},
	}
	return prog, nil
}

func (p *Program) Run() error {
	Logger.Println("Program.Run")
	for p.programCounter < len(p.modules[p.name].instructions) {
		inst := p.modules[p.name].instructions[p.programCounter]
		if err := inst.Exec(p); err != nil {
			if err == Exit {
				Logger.Println("Exit")
				break
			}
			return err
		}
		p.programCounter += 1
	}
	return nil
}
