package iteration

import ("testing"; "fmt")


func TestRepeat(t *testing.T) {
    t.Run("Amount is 6 should return the character repeated 6 times ", func (t *testing.T) {
        repeated := Repeat("t", 6)
        expected := "tttttt"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Amount is 0 should return an empty string", func (t *testing.T) {
        repeated := Repeat("a",0)
        expected := ""

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

      t.Run("Amount is 1 should return a single character", func (t *testing.T) {
        repeated := Repeat("a",1)
        expected := "a"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Amount is a neagtive number should return an empty string", func (t *testing.T) {
        repeated := Repeat("a", -1)
        expected := ""

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Character is an empty string should return an empty string",func (t *testing.T) {
        repeated := Repeat("", 6)
        expected := ""

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Character has 2 characters should return all characters repeated",func (t *testing.T) {
        repeated := Repeat("AB", 3)
        expected := "ABABAB"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

    t.Run("Character has a space should return repeated sequence including the space",func (t *testing.T){
        repeated := Repeat("A B", 3)
        expected := "A BA BA B"

        if repeated != expected {
            t.Errorf("expected %q but got %q", expected, repeated)
        }
    })

}

func BenchmarkRepeat(b *testing.B) {
    for i:= 0; i < b.N; i ++ {
        Repeat("a", 100)
    }
}

func ExampleRepeat() {
    repeated := Repeat("A", 10)
    fmt.Println(repeated)
    // Output: AAAAAAAAAA
}
