package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error reading from standard input.")
		os.Exit(1)
	}

	netblockStr := strings.TrimSpace(scanner.Text())
	ip, ipnet, err := net.ParseCIDR(netblockStr)
	if err != nil {
		fmt.Println("Error parsing netblock:", err)
		os.Exit(1)
	}

	fmt.Printf("IP addresses in %s:\n", netblockStr)
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incrementIP(ip) {
		fmt.Println(ip)
	}
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
