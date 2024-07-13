package iteration

import "testing"

func TestRepeat(t *testing.T) {
    t.Run("Passing 6 to function should have the string repeated 6 times ", func (t *testing.T) {
      repeated := Repeat("t", 6)
      expected := "tttttt"

      if repeated != expected {
          t.Errorf("expected %q but got %q", expected, repeated)
      }
    })
}


