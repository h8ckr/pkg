package log

import (
	"github.com/pkg/errors"
	"testing"
)

func init() {
	if err := Init("", &infoDepth); err != nil {
		panic(err)
	}
}

func returnErr() error {
	return errors.Errorf("valid test error")
}

func TestInfo(t *testing.T) {
	Info("test")
}

func TestInfof(t *testing.T) {
	Infof("test %s %s", "with", "format")
}

func TestInfoln(t *testing.T) {
	Infoln("test with break")
}

func BenchmarkInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("test")
	}
}

func BenchmarkInfof(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Infof("test %s %s", "with", "format")
	}
}

func BenchmarkInfoln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Infoln("test with break")
	}
}

func TestFatal(t *testing.T) {
	Fatal("fatal")
}

func TestFatalf(t *testing.T) {
	Fatalf("fatal %s %s", "with", "format")
}

func TestFatalln(t *testing.T) {
	Fatalln("fatal with break")
}

func BenchmarkFatal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fatal("fatal")
	}
}

func BenchmarkFatalf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fatalf("fatal %s %S", "with", "format")
	}
}

func BenchmarkFatalln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fatalln("fatal with break")
	}
}

func TestPanic(t *testing.T) {
	validTestErr := returnErr()
	err := errors.Wrap(validTestErr, "wrapped")
	Panic(err, "panic situation")
}

func BenchmarkPanic(b *testing.B) {
	validTestErr := returnErr()
	err := errors.Wrap(validTestErr, "wrapped")
	for i := 0; i < b.N; i++ {
		Panic(err, "panic situation")
	}
}
