package auth

import "testing"

func TestAlwaysPasses(t *testing.T) {
    if true != true {
        t.Error("Somehow true isn't true... that's weird!")
    }
}

// func TestAlwaysFails(t *testing.T) {
//     t.Fatal("This test is guaranteed to fail.")
// }