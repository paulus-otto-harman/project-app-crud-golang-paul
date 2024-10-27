package main

import (
	"project/service"
	"sync"
)

func main() {

	// session lifetime satuan detik. gunakan 0 untuk 10 tahun.
	const SessionTimeout = 50

	wg := sync.WaitGroup{}

	wg.Add(1)
	go service.AppContainer(&wg, SessionTimeout)
	wg.Wait()
}
