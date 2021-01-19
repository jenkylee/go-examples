package main

/**
 * 实现一个日志记录器满足:文件记录和数据库记录2种方式
 */

import "log"

type LogManager struct {
	Logging
}

func NewLogManager(logging Logging) *LogManager {
	return &LogManager{Logging: logging}
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

	manager := NewLogManager(FileLog)
	manager.Info()
	manager.Error()

	dbLog := &DbLogging{}

	manager.Logging = dbLog
	manager.Info()
	manager.Error()
}