package tcp

import (
	"github.com/palSagnik/blackhat-go"
	
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection Succesful")
	}
}