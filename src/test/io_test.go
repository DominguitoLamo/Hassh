package test

import (
	"fmt"
	"os"
	"testing"
	"path/filepath"
)

func TestIo(t *testing.T) {
	f, _ := os.Create("hello")
	f.Write([]byte("hello"))
	
	stat, _ := f.Stat()
	b := make([]byte, stat.Size())
	f.Read(b)
	fmt.Println(string(b))
}

func TestPath(t *testing.T) {
	getExePath()
}

func getExePath() string {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exePath := filepath.Dir(ex)
    fmt.Println("exePath:", exePath)
    return exePath
}
