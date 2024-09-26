package logWriter

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	//logger := &Logger{
	//	Filename: "./test.log",
	//	MaxSize:  1,
	//	//MaxBackups: 2,
	//	MaxAge: 1,
	//	//BackupTimeFormat: "2006-01-02-15-04-05",
	//	BackupTimeFormat: "2006-01-02-15-04-05",
	//	LocalTime:        true,
	//	Compress:         true,
	//}
	logger := NewLogger("./test.log")
	logger.SetMaxSize(1)
	logger.SetMaxAge(1)
	logger.SetFormatTime("2006-01-02-15-04-05")
	logger.DisLocalTime()
	logger.DisCompress()
	log.SetOutput(logger)
	for {
		for i := 0; i < 10000; i++ {
			_, err := logger.Write([]byte("hello worldhello worldhello worldhello world\n"))
			if err != nil {
				fmt.Println(err)
			}
		}
		time.Sleep(time.Second * 2)
	}
}
