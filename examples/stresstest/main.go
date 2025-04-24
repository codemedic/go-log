package main

import (
	"context"
	"flag"
	"fmt"
	golog "github.com/codemedic/go-log"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"runtime/pprof"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var randomStrings = []string{}

func init() {
	// Initialize randomStrings with 100 strings of various lengths, between 50 and 1000 characters
	for i := 0; i < 100; i++ {
		str := make([]byte, 50+rand.Intn(950))
		for j := range str {
			str[j] = byte('a' + rand.Intn(26))
		}
		randomStrings = append(randomStrings, string(str))
	}
}

func main() {
	// Parse command-line flags
	backend := flag.String("b", "file", "logging backend: 'syslog' or 'file'")
	concurrency := flag.Int("c", 2, "number of concurrent goroutines")
	duration := flag.Duration("d", 10*time.Second, "how long to run the test")
	cpuprofile := flag.String("cpu-profile", "", "write cpu profile to `file`")
	memprofile := flag.String("mem-profile", "", "write memory profile to `file`")

	flag.Parse()

	// Initialize logger based on backend
	var l golog.Log
	switch strings.ToLower(*backend) {
	case "syslog":
		l = golog.Must(golog.NewSyslog(
			golog.OptionsMust(golog.Options(
				golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
				golog.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
				golog.WithSyslogTag("stress-test"),
			))))
	case "file":
		l = golog.Must(golog.NewLogfile("stress-test.log", 0644,
			golog.OptionsMust(golog.Options(
				golog.WithLevelFromEnv("LOG_LEVEL", golog.Info),
				golog.WithUTCTimestampFromEnv("LOG_UTC", true),
				golog.WithSourceLocationFromEnv("LOG_SOURCE_LOCATION", "short"),
				golog.WithMicrosecondsTimestamp,
			))))
	default:
		log.Fatalf("Unsupported backend: %s", *backend)
	}

	defer l.Close()

	// Set up CPU profiling if requested
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatalf("could not create CPU profile: %v", err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatalf("could not start CPU profile: %v", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Set up memory profiling if requested
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatalf("could not create memory profile: %v", err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		defer func() {
			if err := pprof.Lookup("heap").WriteTo(f, 0); err != nil {
				log.Fatalf("could not write memory profile: %v", err)
			}
		}()
	}

	// The root context
	ctx := context.Background()

	var cancel context.CancelFunc
	// Set a timeout for the test
	if *duration > 0 {
		ctx, cancel = context.WithTimeout(ctx, *duration)
		defer cancel()
	}

	// get a context stoppable by SIGTERM or SIGINT
	stoppableCtx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var counter atomic.Int64

	// Start stress test
	wg := sync.WaitGroup{}
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func(ctx context.Context, id int) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				default:
					// Log a random string
					l.Info("Goroutine %d: %s", id, randomStrings[rand.Intn(len(randomStrings))])
					// Increment the counter
					counter.Add(1)
				}
			}
		}(stoppableCtx, i)
	}

	wg.Wait()
	fmt.Println("Total log messages:", counter.Load())
}
