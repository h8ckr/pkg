package log

import (
	"github.com/pkg/errors"
	"reflect"
	"syscall"
)

var (
	debugDepth = 4
	infoDepth = 3
	errOnlyDepth = 2
	silentDepth = 1
)

type level struct {
	depth *int
	validDepths struct{
		debugDepth *int
		infoDepth *int
		errOnlyDepth *int
		silentDepth *int
	}
	minDepth *int
	maxDepth *int
}

func (l *level) get() *int {
	return l.depth
}

func (l *level) setMaxDepth() {
	values := l.getStructFieldIntPtrValues()
	if len(values) == 0 {
		return
	}
	l.maxDepth = values[0]
	if len(values) > 1 {
		for candidate := range values {
			if candidate < *l.maxDepth {
				l.maxDepth = &candidate
			}
		}
	}
}

func (l *level) setMinDepths() {
	values := l.getStructFieldIntPtrValues()
	if len(values) == 0 {
		return
	}
	l.minDepth = values[0]
	if len(values) > 1 {
		for candidate := range values {
			if candidate < *l.minDepth {
				l.minDepth = &candidate
			}
		}
	}
}

func (l *level) getStructFieldIntPtrValues() []*int {
	var validDepthSlice []*int
	validDepthsMap := make(map[int]*int)
	validDepthsValues := reflect.ValueOf(l.validDepths)
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
	if *newDepth < *l.minDepth {
		return errors.Errorf("minimum allowed log depth is %d", *l.minDepth)
	} else if *newDepth > *l.maxDepth {
		return errors.Errorf("maximum allowed log depth is %d", *l.maxDepth)
	}

	l.depth = newDepth
	return nil
}

type logger struct {
	level *level
	prefix string
}

func (l *logger) info(bytes []byte) {
	syscall.Write(1, bytes)
}

func (l *logger) error(bytes []byte) {
	syscall.Write(2, bytes)
}

func (l *logger) setPrefix(prefix string) {
	l.prefix = prefix
}

type commander interface {
	info(arg interface{})
	error(arg interface{})
}