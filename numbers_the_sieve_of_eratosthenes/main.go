package main

import (
	"fmt"
)

func generate(ch chan<- int, N int) {
    for i := 2; i <= N; i++ {
        ch <- i
    }
    close(ch)
}

func filter(in <- chan int, p int) chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			if num % p != 0 {
				out <- num
			}
		}
		close(out)
	}()
	return out
}

func main() {
	const N = 100
    ch := make(chan int)
    go generate(ch, N)
	fmt.Println(ch)

	for {
        p, ok := <-ch
        if !ok {
            break
        }
        fmt.Println(p) 
        ch = filter(ch, p)
    }
}