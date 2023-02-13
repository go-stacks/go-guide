package sync

import "fmt"

func newErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("下标超出范围，长度 %d, 下标 %d", length, index)
}
