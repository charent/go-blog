package mylog

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	logPath := "log"

	isExist, _ := PathExist(logPath)

	if !isExist {
		err := os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir %v fail, message: %v", logPath, err)
		}
	}

	file, err := os.OpenFile("log/run.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Fail to open logger file, message: ", err)
	}

	Info = log.New(io.MultiWriter(file, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}
