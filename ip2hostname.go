package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func getHostnames(ip string) ([]string, error) {
        hostnames, _ := net.LookupAddr(ip)
        return hostnames, nil
}

func main() {
        var ipAddress string

        if len(os.Args) > 1 {
                ipAddress = os.Args[1]
        } else {

        scanner := bufio.NewScanner(os.Stdin)

        if scanner.Scan() {
        ipAddress = strings.TrimSpace(scanner.Text())
        } else {
        fmt.Println("Error reading IP address from standard input.")
        return
        }
        }

        hostnames, _ := getHostnames(ipAddress)
        if hostnames != nil {
          fmt.Println(hostnames[0][:len(hostnames[0])-1])
        return
        }
}
