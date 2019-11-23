package swir

import (
	"bytes"
	"encoding/binary"
)

type Writer struct {
	expectedInputKeyCount int32
	buf                   bytes.Buffer
}

func NewWriter(expectedInputKeyCount int) *Writer {
	w := &Writer{
		expectedInputKeyCount: int32(expectedInputKeyCount),
	}
	w.writeHeader()
	return w
}

func (w *Writer) writeHeader() {
	// Write file format
	{
		fileTypeSize := byte(len(fileType))
		if err := binary.Write(&w.buf, binary.LittleEndian, fileTypeSize); err != nil {
			panic(err)
		}
		if err := binary.Write(&w.buf, binary.LittleEndian, []byte(fileType)); err != nil {
			panic(err)
		}
	}

	// Write version
	{
		strSize := byte(len(formatVersion))
		if strSize > versionSizeByteMax {
			panic(errInvalidVersionSize)
		}
		if err := binary.Write(&w.buf, binary.LittleEndian, strSize); err != nil {
			panic(err)
		}
		if err := binary.Write(&w.buf, binary.LittleEndian, []byte(formatVersion)); err != nil {
			panic(err)
		}
	}

	if err := binary.Write(&w.buf, binary.LittleEndian, w.expectedInputKeyCount); err != nil {
		panic(err)
	}
}

func (r *Writer) Bytes() []byte {
	return r.buf.Bytes()
}

func (r *Writer) String() string {
	return r.buf.String()
}

func (w *Writer) WriteFrame(keysDown []bool) {
	if int(w.expectedInputKeyCount) != len(keysDown) {
		panic(errInvalidKeySize)
	}
	if err := binary.Write(&w.buf, binary.LittleEndian, eventKeysDown); err != nil {
		panic(err)
	}
	const chunkSize = 8
	for i := 0; i < len(keysDown); i += chunkSize {
		end := i + chunkSize
		if end > len(keysDown) {
			end = len(keysDown)
		}
		var b byte
		var pos uint = 0
		for j := i; j < end; j++ {
			if keysDown[j] {
				b |= (1 << pos)
			}
			pos++
		}
		if err := binary.Write(&w.buf, binary.LittleEndian, b); err != nil {
			panic(err)
		}
	}
}
