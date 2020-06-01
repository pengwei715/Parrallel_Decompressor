//Package pipeline contains pipeline pattern to process all the compressed files
package pipeline

import (
	"mpcs52060/proj3/src/util"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

//Generator walk throught every file in the root and push the file name into one channel, and push the err information into the err channel
func Generator(root string) (<-chan string, <-chan error) {
	pathBuffer := make(chan string)
	errBuffer := make(chan error, 1)
	go func() {
		defer close(pathBuffer)
		defer close(errBuffer)
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			pathBuffer <- path
			return nil
		})
		if err != nil {
			errBuffer <- err
		}
	}()
	return pathBuffer, errBuffer
}

//Distribute use fanout pattern to get the file from the total channel, based on the type of file distribute it into three different channels.
func Distribute(done <-chan interface{}, total <-chan string) (<-chan string, <-chan string, <-chan string) {
	gzBuffer := make(chan string, 1)
	bzBuffer := make(chan string, 1)
	zlibBuffer := make(chan string, 1)
	go func() {
		defer close(gzBuffer)
		defer close(bzBuffer)
		defer close(zlibBuffer)
		//use regex to get the file's type
		re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
		for path := range total {
			select {
			case <-done:
				return
			default:
				match := re.FindStringSubmatch(path)
				tail := match[3]
				switch tail {
				case ".gz":
					gzBuffer <- path
				case ".bz2":
					bzBuffer <- path
				case ".zlib":
					zlibBuffer <- path
				}
			}
		}
	}()
	return gzBuffer, bzBuffer, zlibBuffer
}

//GzDecompressor decompress the .gz files and return a err channel
func GzDecompressor(done <-chan interface{}, gzBuffer <-chan string) <-chan error {
	gzErrorBuffer := make(chan error, 1)
	go func() {
		defer close(gzErrorBuffer)
		for path := range gzBuffer {
			select {
			case <-done:
				return
			default:
				err := util.GzipFiles(path)
				if err != nil {
					gzErrorBuffer <- err
				}
			}
		}
	}()
	return gzErrorBuffer
}

//BzDecompressor decompress the .bz2 files and return a err channel
func BzDecompressor(done <-chan interface{}, bzBuffer <-chan string) <-chan error {
	bzErrorBuffer := make(chan error, 1)
	go func() {
		defer close(bzErrorBuffer)
		for path := range bzBuffer {
			select {
			case <-done:
				return
			default:
				err := util.BzipFiles(path)
				if err != nil {
					bzErrorBuffer <- err
				}
			}
		}
	}()
	return bzErrorBuffer
}

//ZlibDecompressor decompress the .zlib files and return a err channel
func ZlibDecompressor(done <-chan interface{}, ZlibBuffer <-chan string) <-chan error {
	ZlibErrorBuffer := make(chan error, 1)
	go func() {
		defer close(ZlibErrorBuffer)
		for path := range ZlibBuffer {
			select {
			case <-done:
				return
			default:
				err := util.ZlibFiles(path)
				if err != nil {
					ZlibErrorBuffer <- err
				}
			}
		}
	}()
	return ZlibErrorBuffer
}

//CollectErr use fanin pattern here to collect all the error information
func CollectErr(done <-chan interface{}, errorChans ...<-chan error) <-chan error {
	allErrorBuffer := make(chan error, 1)
	var wg sync.WaitGroup
	output := func(c <-chan error) {
		defer wg.Done()
		for err := range c {
			select {
			case allErrorBuffer <- err:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(errorChans))
	for _, c := range errorChans {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(allErrorBuffer)
	}()
	return allErrorBuffer
}
