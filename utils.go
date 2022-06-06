package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetConnectionStatus() bool {
	fmt.Println("checking network connection:")
	duration, _ := time.ParseDuration("10s")
	conn, err := net.DialTimeout("ip:icmp", "1.1.1.1", duration)

	if err != nil {
		fmt.Print("error: ")
		fmt.Println(err)
		fmt.Println("done. \n ")
		return false
	}

	if conn != nil {
		fmt.Println("ping 1.1.1.1 .... ok")
	} else {
		fmt.Println("ping 1.1.1.1 .... fail")
	}
	conn.Close()

	connDNS, errDNS := net.DialTimeout("tcp", "ya.ru:80", duration)

	if errDNS != nil {
		fmt.Print("error: ")
		fmt.Println(errDNS)
		fmt.Println("done. \n ")
		return false
	}

	if connDNS != nil {
		fmt.Println("dns to ya.ru .... ok")
	} else {
		fmt.Println("dns to ya.ru .... fail")
	}

	fmt.Println("done. \n ")
	return true
}
