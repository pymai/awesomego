package igoroutine

import (
	"fmt"
	"time"
)

func SleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...")
}
