package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"

	log "github.com/sirupsen/logrus"
)

func main() {
	//log.SetFormatter(&log.JSONFormatter{})
	var file, err = os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}
	log.SetOutput(file)

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, identifier int) {
			defer wg.Done()
			log.WithFields(
				log.Fields{
					"goid":       getGID(),
					"identifier": identifier,
				},
			).Info("Something happened")
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
