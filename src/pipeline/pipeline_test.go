package pipeline

import (
	"fmt"
	"testing"
)

var root string = "../../test_data/"

func TestGenerator(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	valchan, errchan := Generator(root)
	for w := range valchan {
		fmt.Printf("%v\n", w)
	}
	for err := range errchan {
		if err != nil {
			t.Errorf("something wrong here %v:", err)
		}
	}
}

func TestDistribute(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	valchan, errchan := Generator(root)
	for w := range valchan {
		fmt.Printf("%v\n", w)
	}

	chan1, chan2, chan3 := Distribute(done, valchan)
	for i := range chan1 {
		fmt.Printf("%v\n", i)
	}
	for j := range chan2 {
		fmt.Printf("%v\n", j)
	}
	for k := range chan3 {
		fmt.Printf("%v\n", k)
	}

	for err := range errchan {
		if err != nil {
			t.Errorf("something wrong here %v:", err)
		}
	}
}

func TestGzDecompressor(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	valchan, _ := Generator(root)
	for w := range valchan {
		fmt.Printf("%v\n", w)
	}

	chan1, _, _ := Distribute(done, valchan)
	gzErrBuffer := GzDecompressor(done, chan1)

	for err := range gzErrBuffer {
		if err != nil {
			t.Errorf("something wrong here %v:", err)
		}
	}
}

func TestBzDecompressor(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	valchan, _ := Generator(root)
	for w := range valchan {
		fmt.Printf("%v\n", w)
	}
	_, chan2, _ := Distribute(done, valchan)
	bzErrBuffer := BzDecompressor(done, chan2)

	for err := range bzErrBuffer {
		if err != nil {
			t.Errorf("something wrong here %v:", err)
		}
	}
}

func TestZlibDecompressor(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	valchan, _ := Generator(root)
	for w := range valchan {
		fmt.Printf("%v\n", w)
	}
	_, _, chan3 := Distribute(done, valchan)
	zlibErrBuffer := ZlibDecompressor(done, chan3)

	for err := range zlibErrBuffer {
		if err != nil {
			t.Errorf("something wrong here %v:", err)
		}
	}
}

func TestAll(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	valchan, errchan := Generator(root)
	for w := range valchan {
		fmt.Printf("%v\n", w)
	}
	chan1, chan2, chan3 := Distribute(done, valchan)

	zlibErrBuffer := ZlibDecompressor(done, chan3)
	bzErrBuffer := BzDecompressor(done, chan2)
	gzErrBuffer := GzDecompressor(done, chan1)

	totalerr := CollectErr(done, zlibErrBuffer, bzErrBuffer, gzErrBuffer, errchan)
	for err := range totalerr {
		if err != nil {
			t.Errorf("something wrong here %v:", err)
		}
	}
}
