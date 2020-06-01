package util

import (
	"testing"
)

var prePath string = "../../test_data/"

func TestBzipFiles(t *testing.T) {
	filePath := prePath + "file1.bz2"
	err := BzipFiles(filePath)
	if err != nil {
		t.Errorf("something wrong here %v", err)
	}
}

func TestGzipFiles(t *testing.T) {
	filePath := prePath + "file1.gz"
	err := GzipFiles(filePath)
	if err != nil {
		t.Errorf("something wrong here %v", err)
	}
}

func TestZlibFiles(t *testing.T) {
	filePath := prePath + "file1.zlib"
	err := ZlibFiles(filePath)
	if err != nil {
		t.Errorf("something wrong here %v", err)
	}
}
