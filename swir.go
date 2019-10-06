package swir

import "errors"

const (
	eventNull byte = 0

	// 1 - 127 is reserved for custom user-code events

	eventKeysDown byte = 128 + iota
)

const formatVersion = "0.1.0"

const versionSizeByteMax = 64

var (
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
