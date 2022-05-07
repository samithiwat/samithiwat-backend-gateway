package main

import (
	"context"
	"fmt"
	"github.com/samithiwat/samithiwat-backend-gateway/src/config"
	"github.com/samithiwat/samithiwat-backend-gateway/src/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-s

		log.Printf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Printf("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}

	r := router.NewFiberRouter()

	go func() {
		if err := r.Listen(fmt.Sprintf(":%v", conf.App.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"server": func(ctx context.Context) error {
			return r.Shutdown()
		},
	})

	<-wait
}
