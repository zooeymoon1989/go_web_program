package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var condLock = *sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	var newYearTime = time.Unix(1577808000 , 0)
	var now = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), newYearTime.Sub(now))
	defer cancel()

	wg.Add(1)

	fmt.Println("新年进入倒计时 ლ(･ ิω･ิლ)")

	go func() {
		wg.Done()
		condLock.L.Lock()
		defer condLock.L.Unlock()
		condLock.Wait()
		fmt.Println("2020✅")
		fmt.Println("💃💃💃💃💃💃💃💃💃💃💃💃💃💃💃")
		fmt.Println("🎆🎆🎆🎆🎆🎆🎆🎆🎆🎆🎆🎆🎆🎆🎆")
	}()

	wg.Wait()

	select {
	case <-ctx.Done():
		fmt.Println("2019☑️")
		condLock.Signal()
	}
}
