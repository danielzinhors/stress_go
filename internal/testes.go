package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

const APP_NAME = "STRESS_GO"

func RunStressTester(url string, requests int64, concurrency int64) {

	MessageLog("Rodando com os parametross url=\"%s\" request=%d concurrency=%d", url, requests, concurrency)

	if concurrency > requests {
		concurrency = requests
		MessageLog("Reduzindo o numero de workers to %d", concurrency)
	}

	loadPerRacer, err := GetLoadRacer(requests, concurrency)
	if err != nil {
		MessageLog("error: %s", err)
		os.Exit(1)
	}

	MessageLog("Usando Racer load: %v", *loadPerRacer)

	ctx := context.Background()
	relatorio := NewRelatorio(url, concurrency)
	wg := &sync.WaitGroup{}
	start := time.Now()
	printReportOnExit(relatorio, &start)

	for i, load := range *loadPerRacer {
		racer := NewRacer(strconv.FormatInt(int64(i+1), 10), url, load, relatorio)
		wg.Add(1)
		go racer.Run(ctx, wg)
	}

	wg.Wait()

	duration := time.Since(start)
	relatorio.Duration = &duration
	relatorio.PrintRelatorio(true)
}

func printReportOnExit(relatorio *Relatorio, start *time.Time) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Print("\n\n")
		duration := time.Since(*start)
		relatorio.Duration = &duration
		relatorio.PrintRelatorio(false)
		os.Exit(1)
	}()
}
