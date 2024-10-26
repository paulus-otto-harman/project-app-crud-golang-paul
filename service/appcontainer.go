package service

import (
	"context"
	"project/view"
	"sync"
	"time"
)

func AppContainer(wg *sync.WaitGroup, timeout int) {
	defer wg.Done()
	sessionLifetime := time.Duration(timeout) * time.Second
	for {
		loginScreen := view.Login{}
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), sessionLifetime)
			defer cancel()

			(&loginScreen).Render(ctx)
		}()
		if loginScreen.Username == "0" {
			return
		}
	}
}
