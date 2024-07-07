package main

import (
	"bufio"
	"log"
	"os"

	"github.com/konradmalik/kmls/rpc"
)

func main() {
	logger := getLogger("/tmp/kmls.log")

	logger.Println("Starting...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("no good file for a logger")
	}

	return log.New(logfile, "[kmls] ", log.Ldate|log.Ltime|log.Lshortfile)
}
