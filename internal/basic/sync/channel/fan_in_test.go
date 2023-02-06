package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	ch1 := search("test1")
	ch2 := search("test2")
	ch3 := search("test3")

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case msg := <-ch3:
			fmt.Println(msg)
		}
	}
}

func search(msg string) chan string {

	ch := make(chan string)

	go func() {
		var i int
		for {
			ch <- fmt.Sprintf("get %s %d", msg, i)
			i++
			time.Sleep(time.Second)
		}
	}()

	return ch
}
