# SilbinaryWolf's Input Recorder for Go

[![Build Status](https://travis-ci.com/silbinarywolf/swir.svg?branch=master)](https://travis-ci.com/silbinarywolf/swir)
[![Documentation](https://godoc.org/github.com/silbinarywolf/swir?status.svg)](https://godoc.org/github.com/silbinarywolf/swir)
[![Report Card](https://goreportcard.com/badge/github.com/silbinarywolf/swir)](https://goreportcard.com/report/github.com/silbinarywolf/swir)

**WARNING: Backwards compatibility on the file format is currently not guaranteed. This will change once the repository is tagged at v1.0.0 and I complete my game. The reason for this is because I want to ensure it has key functionality I require before calling it stable at v1.0.0.**

Basic input recording system that writes what keys are held down per frame to a stream of bytes. Each key that can be held down takes up 1-bit of space per frame. This stream of bytes can used to play a users gameplay back for purposes such as integration testing, replays and more.

## Quick Start

1) Install
```
go get github.com/silbinarywolf/swir
```

2) Watch the existing recording test
```
go test ./...
```

3) Look at the code in example/game to see how this system works and can be tied together.

## Requirements

* Golang 1.13+

## Documentation

* [Documentation](https://godoc.org/github.com/silbinarywolf/swir)
* [License](LICENSE.md)

## Features that may be considered in the future
I have ideas and features that I'd like to look into adding in the future, however, for now this library achieves the bare-minimum of what I need.

- Recording Random Number Seeds
	- Not sure if this needs its own built-in API or if this is something a user should handle themselves. The Retro City Rampage talk on [automated testing](https://www.youtube.com/watch?v=W20t1zCZv8M) mentions that to keep tests working over time, you will want to decouple random seeds for gameplay and visual effects that don't actually affect gameplay. 
- Expose API to allow a user to write into the stream with custom events or behaviour (ie. level changed, log characters X/Y, etc)
	- ie. `WriteEvent(event EventID, data []byte)`, where `EventID` is a constant between 1-127. (127-255 is reserved for internal events like keypresses)
- Recording Mouse Movement / Click
  - Ideally this would use deltas between frames to try and keep the recording file as small as possible

## Credits

* [Brian Provinciano](https://www.youtube.com/watch?v=W20t1zCZv8M) for inspiration from their GDC Talk, "Automated Testing and Instant Replays in Retro City Rampage"
