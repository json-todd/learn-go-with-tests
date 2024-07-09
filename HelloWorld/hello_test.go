package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Tri Anh")
    want := "Hello, Tri Anh"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

  t.Run("say 'Hello, World' as a default when nothing (emptry string) is the input", func (t *testing.T){
    got := Hello("")
    want := "Hello, World"

    if got != want {
      t.Errorf("got %q want %q", got ,want)
    }
  })
}
