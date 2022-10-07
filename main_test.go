package main

import "testing"

func TestReverse(t *testing.T) {
	if got, want := Reverse("reverse"), "esrever"; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	t.Run("empty string", func(t *testing.T) {
		if got, want := Reverse(""), ""; got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}
