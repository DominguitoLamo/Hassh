package test

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func TestPath(t *testing.T) {
	getExePath()
}

func TestZip(t *testing.T) {
    fmt.Println("creating zip archive...")
    archive, err := os.Create("archive.zip")
    if err != nil {
        panic(err)
    }
    defer archive.Close()
    zipWriter := zip.NewWriter(archive)

    fmt.Println("writing first file to archive...")
    w1, err := zipWriter.Create("csv/test.csv")
    if err != nil {
        panic(err)
    }
    if _, err := io.WriteString(w1, "wewe,we,we,we,we,we"); err != nil {
        panic(err)
    }

    w2, err := zipWriter.Create("txt/test.txt")
    if err != nil {
        panic(err)
    }
    if _, err := io.WriteString(w2, "sdfdfdfdfdf"); err != nil {
        panic(err)
    }

    fmt.Println("closing zip archive...")
    zipWriter.Close()

	file, err := os.Open("archive.zip")
	if err != nil {
		 fmt.Println(err)
	   return
	 }
	defer file.Close()
 
	// Get the file size
	stat, err := file.Stat()
	if err != nil {
	   fmt.Println(err)
	   return
	}
 
	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
	   fmt.Println(err)
	   return
	}
	fmt.Println(bs)
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
