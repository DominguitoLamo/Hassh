package test

import (
	"fmt"
	"os"
	"testing"
)

func TestIo(t *testing.T) {
	f, _ := os.Create("hello")
	f.Write([]byte("hello"))
	
	stat, _ := f.Stat()
	b := make([]byte, stat.Size())
	f.Read(b)
	fmt.Println(string(b))
}