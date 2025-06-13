package auth

import "testing"

func TestAlwaysPasses(t *testing.T) {
	a := 1
	if a != 1 {
		t.Error("Somehow true isn't true... that's weird!")
	}
}

// func TestAlwaysFails(t *testing.T) {
//     t.Fatal("This test is guaranteed to fail.")
// }
