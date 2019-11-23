package swir

import "errors"

const (
	eventNull byte = 0

	// 1 - 127 is reserved for custom user-code events

	eventKeysDown byte = 128 + iota
)

// fileType is the first bit of data in the file, this helps
// the reader check if its the correct file type.
// This should never change for the SWIRF file format, even 1000 years into the future.
// its the *one* constant.
const fileType = "SWIRF"

// formatVersion is followed after the header. This version string can be up to 64 bytes so that
// users can fork the versions / namespace however they want in a flexible way.
// ie. "jennysfork-1.2.0" or "samuels-swirf-1.0.0"
const formatVersion = "0.2.0"

// versionSizeByteMax is 64 bytes. We allow a fairly large version number so that
// this format can be forked easily and have many variations deep into the future
const versionSizeByteMax = 64

var (
	errInvalidFileType    = errors.New("Invalid file type, expected SWIRF format")
	errInvalidVersionSize = errors.New("Version string cannot be above 64 bytes")
	errInvalidKeySize     = errors.New("Incorrectly configured key count.")
	errExpectedEventCode  = errors.New("Data is malformed. Expected event code for key down.")
)

type errInvalidVersionString struct {
	badVersion string
}

func newErrInvalidVersionString(badVersion string) errInvalidVersionString {
	return errInvalidVersionString{
		badVersion: badVersion,
	}
}

func (err errInvalidVersionString) Error() string {
	return "Unexpected version: '" + err.badVersion + "', expected '" + formatVersion + "'"
}
