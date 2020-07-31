package getenv

import (
	"os"
	"testing"
)

func TestInt(t *testing.T) {
	os.Setenv("FOO", "123")
	want := 123
	got := 0
	Int("FOO", &got, false)
	if got != want {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}

func TestBool(t *testing.T) {
	os.Setenv("FOO", "true")
	want := true
	got := false
	Bool("FOO", &got, false)
	if got != want {
		t.Errorf("GetEnvBool() = %v, want %v", got, want)
	}
}

func TestStr(t *testing.T) {
	os.Setenv("FOO", "BAR")
	want := "BAR"
	got := ""
	Str("FOO", &got, false)
	if got != want {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}
