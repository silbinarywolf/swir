package swir

import (
	"io/ioutil"
	"testing"
)

const (
	// These constants represent the inputs
	// from https://github.com/silbinarywolf/swir-examples repo
	//Right = 0
	//Up    = 1
	//Left  = 2
	//Down  = 3
	Last = 4
)

func TestReadAndProcessValidSWIRFFile(t *testing.T) {
	const recordPath = "testdata/valid.swirf"
	recordData, err := ioutil.ReadFile(recordPath)
	if err != nil {
		t.Errorf("Failed to load %s: %s", recordPath, err)
		return
	}
	recordPlayer := NewReader(Last, recordData)
	for {
		keysDown := recordPlayer.ReadFrame()
		if keysDown == nil {
			break
		}
	}
}

func TestReadInvalidSWIRFFile(t *testing.T) {
	var err error

	const recordPath = "testdata/corrupt.swirf"
	recordData, err := ioutil.ReadFile(recordPath)
	if err != nil {
		t.Errorf("Failed to load %s: %s\n", recordPath, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(errInvalidVersionString)
			if !ok {
				t.Errorf("Unexpected error with panic: %s\n", r)
				return
			}
			// Hooray! We got the expected error.
		}
	}()
	_ = NewReader(Last, recordData)
	t.Errorf("Expected test to panic before reaching this line\n")
}
