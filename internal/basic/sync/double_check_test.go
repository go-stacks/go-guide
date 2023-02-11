package sync

import (
	"fmt"
	"testing"
)

func TestDoubleCheck(t *testing.T) {
	safeMap := SafeMap[string, int]{
		m: make(map[string]int),
	}
	v, b := safeMap.LoadOrStore("ch", 10)
	v, b = safeMap.LoadOrStore("ch", 10)
	fmt.Println(v, b)
}
