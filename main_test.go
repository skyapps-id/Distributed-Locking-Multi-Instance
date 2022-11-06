package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromo(t *testing.T) {
	// Init Promo
	if err := rds.Set(ctx, "PROMO", 30000, 0); err != nil {
		fmt.Println(err)
	}

	t.Run("Positive test multi thread", func(t *testing.T) {
		var wg sync.WaitGroup
		doIncrement := func(n int) {
			for i := 0; i < n; i++ {
				decriment()
			}
			wg.Done()
		}

		wg.Add(3)
		go doIncrement(10000)
		go doIncrement(10000)
		go doIncrement(10000)
		wg.Wait()

		val, _ := rds.Get(ctx, "PROMO")
		assert.Equal(t, val, "0", "Response not valid")
	})

	t.Run("Negative test multi thread", func(t *testing.T) {
		err := decriment()
		assert.Equal(t, err.Error(), "promo is over", "Response not valid")
	})
}
