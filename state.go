package esc

import "io"

type memory struct {
	data		[]uint32
	pointers	[]uint32
}

type ioStream struct {
	input	io.RuneReader
	output	io.StringWriter
}

type state struct {
	memory
	primalPtr      uint32
}
