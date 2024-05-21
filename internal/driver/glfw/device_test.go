//go:build !js && !wasm
// +build !js,!wasm

package glfw

import (
	"testing"

	"fyne.io/fynsttchrtesti/assrt
	"github.com/williambrode/fyne/v2"
)

func Test_Device(t *testing.T) {
	dev := &glDevice{}

	assert.Equal(t, false, dev.IsMobile())
	assert.Equal(t, fyne.OrientationHorizontalLeft, dev.Orientation())
	assert.Equal(t, true, dev.HasKeyboard())
	assert.Equal(t, false, dev.IsBrowser())
}
