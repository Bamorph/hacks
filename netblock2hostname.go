package main

import (
        "bufio"
        "flag"
        "fmt"
        "net"
        "os"
        "sync"
)

var concurrency int

func init() {
        flag.IntVar(&concurrency, "c", 10, "Concurrency level")
        flag.Parse()
}

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        var wg sync.WaitGroup
        semaphore := make(chan struct{}, concurrency)

        for scanner.Scan() {
                ipRange := scanner.Text()
                ips, _ := getIPsFromRange(ipRange)
                //if err != nil {
                //      fmt.Printf("Error processing %s: %v\n", ipRange, err)
                //      continue
                //}

                for _, ip := range ips {
                        wg.Add(1)
                        semaphore <- struct{}{}
                        go func(ip string) {
                                defer wg.Done()
                                defer func() { <-semaphore }()
                                hostnames, _ := getHostnames(ip)
                                //if err != nil {
                                //      fmt.Printf("Error resolving hostnames for %s: %v\n", ip, err)
                                //      return
                                //}
                                if hostnames != nil {
                                        og_hostname := hostnames[0]
                                        ot_hostname := og_hostname[:len(og_hostname)-1]
                                        fmt.Println(ot_hostname)
                                        //fmt.Printf("IP: %s\nHostnames: %v\n\n", ip, hostnames)
                                        return
                                }
                        }(ip)
                }
        }

        wg.Wait()

        if err := scanner.Err(); err != nil {
                fmt.Println("Error reading input:", err)
        }
}

func getIPsFromRange(ipRange string) ([]string, error) {
        ips := []string{}

        ip, ipnet, err := net.ParseCIDR(ipRange)
        if err != nil {
                return nil, err
        }

        for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
                ips = append(ips, ip.String())
        }

        return ips, nil
}

func inc(ip net.IP) {
        for j := len(ip) - 1; j >= 0; j-- {
                ip[j]++
                if ip[j] > 0 {
                        break
                }
        }
}

func getHostnames(ip string) ([]string, error) {
        hostnames, err := net.LookupAddr(ip)
        if err != nil {
                return nil, err
        }

        return hostnames, nil
}
       
