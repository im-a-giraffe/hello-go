package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Run() {
	execWithChannels()
}

func execWithChannels() {
	c := make(chan string)
	go count("sheep", c)

	/**
	 * Break when the channel is not open any more
	 */
	/*
		for {
			msg, open := <- c
			if !open {
				break
			}
			fmt.Println(msg)
		}

	*/

	/**
	 * Only iterate over the results of the channel
	 */
	cx := make(chan string)
	go count("sheep", cx)
	for msg := range cx {
		fmt.Println(msg)
	}
}

func execWithAwait() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("Hello", c)
		wg.Done()
	}()
	wg.Wait()

}

func count(thing string, c chan string) {
	for i := 1; i < 10; i++ {
		fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 200)
	}
	close(c)
}
