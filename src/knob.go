package main

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>
void scroll(int x, int y) {
	CGEventRef event = CGEventCreateScrollWheelEvent(NULL, kCGScrollEventUnitPixel, 2, y, x);
	CGEventSetType(event, kCGEventScrollWheel);
	CGEventPost(kCGSessionEventTap, event);
	CFRelease(event);
}
*/
import "C"

import (
	"strconv"

	"github.com/andrewrynhard-audio/streamdeck-go-sdk/sdk"
)

var (
	vertical        bool = true
	dialRotateSpeed int  = 1
)

func WillAppear(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.WillAppearEvent)

	setOrientationFeedback(plugin, p.Context)
	setRotateMultiplierFeedback(plugin, p.Context)
}

func DialRotate(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.DialRotateEvent)

	if vertical {
		C.scroll(C.int(0), C.int(dialRotateSpeed*p.Payload.Ticks))
	} else {
		C.scroll(C.int(dialRotateSpeed*p.Payload.Ticks), C.int(0))
	}
}

func DialPress(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.DialPressEvent)

	if p.Payload.Pressed {
		if dialRotateSpeed == 32 {
			setRotateMulitplier(plugin, p.Context, 1)
		} else {
			setRotateMulitplier(plugin, p.Context, dialRotateSpeed*2)
		}
	}
}

func TouchTap(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.TouchTapEvent)

	if p.Payload.Hold {
		setRotateMulitplier(plugin, p.Context, 1)
	} else {
		setOrientation(plugin, p.Context)
	}
}

func setOrientation(plugin *sdk.Plugin, c string) error {
	vertical = !vertical

	setOrientationFeedback(plugin, c)

	return nil
}

func setOrientationFeedback(plugin *sdk.Plugin, c string) {
	var icon string

	if vertical {
		icon = "arrows-up-down.svg"
	} else {

		icon = "arrows-left-right.svg"
	}

	payload := map[string]string{"icon": icon}

	plugin.SetFeedback(c, payload)
}

func setRotateMulitplier(plugin *sdk.Plugin, c string, n int) error {
	dialRotateSpeed = n

	setRotateMultiplierFeedback(plugin, c)

	return nil
}

func setRotateMultiplierFeedback(plugin *sdk.Plugin, c string) {
	payload := map[string]string{"value": strconv.Itoa(dialRotateSpeed)}

	plugin.SetFeedback(c, payload)
}
