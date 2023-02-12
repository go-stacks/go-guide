package sync

import (
	"fmt"
	"sync"
)

type OnceClose struct {
	once sync.Once
}

func (o *OnceClose) Close() error {
	o.once.Do(func() {
		fmt.Println("close")
	})

	return nil
}
