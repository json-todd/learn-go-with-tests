package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
    t.Run("collection of 5 numbers", func(t *testing.T) {
      numbers := []int{1,2,3,4,5}

      got := Sum(numbers)
      want := 15

      if got != want {
          t.Errorf("got %d want %d given, %v", got, want, numbers)
      }
    })

    t.Run("collection of any size", func(t *testing.T){
        numbers := []int{1,2,3}

        got := Sum(numbers)
        want := 6

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("an empty slice",func(t *testing.T){
        numbers := []int{}

        got := Sum(numbers)
        want := 0

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })


    t.Run("negative and positive numbers", func(t *testing.T){
        numbers := []int {-3, -2, -1, 1, 2, 3}

        got := Sum(numbers)
        want := 0

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }
    })

    t.Run("zero are supported in collection", func(t *testing.T){
        numbers:= []int{0,0,0,0,0,0,0,0,0,0}

        got := Sum(numbers)
        want := 0

        if got != want {
            t.Errorf("got %d want %d given, %v", got, want, numbers)
        }

    })
}

func TestSumAll(t *testing.T) {
    got := SumAll([]int{1,2}, []int{0,9})
    want:= []int{3,9}

    if ! reflect.DeepEqual(got,want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func TestSumTail(t *testing.T) {
		got := SumAllTail([]int{1,2}, []int{0,9})
		want := []int{2,9}

		if !reflect.DeepEqual(got,want) {
				t.Errorf("got %v want %v", got, want)
		}
}