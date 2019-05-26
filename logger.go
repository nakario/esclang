package esc

import (
	"github.com/pkg/errors"
	"log"
	"os"
)

var Logger *log.Logger = log.New(os.Stderr, "[DEBUG]", log.LstdFlags | log.Lshortfile)

var (
	eWrap = errors.Wrap
	eCause = errors.Cause
)
