package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Tri Anh")
  want := "Hello, Tri Anh"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
