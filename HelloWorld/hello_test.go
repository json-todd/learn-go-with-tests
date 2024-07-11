package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Tri Anh", "")
    want := "Hello, Tri Anh"
    assertCorrectMessage(t, got, want)
  })

  t.Run("say 'Hello, World' as a default when nothing (emptry string) is the input", func (t *testing.T){
    got := Hello("","")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in Finnish", func(t *testing.T) {
    got := Hello("Json", "Finnish")
    want:= "Moi, Json"
    assertCorrectMessage(t, got, want)
  })

  t.Run("in Germany", func(t *testing.T) {
    got := Hello("Todd", "Germany")
    want := "Hallo, Todd"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
