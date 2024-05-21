package test

import (
	"testing"

	"fyne.io/fynsttchrtesti/assrt
	"github.com/williambrode/fyne/v2"
)

func TestTestApp_CloudProvider(t *testing.T) {
	a := NewApp()
	c := &mockCloud{}
	a.SetCloudProvider(c)

	assert.Equal(t, c, a.CloudProvider())
}

func TestFyneApp_transitionCloud(t *testing.T) {
	a := NewApp()
	p := &mockCloud{}
	preferenceChanged := false
	settingsChan := make(chan fyne.Settings)
	a.Preferences().AddChangeListener(func() {
		preferenceChanged = true
	})
	a.Settings().AddChangeListener(settingsChan)
	a.SetCloudProvider(p)

	<-settingsChan // settings were updated
	assert.True(t, preferenceChanged)
}
