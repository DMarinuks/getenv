package getenv

import (
	"os"
	"testing"
)

func TestInt(t *testing.T) {
	os.Setenv("FOO", "123")
	want := 123
	got := 0
	Int("FOO", &got,false)
	if got != want {
		t.Errorf("GetInt() = %v, want %v", got, want)
	}
}

func TestBool(t *testing.T) {
	os.Setenv("FOO", "true")
	want := true
	got := false
	Bool("FOO", &got,false)
	if got != want {
		t.Errorf("GetBool() = %v, want %v", got, want)
	}
}

func TestStr(t *testing.T) {
	os.Setenv("FOO", "BAR")
	want := "BAR"
	got := ""
	Str("FOO", &got,false)
	if got != want || len(got) != len(want){
		t.Errorf("GetStr() = %v, want %v", got, want)
	}
}

func TestStrSlice(t *testing.T) {
	os.Setenv("FOO", "BAR, CAT")
	want := []string{"BAR", "CAT"}
	got := StrSlice("FOO",",", false)
	if len(got) == 0 || (got[0] != want[0] && got[1] != want[1]) || len(got[1]) != len(want[1]) {
		t.Errorf("StrSlice() = %v, want %v", got, want)
	}
}

func TestIntSlice(t *testing.T) {
	os.Setenv("FOO", "1, 2")
	want := []int{1, 2}
	got, _ := IntSlice("FOO",",", false)
	if len(got) == 0 || (got[0] != want[0] && got[1] != want[1]) {
		t.Errorf("IntSlice() = %v, want %v", got, want)
	}
}

func TestBoolSlice(t *testing.T) {
	os.Setenv("FOO", "True, false")
	want := []bool{true, false}
	got, _ := BoolSlice("FOO",",", false)
	if len(got) == 0 || (got[0] != want[0] && got[1] != want[1]) {
		t.Errorf("BoolSlice() = %v, want %v", got, want)
	}
}