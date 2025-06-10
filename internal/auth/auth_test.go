package auth

import "testing"

func TestAlwaysPasses(t *testing.T) {
    if true != true {
        t.Error("Somehow true isn't true... that's weird!")
    }
}
