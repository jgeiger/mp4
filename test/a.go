
package main

import (
	"github.com/go-av/mp4"
	"fmt"
)

func write() {
	m, _ := mp4.Open("/Users/Xieran/Desktop/a.mp4")
	fmt.Println(m.W, m.H)
	pkts := m.ReadDur(20)
	m2, _ := mp4.Create("/tmp/out.mp4")
	fmt.Println(len(pkts))
	for _, p := range pkts {
		fmt.Println(p)
		m2.Write(p)
	}
	m2.Close()
}

func read() {
	m, err := mp4.Open("/tmp/out.mp4")
	mp4.LogLevel(0)
	fmt.Println(m.W)
}

func main() {
	read()
}

