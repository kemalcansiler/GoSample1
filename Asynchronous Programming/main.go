package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(8)
	go abcGen()

	go println("this comes first!")

	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
