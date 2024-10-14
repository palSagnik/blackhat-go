package main

import (
	"fmt"
	"net"
	"sort"
)

var (
	hostName = "scanme.nmap.org"
	portStart = 1
	portEnd = 10240
	workerCount = 300
)

// Providing a pool of workers
func worker(ports chan int, results chan int) {

	for p := range ports {
		address := fmt.Sprintf("%s:%d", hostName, p)
		conn, err := net.Dial("tcp", address) 
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func scannerLogic() (res []int) {
	
	ports := make(chan int, workerCount)
	results := make(chan int)

	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := portStart; i <= portEnd; i++ {
			ports <- i
		}
	}()
	
	// getting openPorts from results channel
	// 0 -> notOpen
	// !0 -> open
	for i := portStart; i <= portEnd; i++ {
		port := <- results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)

	return openPorts
}

func main() {
	
	openPorts := scannerLogic()

	sort.Ints(openPorts)
	fmt.Println("Open Ports")
	for _, port := range(openPorts) {
		fmt.Printf("%s:%d\n", hostName, port)
	}
}
