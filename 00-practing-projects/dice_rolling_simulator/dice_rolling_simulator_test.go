package main

import (
	"fmt"
	"testing"
)

func TestDiceRolling(t *testing.T) {
	t.Run("message", func(t *testing.T) {
		got := Message(1)

		want := fmt.Sprint(DefaultMessage, 1)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("dice rolling should be within 1-6", func(t *testing.T) {
		for i := 0; i < 10000; i++ {
            got := RollDice()

			if got < 1 || got > 6 {
				t.Fatalf("dice result need to be between 1-6 got %d", got)
			}
		}
	})

    t.Run("roll dice n times", func(t *testing.T) {
        got := TimesRoll(3)

        if len(got) != 3 {
            t.Errorf("Expected dice to be rolled 3 times got %d", len(got))
        }
    })

    t.Run("should roll the dice only once if no number of times are provided", func(t *testing.T) {
        got := TimesRoll(0)

        if len(got) != 1 {
            t.Errorf("Expect TimesRoll to roll the dice only once if no quantity is provided got %d", got)
        }
    })
}
