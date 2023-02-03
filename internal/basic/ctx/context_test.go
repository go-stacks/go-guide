// MIT License
//
// Copyright (c) 2023 henrysworld
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ctx

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type key string

// TestContext 展示Context WithTimeout、WithValue、WithDeadline的用法
func TestContext(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	key := key("key")
	vCtx := context.WithValue(timeoutCtx, key, "value1")
	value := vCtx.Value(key)
	fmt.Println(value)
	dCtx, dCancel := context.WithDeadline(vCtx, time.Now().Add(time.Second))
	defer dCancel()

	time.Sleep(2 * time.Second)

	err := timeoutCtx.Err()
	dErr := dCtx.Err()
	fmt.Println("t = ", err)
	fmt.Println("d = ", dErr)

}

// TestContextValue 验证Context传值功能，子Context可以获取到父Context中放入的值，但是父Context不能获取到子Context放入的值
func TestContextValue(t *testing.T) {
	parentKey := key("parentKey")
	ctx := context.Background()
	parent := context.WithValue(ctx, parentKey, "parent")

	subKey := key("subKey")
	sub := context.WithValue(parent, subKey, "sub")

	err := sub.Err()
	if err != nil {
		fmt.Println(err)
	}

	p2PValue := parent.Value(parentKey)
	p2SValue := parent.Value(subKey)

	s2PValue := sub.Value(parentKey)
	S2SValue := sub.Value(subKey)

	fmt.Println("p2PValue", p2PValue)
	fmt.Println("p2SValue", p2SValue)
	fmt.Println("s2PValue", s2PValue)
	fmt.Println("S2SValue", S2SValue)
}

// TestContextCancel Context Cancel用法。当父Context被cancel以后，其下面的所有子context都会被取消，一般用在超时控制
func TestContextCancel(t *testing.T) {
	bg := context.Background()
	parentCtx, cancel1 := context.WithTimeout(bg, time.Second)
	subCtx, _ := context.WithTimeout(parentCtx, time.Second*10)

	go func() {
		<-subCtx.Done()
		fmt.Println("timeout")
	}()

	time.Sleep(500 * time.Millisecond)

	//cancel2()
	cancel1()

	fmt.Println("parentCtx", parentCtx.Err())
	fmt.Println("subCtx", subCtx.Err())
}

func slowBusiness() {
	time.Sleep(5 * time.Second)
}

// TestContextListen 监听goruntine执行完成还是context超时。
func TestContextListen(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	end := make(chan struct{}, 1)

	go func() {
		slowBusiness()
		end <- struct{}{}
	}()

	status := timeoutCtx.Done()

	select {
	case <-end:
		fmt.Println("finish")
	case <-status:
		fmt.Println("timeout")
	}

}
