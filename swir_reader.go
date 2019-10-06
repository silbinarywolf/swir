package swir

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Reader struct {
	expectedInputKeyCount int32
	// temporaryFrameData is reused each frame to provide a slice of
	// bools to the user
	temporaryFrameData [32]bool
	buf                bytes.Reader
}

func NewReader(expectedInputKeyCount int, data []byte) *Reader {
	w := &Reader{
		expectedInputKeyCount: int32(expectedInputKeyCount),
		buf:                   *bytes.NewReader(data),
	}
	w.readHeader()
	return w
}

func (r *Reader) readHeader() {
	var strSize byte
	if err := binary.Read(&r.buf, binary.LittleEndian, &strSize); err != nil {
		panic(err)
	}
	if strSize > versionSizeByteMax {
		panic(errInvalidVersionSize)
	}
	var versionData [versionSizeByteMax]byte
	for i := byte(0); i < strSize; i++ {
		var c byte
		if err := binary.Read(&r.buf, binary.LittleEndian, &c); err != nil {
			panic(err)
		}
		versionData[i] = c
	}
	version := string(versionData[:strSize])
	if version != formatVersion {
		panic("Unexpected version: '" + version + "', expected '" + formatVersion + "'")
	}
	var expectedInputKeyCount int32
	if err := binary.Read(&r.buf, binary.LittleEndian, &expectedInputKeyCount); err != nil {
		panic(err)
	}
	if expectedInputKeyCount != r.expectedInputKeyCount {
		panic("Key count no longer matches on files")
	}
}

// ReadFrame will return either a temporary slice of held down buttons or nil if there is no more data to read
func (r *Reader) ReadFrame() []bool {
	var eventCode byte
	if err := binary.Read(&r.buf, binary.LittleEndian, &eventCode); err != nil {
		if err == io.EOF {
			return nil
		}
		panic(err)
	}
	if eventCode != eventKeysDown {
		panic(errExpectedEventCode)
	}
	result := r.temporaryFrameData[:0]
	const chunkSize = 8
	keyCount := int(r.expectedInputKeyCount)
	for i := 0; i < keyCount; i += chunkSize {
		end := i + chunkSize
		if end > keyCount {
			end = keyCount
		}
		var b byte
		if err := binary.Read(&r.buf, binary.LittleEndian, &b); err != nil {
			panic(err)
		}
		var pos uint = 0
		for j := i; j < end; j++ {
			if b&(1<<pos) != 0 {
				result = append(result, true)
			} else {
				result = append(result, false)
			}
			pos++
		}
	}
	return result
}
