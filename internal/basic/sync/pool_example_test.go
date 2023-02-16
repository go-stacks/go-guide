package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &User{}
		},
	}

	u1 := pool.Get().(*User)
	u1.ID = 12
	u1.Name = "TOme"

	u3 := pool.Get().(*User)
	u3.ID = 13
	u3.Name = "TO1111me"

	// u1.Reset()
	pool.Put(u1)
	pool.Put(u3)

	u2 := pool.Get().(*User)
	fmt.Println(u2)

	u4 := pool.Get().(*User)
	fmt.Println(u4)
}

type User struct {
	ID   uint64
	Name string
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
}

func (u *User) ChangeName(name string) {
	u.Name = name
}
