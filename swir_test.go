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

func TestReadInvalidVersionFile(t *testing.T) {
	var err error

	const recordPath = "testdata/invalid_version.swirf"
	recordData, err := ioutil.ReadFile(recordPath)
	if err != nil {
		t.Errorf("Failed to load %s: %s\n", recordPath, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			var err error
			err, _ = r.(errInvalidVersionString)
			if err == nil {
				t.Errorf("Expected to get version error with panic: %s\n", r)
				return
			}
			// Hooray! We got the expected error.
		}
	}()
	_ = NewReader(Last, recordData)
	t.Errorf("Expected test to panic before reaching this line\n")
}

func TestReadInvalidFileFormat(t *testing.T) {
	var err error

	const recordPath = "testdata/invalid_file_format.png"
	recordData, err := ioutil.ReadFile(recordPath)
	if err != nil {
		t.Errorf("Failed to load %s: %s\n", recordPath, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			var err error
			err, _ = r.(error)
			if err != errInvalidFileType {
				t.Errorf("Expected to get file format error with panic: %s\n", r)
				return
			}
			// Hooray! We got the expected error.
		}
	}()
	_ = NewReader(Last, recordData)
	t.Errorf("Expected test to panic before reaching this line\n")
}

func TestThreeInputsWriteThenReadInMemorySWIRFFile(t *testing.T) {
	const (
		Left  = 0
		Right = 1
		Jump  = 2
		Last  = 3
	)

	// Write recording information
	var recordData []byte
	{
		w := NewWriter(Last)
		for i := 0; i < 5; i++ {
			w.WriteFrame([]bool{
				Left:  false,
				Right: true,
				Jump:  false,
			})
		}
		recordData = w.Bytes()
		const expectedBytes = 26
		if len(recordData) != expectedBytes {
			t.Errorf("Expected recording file to be %d bytes but got %d bytes", expectedBytes, len(recordData))
		}
	}

	// Read recording information
	{
		r := NewReader(Last, recordData)
		for i := 0; i < 5; i++ {
			keysDown := r.ReadFrame()
			if keysDown[Left] {
				t.Errorf("Frame %d: Expected Left input to be false, not true", i)
			}
			if keysDown[Jump] {
				t.Errorf("Frame %d: Expected Jump input to be false, not true", i)
			}
			if !keysDown[Right] {
				t.Errorf("Frame %d: Expected Right input to be true, not false", i)
			}
		}
		keysDown := r.ReadFrame()
		if keysDown != nil {
			t.Errorf("Expected end of recording data, but there was more data")
			return
		}
	}
}
