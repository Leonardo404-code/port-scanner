package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var mutex sync.Mutex

type ScanResult struct {
	Port    string
	State   string
	Service string
}

func ScanPort(protocol, host string, port int) ScanResult {
	result := ScanResult{Port: strconv.Itoa(port) + string("/") + protocol}
	address := host + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}

	defer conn.Close()

	result.State = "Open"

	return result
}

func InitialScan(host string, wg *sync.WaitGroup) {
	mutex.Lock()
	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("udp", host, i))
	}

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", host, i))
	}
	fmt.Println(results)

	mutex.Unlock()
	wg.Done()
}

func WideScan(host string, wg *sync.WaitGroup) {
	mutex.Lock()
	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", host, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", host, i))
	}

	fmt.Println(results)

	mutex.Unlock()
	wg.Done()
}
