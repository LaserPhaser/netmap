package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	portStart := flag.Int("port_start", 0, "scan from")
	portEnd := flag.Int("port_end", 65535, "scan to")
	scanProtocol := flag.String("scan_protocol", "tcp", "Protocol to use [TCP/UDP]")
	scanHost := flag.String("host", "localhost", "Host to scan")
	flag.Parse()

	numCpu := runtime.NumCPU()
	chunks := chunks(numCpu, *portStart, *portEnd)

	var wg sync.WaitGroup
	for _, ch := range chunks {
		wg.Add(1)
		go scanPort(&wg, ch[0], ch[1], *scanProtocol, *scanHost)
	}
	fmt.Println("Scanning started for " + *scanHost + " for ports range[" + strconv.Itoa(*portStart) + ":" + strconv.Itoa(*portEnd) + "], protocol " + *scanProtocol)
	wg.Wait()
	fmt.Println("Scanning finished")

}

func scanPort(wg *sync.WaitGroup, from int, to int, scan_protocol string, scan_host string) {

	defer wg.Done()
	for port := from; port <= to; port++ {

		conn, err := net.DialTimeout(scan_protocol, scan_host+":"+strconv.Itoa(port), time.Second)
		if err == nil {
			fmt.Println("Found open port", port)
			conn.Close()
		}
	}
}

func chunks(num_cpu int, from int, to int) [][2]int {

	var divided [][2]int
	var chunk [2]int
	chunkSize := (to - from + num_cpu - 1) / num_cpu

	for start := from; start < to; start += chunkSize + 1 {
		end := start + chunkSize

		if end > to {
			end = to
		}
		chunk[0] = start
		chunk[1] = end
		divided = append(divided, chunk)

	}
	return divided

}
