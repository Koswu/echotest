package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

var isServer bool

const TestCount = 100

func main() {
	ip := getIP()
	choice := getMenu()
	switch choice {
	case 1:
		checkTrans("tcp", ip, getPort())
	case 2:
		checkTrans("udp", ip, getPort())
	}
	fmt.Println("Press enter to exit...")
	fmt.Scanln()
}

func getIP() string {
	var inputIP string
	var ip net.IP
	for true {
		fmt.Print("Input the IP address to test:")
		fmt.Scanln(&inputIP)
		ip = net.ParseIP(inputIP)
		if ip != nil {
			break
		}
		fmt.Println("IP address format error")
	}
	return ip.String()
}

func getPort() int {
	var inputPort string
	var port int
	var err error
	for true {
		fmt.Print("Input the port number to test:")
		fmt.Scanln(&inputPort)
		port, err = strconv.Atoi(inputPort)
		if err == nil && port > 0 && port < 65535 {
			break
		}
		fmt.Println("Port number format error")
	}
	return port
}

func checkTrans(netMode string, ip string, port int) bool {
	connStr := ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial(netMode, connStr)
	if err != nil {
		fmt.Println("Failed to connect", connStr)
		return false
	}
	defer conn.Close()
	io := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for i := 0; i < TestCount; i++ {
		content := rand.Int()
		s := strconv.Itoa(content)
		_, err := io.WriteString(s + "\n")
		io.Flush()
		if err != nil {
			fmt.Println("Write remote error", err)
		}
		back, err := io.ReadString('\n')
		if err != nil {
			fmt.Println("Read remote error", err)
		}
		back = strings.TrimSpace(back)
		if back != s {
			fmt.Printf("%d/%d tests error, origin data：%s， get：%s\n", i+1, TestCount, s, back)
			return false
		} else {
			fmt.Printf("%d/%d tests success\n", i+1, TestCount)
		}
	}
	fmt.Println(netMode + "test fully success!!")
	return true
}

func getMenu() int {
	var input int
	for true {
		fmt.Println("(1) TCP test")
		fmt.Println("(2) UDP test")
		fmt.Print("Input your choice：")
		fmt.Scanf("%d\n", &input)
		for _, choice := range []int{1, 2} {
			if input == choice {
				return input
			}
		}
		fmt.Println("Input is illegal", input)
	}
	return 0
}
