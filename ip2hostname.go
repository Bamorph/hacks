package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	// Parse command line flags
	concurrency := flag.Int("c", 1, "Number of concurrent workers")
	flag.Parse()

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Create a channel to send IP addresses to workers
	ipChannel := make(chan string)

	// Start workers
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker(&wg, ipChannel)
	}

	// Read IP addresses from standard input and send them to the channel
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ipAddress := scanner.Text()
		ipChannel <- ipAddress
	}

	// Close the channel once all IP addresses are sent
	close(ipChannel)

	// Wait for all workers to finish
	wg.Wait()
}

func worker(wg *sync.WaitGroup, ipChannel <-chan string) {
	defer wg.Done()

	for ipAddress := range ipChannel {
		hostname := resolveHostname(ipAddress)
		if hostname != "" {
			fmt.Println(hostname)
		}
	}
}

func resolveHostname(ipAddress string) string {
	// Use net.LookupAddr to resolve the hostname
	hostnames, err := net.LookupAddr(ipAddress)
	if err != nil || len(hostnames) == 0 {
		return "" // Return an empty string if no hostname is found or an error occurs
	}

	// Clean up the hostname by removing the trailing dot, if present
	hostname := strings.TrimSuffix(hostnames[0], ".")

	// Return the cleaned-up hostname
	return hostname
}
