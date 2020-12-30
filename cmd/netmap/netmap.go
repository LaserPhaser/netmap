package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	numCpu := runtime.NumCPU()

	chunks := chunks(numCpu, 65535)

	var wg sync.WaitGroup
	for _, ch := range chunks {
		wg.Add(1)
		go scanPort(&wg, ch[0], ch[1])
	}
	fmt.Println("Scanning started")

	wg.Wait()
	fmt.Println("Scanning finished")

}

func scanPort(wg *sync.WaitGroup, from int, to int) {

	defer wg.Done()
	for port := from; port <= to; port++ {

		conn, err := net.DialTimeout("tcp", "localhost:"+strconv.Itoa(port), time.Second)
		if err == nil {
			fmt.Println("Found open port", port)
			conn.Close()
		}
	}
}

func chunks(num_cpu int, max_value int) [][2]int {

	var divided [][2]int
	var chunk [2]int
	chunkSize := (max_value + num_cpu - 1) / num_cpu

	for start := 0; start < max_value; start += chunkSize + 1 {
		end := start + chunkSize

		if end > max_value {
			end = max_value
		}
		chunk[0] = start
		chunk[1] = end
		divided = append(divided, chunk)

	}
	return divided

}
