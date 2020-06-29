package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	logFilePath := "logs.txt"
	noOfGoRoutines := 10

	if len(os.Args) > 1 {
		logFilePath = os.Args[1]
	}

	if len(os.Args) > 2 {
		noOfGoRoutines, _ = strconv.Atoi(os.Args[2])
	}

	var file, err = os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	var wg sync.WaitGroup
	for i := 0; i < noOfGoRoutines; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, identifier int) {
			defer wg.Done()
			for {
				log.WithFields(
					log.Fields{
						"goid":       getGID(),
						"identifier": identifier,
					},
				).Info("Something happened")

				log.WithFields(
					log.Fields{
						"goid":       getGID(),
						"identifier": identifier,
					},
				).Debug("Debugging something")

				time.Sleep(1 * time.Second)
			}
		}(&wg, i)
	}

	wg.Wait()
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
