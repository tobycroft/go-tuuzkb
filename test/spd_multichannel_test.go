package test

import (
	"testing"
)

var testChannel1 = make(chan []byte, 1)
var testChannel2 = make(chan []byte, 0)
var testChannel3 = make(chan []byte, 0)
var testChannel4 = make(chan []byte, 0)

func BenchmarkMulChannel(b *testing.B) {
	go run()
	go run2()
	go run3()
	b.StartTimer()
	c := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		testChannel1 <- c
		<-testChannel4
	}
	b.StopTimer()
}

func run() {
	for c := range testChannel1 {
		testChannel2 <- c
	}
}
func run2() {
	for c := range testChannel2 {
		testChannel3 <- c
	}
}
func run3() {
	for c := range testChannel3 {
		testChannel4 <- c
	}
}
