package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	TESTFILE string = "test.txt"
	OUTFILE  string = "out.txt"
)

type customWriter struct {

}

func (w *customWriter) Write(p []byte) (n int, err error) {
	writeToFile(string(p), OUTFILE)
	return len(p), nil
}

func main() {
	writer := &customWriter{}
	go infinitelyWriteToFile()

	cmd := exec.Command("tail", "-f", TESTFILE)
	//cmd := exec.Command("ls", "-la")
	//cmd := exec.Command("git", "clone", "https://github.com/sophiewakely/sophiewakely", "-v")
	cmd.Stdout = io.MultiWriter(os.Stdout, writer)
	cmd.Stderr = io.MultiWriter(os.Stdout, writer)
	cmd.Run()
}

func infinitelyWriteToFile() {
	for {
		t := time.Now()
		writeToFile(t.Format(time.RFC3339) + "\n", TESTFILE)
		time.Sleep(time.Second *2)
	}
}

func writeToFile(s string, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(s); err != nil {
		log.Println(err)
	}
}

