package sync

import "testing"

func TestOnce(t *testing.T) {
	cc := &OnceClose{}

	for i := 0; i < 100; i++ {
		cc.Close()
	}
}
