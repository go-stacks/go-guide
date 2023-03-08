package channel

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

//使用30个并发，打印0-99

//1.使用chan完成，并且使用close chan来防止goroutine泄露

type Token int

func TestChanClose(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	var wg sync.WaitGroup
	var m sync.Map
	ch := make(chan Token, 100)
	done := make(chan struct{})
	maxNum := 100
	sum := 0

	for i := 0; i <= 99; i++ {
		ch <- Token(i)
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go task(ch, done, maxNum, i, &m, &wg)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		if _, ok := m.Load(i); ok {
			sum++
		}
	}

	fmt.Println("总数： ", sum)

}

func task(ch chan Token, done chan struct{}, maxNum int, index int, m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("go = %d, value = %d\n", index, v)
			m.Store(int(v), struct{}{})
			v++
			if int(v) == maxNum {
				close(done)
			}
		case <-done:
			return
		}
	}
}

// 2.控制goroutine 方法的调度，一个挨着一个执行。并使用ctx来控制goroutine的生命周期
func TestSerial(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	var m sync.Map

	maxNum := 100
	goroutineCnt := 30
	chs := make([]chan Token, maxNum)
	done := make(chan Token)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < maxNum; i++ {
		chs[i] = make(chan Token)
	}

	for i := 0; i < goroutineCnt; i++ {
		go taskSerial(ctx, chs[i], chs[(i+1)%30], done, &m, i, maxNum)
	}

	chs[0] <- 0

	select {
	case <-done:
		cancel()
	}

	time.Sleep(time.Second)
}

func taskSerial(ctx context.Context, ch chan Token, nextCh chan Token, done chan Token, m *sync.Map, index int, maxNun int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v := <-ch:
			fmt.Printf("go = %d, value = %d\n", index, v)
			v++
			if int(v) == maxNun {
				close(done)
				return
			}
			nextCh <- v
		}
	}
}
