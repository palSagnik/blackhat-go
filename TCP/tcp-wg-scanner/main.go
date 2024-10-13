package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	portStart := 1
	portEnd := 1024
	hostName := "scanme.nmap.org"

	var wg sync.WaitGroup
	for i := portStart; i <= portEnd; i++ {
		
		wg.Add(1)
		fmt.Printf("Checking Port %d\n", i)
		go func(j int) {
			defer wg.Done()
			
			address := fmt.Sprintf("%s:%d", hostName, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}

			conn.Close()
			fmt.Printf("Port %d is open\n", j)
		}(i)

		wg.Wait()
	}
}