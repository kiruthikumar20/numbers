package main

import "testing"

func AddTesting(s *testing.T) {
	got := Add(6, 6)
	want := 11
	if got != want {
		s.Errorf("Sum was incorrect,got:%d, want : %d", got, want)
	}
}

/*
func TestAdd(t *testing.T) {
total := Add(5, 5)
if total != 10 {
t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
}
}
```

*/
