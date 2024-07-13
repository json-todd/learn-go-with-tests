package iteration

import "testing"

func TestRepeat(t *testing.T) {
    t.Run("Supply 6 as amount arg to function should return the string repeated 6 times ", func (t *testing.T) {
        repeated := Repeat("t", 6)
        expected := "tttttt"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Supply an empty string to function should return an empty string",func (t *testing.T){
        repeated := Repeat("", 6)
        expected := ""

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Supply 0 as amount arg should return an empty strinbg", func (t *testing.T) {
        repeated := Repeat("a",0)
        expected := ""

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

      t.Run("Supply 1 as amount arg should return a single character", func (t *testing.T) {
        repeated := Repeat("a",1)
        expected := "a"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Supply a neagtive number as amount arg should return an empty string", func (t *testing.T) {
        repeated := Repeat("a", -1)
        expected := ""

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })
}


