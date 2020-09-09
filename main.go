package main

import (
	"context"
	"github.com/WWWWaheb/goKimsufi/pkg/kimsufi"
	"github.com/WWWWaheb/goKimsufi/pkg/logx"
	"github.com/WWWWaheb/goKimsufi/pkg/telegram"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger := logx.NewLogger()

	ctx := context.Background()
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	hwChan := make(chan string)
	notifyChan := make(chan string)

	go kimsufi.StartBot(logger, hwChan, notifyChan)

	go telegram.StartBot(logger, hwChan, notifyChan)

	select {
	case <-ctx.Done():
	case <-sig:
	}
}
