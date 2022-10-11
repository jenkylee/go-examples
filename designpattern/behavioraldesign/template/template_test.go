package template

import "testing"

func ExampleHTTPDownloder() {
	var downloader Downloader = NewHTTPDownloader()
	downloader.Download("http://example.com/abc.zip")
}

func ExampleFTPDownloader() {
	var downloader Downloader = NewFTPDownloader()
	downloader.Download("ftp://example.com/abc.zip")
}

func TestMain(m *testing.M) {
	ExampleHTTPDownloder()
	ExampleFTPDownloader()
}
