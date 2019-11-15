package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main()  {
	ips := []string{
		"127.0.0.1",
		"google.com",
		"baidu.com",
	}
	// 扫描端口
	Dials(ips)
}

func Dials(ips []string)  {

	var (
		wg 		= &sync.WaitGroup{}
		timeOut = time.Millisecond * 200
		unUse 	= map[string][]int{}
	)
	for port := 1; port <= 100; port++ {
		wg.Add(len(ips))
		for _, h := range ips {
			go func(host string, port int) {
				status := isOpen(host, port, timeOut)
				if status {
					unUse[host] = append(unUse[host], port)
				}
				wg.Done()
			}(h, port)
		}
	}
	wg.Wait()

	fmt.Println(unUse)
}

func isOpen(host string, port int, timeOut time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeOut)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false

}