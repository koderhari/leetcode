package justgo

import (
	"fmt"
	"sync"
	"testing"
)

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		wg.Add(len(cs))
		for _, ch := range cs {
			go func(ch <-chan int) {
				defer wg.Done()
				for v := range ch {
					out <- v
				}
			}(ch)
		}

		wg.Wait()
	}()

	return out
}

func Test_merge(t *testing.T) {

	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)
	out := merge(ch1, ch2, ch3)

	var wg sync.WaitGroup
	for i, v := range []chan int{ch1, ch2, ch3} {
		wg.Add(1)
		go func(v int, ch chan int) {
			defer wg.Done()
			defer close(ch)
			for i := 0; i < 10; i++ {
				ch <- v
			}
		}(i, v)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		idx := 0
		for v := range out {
			fmt.Println(idx)
			idx++
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
