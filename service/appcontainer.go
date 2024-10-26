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
	if timeout == 0 {
		sessionLifetime = time.Until(time.Now().AddDate(10, 0, 0))
	}
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
