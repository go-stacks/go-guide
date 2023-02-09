package channel

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}

func TestPipelineCalculation(t *testing.T) {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, v := range integers {
				select {
				case <-done:
					return
				case intStream <- v:

				}
			}
		}()

		return intStream
	}

	mutiply := func(done <-chan interface{}, intStream <-chan int, mutiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- v * mutiplier:

				}
			}
		}()

		return multipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case addedStream <- v + additive:

				}
			}
		}()

		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := mutiply(done, add(done, mutiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}

// 第一个阶段，数字的生成器
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// 筛选，排除不能够被prime整除的数
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // 获取上一个阶段的
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

func TestFindPrimeNumber(t *testing.T) {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 100000; i++ {
		prime := <-ch // 获取上一个阶段输出的第一个数，其必然为素数
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1 // 前一个阶段的输出作为后一个阶段的输入。
	}
}
