package getenv

import (
	"os"
	"testing"
)

func TestInt(t *testing.T) {
	os.Setenv("FOO", "123")
	want := 123
	got, _ := Int("FOO",false)
	if got != want {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}

func TestBool(t *testing.T) {
	os.Setenv("FOO", "true")
	want := true
	got, _ := Bool("FOO", false)
	if got != want {
		t.Errorf("GetEnvBool() = %v, want %v", got, want)
	}
}

func TestStr(t *testing.T) {
	os.Setenv("FOO", "BAR")
	want := "BAR"
	got := Str("FOO", false)
	if got != want {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}

func TestStrSlice(t *testing.T) {
	os.Setenv("FOO", "BAR, CAT")
	want := []string{"BAR", "CAT"}
	got := StrSlice("FOO",",", false)
	if len(got) == 0 || (got[0] != want[0] && got[1] != want[1]) {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}

func TestIntSlice(t *testing.T) {
	os.Setenv("FOO", "1, 2")
	want := []int{1, 2}
	got, _ := IntSlice("FOO",",", false)
	if len(got) == 0 || (got[0] != want[0] && got[1] != want[1]) {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}

func TestBoolSlice(t *testing.T) {
	os.Setenv("FOO", "True, false")
	want := []bool{true, false}
	got, _ := BoolSlice("FOO",",", false)
	if len(got) == 0 || (got[0] != want[0] && got[1] != want[1]) {
		t.Errorf("GetEnvInt() = %v, want %v", got, want)
	}
}