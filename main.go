package main

import (
	"project/service"
	"sync"
)

func main() {
	const SessionTimeout = 50

	wg := sync.WaitGroup{}

	wg.Add(1)
	go service.AppContainer(&wg, SessionTimeout)
	wg.Wait()
}
