package log

import (
	"github.com/pkg/errors"
	"reflect"
	"syscall"
	"time"
)

var (
	debugDepth   = 4
	infoDepth    = 3
	errOnlyDepth = 2
	silentDepth  = 1
)

type level struct {
	Depth       *int
	ValidDepths struct {
		Debug   *int
		Info    *int
		ErrOnly *int
		Silent  *int
	}
	MinDepth *int
	MaxDepth *int
}

func (l *level) get() *int {
	return l.Depth
}

func (l *level) setMaxDepth() {
	values := l.getStructFieldIntPtrValues()
	if len(values) == 0 {
		return
	}
	l.MaxDepth = values[0]
	if len(values) > 1 {
		for candidate := range values {
			if candidate < *l.MaxDepth {
				l.MaxDepth = &candidate
			}
		}
	}
}

func (l *level) setMinDepths() {
	values := l.getStructFieldIntPtrValues()
	if len(values) == 0 {
		return
	}
	l.MinDepth = values[0]
	if len(values) > 1 {
		for candidate := range values {
			if candidate < *l.MinDepth {
				l.MinDepth = &candidate
			}
		}
	}
}

func (l *level) getStructFieldIntPtrValues() []*int {
	var validDepthSlice []*int
	validDepthsMap := make(map[int]*int)
	validDepthsValues := reflect.ValueOf(l.ValidDepths)
	numFields := validDepthsValues.NumField()
	for i := 0; i < numFields; i++ {
		validDepthsMap[i] = validDepthsValues.Field(i).Interface().(*int)
	}

	for _, val := range validDepthsMap {
		validDepthSlice = append(validDepthSlice, val)
	}
	return validDepthSlice
}

func (l *level) set(newDepth *int) error {
	if *newDepth < *l.MinDepth {
		return errors.Errorf("minimum allowed log Depth is %d", *l.MinDepth)
	} else if *newDepth > *l.MaxDepth {
		return errors.Errorf("maximum allowed log Depth is %d", *l.MaxDepth)
	}

	l.Depth = newDepth
	return nil
}

type logger struct {
	Level  *level
	Prefix string
}

func (l *logger) info(bytes []byte) {
	syscall.Write(1, bytes)
}

func (l *logger) error(bytes []byte) {
	syscall.Write(2, bytes)
}

func (l *logger) setPrefix(prefix string) {
	l.Prefix = prefix
}

func (l *logger) getPrefix() *string {
	if l.Prefix == "" {
		t := time.Now().Format("Mon 2006-01-02 15:04:05 MST")
		return &t
	}
	return &l.Prefix
}

func (l *logger) setDepth(depth *int) {
	l.Level.Depth = depth
}

type commander interface {
	info(bytes []byte)
	error(bytes []byte)
	setPrefix(prefix string)
	getPrefix() *string
	setDepth(depth *int)
}

func new(prefix string, depth *int) *commander {
	if depth == nil {
		depth = &errOnlyDepth
	}
	l := logger{
		Level: &level{
			Depth: nil,
			ValidDepths: struct {
				Debug   *int
				Info    *int
				ErrOnly *int
				Silent  *int
			}{
				Debug:   &debugDepth,
				Info:    &infoDepth,
				ErrOnly: &errOnlyDepth,
				Silent:  &silentDepth,
			},
			MinDepth: nil,
			MaxDepth: nil,
		},
		Prefix: prefix,
	}
	l.Level.setMinDepths()
	l.Level.setMaxDepth()
	l.Level.set(depth)
	var c commander = &l
	return &c
}
