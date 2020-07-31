package getenv

import (
	"os"
	"testing"
)

func TestGetEnvInt(t *testing.T) {
	os.Setenv("FOO", "123")
	want := 123
	got := 0
	GetEnvInt("FOO", &got, false)
	if got != want {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}

func TestGetEnvBool(t *testing.T) {
	os.Setenv("FOO", "true")
	want := true
	got := false
	GetEnvBool("FOO", &got, false)
	if got != want {
		t.Errorf("GetEnvBool() = %v, want %v", got, want)
	}
}

func TestGetEnvStr(t *testing.T) {
	os.Setenv("FOO", "BAR")
	want := "BAR"
	got := ""
	GetEnvStr("FOO", &got, false)
	if got != want {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}