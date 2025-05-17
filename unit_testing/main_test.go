package main

import "testing"

func AddTest(a *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		a.Errorf("Sum was incorrect,got:%d,want:%d", got, want)

	}
}
func SubTest(a *testing.T) {
	got := Sub(2, 3)
	want := 5
	if got != want {
		a.Errorf("Sum was incorrect,got:%d,want:%d", got, want)
	}
}

func TestDoMath(a *testing.T) {
	got := doMath(2, 3, Add)
	want := 5
	if got != want {
		a.Errorf("Sum was incorrect,got:%d,want:%d", got, want)

	}
}

/*

func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(7, 5)
	want := 2
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(7, 5, add)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}


*/
