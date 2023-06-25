//go:build !android && !js

package giohyperlink

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"github.com/gioui-plugins/gio-plugins/hyperlink"
)

// NewConfigFromViewEvent creates a share.Config based on app.ViewEvent.
func NewConfigFromViewEvent(w *app.Window, evt app.ViewEvent) hyperlink.Config {
	r := hyperlink.Config{}
	UpdateConfigFromViewEvent(&r, w, evt)
	return r
}

func UpdateConfigFromViewEvent(config *hyperlink.Config, w *app.Window, evt app.ViewEvent) {}

func UpdateConfigFromStageEvent(config *hyperlink.Config, _ *app.Window, evt system.StageEvent) {}
