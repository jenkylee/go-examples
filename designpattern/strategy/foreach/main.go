package main

/**
 * 遍历策略模式
 */

import (
	"log"
)

type LogManager struct {
	Log []Logging
}

func NewLogManager(logging ...Logging) *LogManager {
	return &LogManager{Log: logging}
}

func (l *LogManager) Info() {
	for _, logging := range l.Log {
		logging.Info()
	}
}

func (l *LogManager) Error() {
	for _, Logging := range l.Log {
		Logging.Error()
	}
}

type Logging interface {
	Info()
	Error()
}

type FileLogging struct {

}

func (f *FileLogging) Info() {
	log.Printf("%v", "文件记录 info")
}

func (f FileLogging) Error() {
	log.Printf("%v", "文件记录 error")
}

type DbLogging struct {

}

func (f *DbLogging) Info() {
	log.Printf("%v", "数据库记录 info")
}

func (f DbLogging) Error() {
	log.Printf("%v", "数据库记录 error")
}

func main()  {
	FileLog := &FileLogging{}
	dbLog := &DbLogging{}

	manager := NewLogManager(FileLog, dbLog)
	manager.Info()
	manager.Error()
}


