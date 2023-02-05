package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestPingKong(t *testing.T) {
	var Ball int
	table := make(chan int)
	go player(table, 1)
	go player(table, 2)

	table <- Ball

	time.Sleep(1 * time.Second)
	ret := <-table
	fmt.Println("last ret = ", ret)
}

func player(table chan int, index int) {
	for {
		ball := <-table
		fmt.Println(ball, "index = ", index)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
