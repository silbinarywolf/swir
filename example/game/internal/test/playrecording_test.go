package test

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/swir/example/game/internal/game"
	"github.com/silbinarywolf/swir/example/game/internal/game/playback"
)

var (
	errPlaybackFinished = errors.New("Playback finished")
)

func testPlayRecordingUpdate(screen *ebiten.Image) error {
	isFinished := playback.RecordUpdate()
	err := game.Update(screen)
	if err != nil {
		return err
	}
	if isFinished {
		return errPlaybackFinished
	}
	return nil
}

func TestPlayRecording(t *testing.T) {
	onMainThread(func() {
		const recordPath = "record.swirf"
		recordData, err := ioutil.ReadFile(recordPath)
		if err != nil {
			t.Fatalf("Failed to load %s: %s", recordPath, err)
		}
		game.Init()
		playback.RecordInit(recordData)
		if err := ebiten.Run(testPlayRecordingUpdate, 320, 240, 2, "Hello world!"); err != nil {
			if err == errPlaybackFinished {
				return
			}
			t.Errorf("error: %s\n", err)
		}
	})
}

// -----------------------------------------
// Force tests to execute on the main thread
// so that OpenGL rendering works
// -----------------------------------------

var mainfunc = make(chan func())

// onMainThread will execute the given function on the main thread
func onMainThread(f func()) {
	done := make(chan struct{})
	mainfunc <- func() {
		f()
		close(done)
	}
	<-done
}

func TestMain(m *testing.M) {
	done := make(chan int, 1)
	go func() {
		done <- m.Run()
	}()
	for {
		runtime.Gosched()
		select {
		case f := <-mainfunc:
			f()
		case res := <-done:
			os.Exit(res)
		default:
			// don't block if no message
		}
	}
}
