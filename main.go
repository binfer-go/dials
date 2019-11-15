package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main()  {

	startPort 	:= flag.Int("start-port", 80, "开始扫描的端口")
	endPort 	:= flag.Int("end-port", 100, "结束扫描的端口")
	timeOut 	:= flag.Duration("timeout", time.Millisecond * 200, "超时时间")
	flag.Parse()
	ips := []string{
		"127.0.0.1",
		"google.com",
		"baidu.com",
	}
	// 扫描端口
	Dials(ips, *startPort, *endPort, *timeOut)
}

func Dials(ips []string, start, end int, timeout time.Duration)  {

	var (
		wg 		= &sync.WaitGroup{}
		timeOut = time.Millisecond * 200
		unUse 	= map[string][]int{}
	)
	for port := start; port <= end; port++ {
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

// Dial
func isOpen(host string, port int, timeOut time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeOut)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false

}