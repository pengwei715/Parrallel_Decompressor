package main

import (
	"bufio"
	"flag"
	"mpcs52060/proj3/src/util"
	"mpcs52060/proj3/src/pipeline"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"fmt"
)

//seqworker walks through every file in the repo and decompress it based on the type of the file
func seqworker(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		err = process(path)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

//process is the helper function for the seqworker
func process(path string) error {
	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	match := re.FindStringSubmatch(path)
	tail := match[3]
	switch tail {
	case ".gz":
		return util.GzipFiles(path)
	case ".bz2":
		return util.BzipFiles(path)
	case ".zlib":
		return util.ZlibFiles(path)
	}
	return nil
}

//parworker is the parallel version of the system
func parworker(root string) <-chan error {
	done := make(chan interface{})
	defer close(done)
	valchan, errchan := pipeline.Generator(root)
	chan1, chan2, chan3 := pipeline.Distribute(done, valchan)

	zlibErrBuffer := pipeline.ZlibDecompressor(done, chan3)
	bzErrBuffer := pipeline.BzDecompressor(done, chan2)
	gzErrBuffer := pipeline.GzDecompressor(done, chan1)

	totalerr := pipeline.CollectErr(done, zlibErrBuffer, bzErrBuffer, gzErrBuffer, errchan)
	for err := range totalerr {
		if err != nil {
			fmt.Printf("something wrong here %v:",err)
		}
	}
	return totalerr
}

func main() {
	if len(os.Args) == 1 {
		fi, _ := os.Stdin.Stat()
		if fi.Size() == 0 {
			fmt.Println("Usage : ./decompressor < textfile\n ./decompressor -p number <textfile")
			return
		}
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			root := scanner.Text()
			seqworker(root)
		}
		return
	}
	numCores := flag.Int("p", 1, "set the cores")
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		root := scanner.Text()
		parworker(root)
	}
	return
}
