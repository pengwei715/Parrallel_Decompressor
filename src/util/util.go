//Package util offers functions to decompress three format compressed file
package util

import (
	"compress/bzip2"
	"compress/gzip"
	"compress/zlib"
	"io"
	"os"
	"regexp"
)

// GzipFiles process gzip files.
func GzipFiles(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	in, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	decompress(in, ".gz", path)
	return nil
}

// BzipFiles process bzip2 files.
func BzipFiles(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	in := bzip2.NewReader(file)
	decompress(in, ".bz2", path)
	return nil
}

// ZlibFiles process zlib files.
func ZlibFiles(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	in, err := zlib.NewReader(file)
	if err != nil {
		return err
	}
	decompress(in, ".zlib", path)
	return nil
}

//decompress is the helper function for the three functions above
func decompress(r io.Reader, ext, path string) error {
	re := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)
	match := re.FindStringSubmatch(path)
	filename := match[2]
	out, err := os.Create(match[1] + filename)
	if err != nil {
		out.Close()
		return err
	}
	_, err = io.Copy(out, r)
	if err != nil {
		out.Close()
		return err
	}
	out.Close()
	return nil
}
