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
	vertical   bool = true
	changeRate int  = 1
)

func WillAppear(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.WillAppearEvent)

	setOrientationFeedback(plugin, p.Context)
	setChangeRateFeedback(plugin, p.Context)
}

func DialRotate(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.DialRotateEvent)

	if vertical {
		C.scroll(C.int(0), C.int(changeRate*p.Payload.Ticks))
	} else {
		C.scroll(C.int(changeRate*p.Payload.Ticks), C.int(0))
	}
}

func DialPress(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.DialPressEvent)

	if p.Payload.Pressed {
		if changeRate == 32 {
			setChangeRate(plugin, p.Context, 1)
		} else {
			setChangeRate(plugin, p.Context, changeRate*2)
		}
	}
}

func TouchTap(plugin *sdk.Plugin, event interface{}) {
	p := event.(*sdk.TouchTapEvent)

	if p.Payload.Hold {
		setChangeRate(plugin, p.Context, 1)
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

func setChangeRate(plugin *sdk.Plugin, c string, n int) error {
	changeRate = n

	setChangeRateFeedback(plugin, c)

	return nil
}

func setChangeRateFeedback(plugin *sdk.Plugin, c string) {
	var n int

	switch changeRate {
	case 1:
		n = 1
	case 2:
		n = 2
	case 4:
		n = 3
	case 8:
		n = 4
	case 16:
		n = 5
	case 32:
		n = 6
	default:
		n = 0
	}

	payload := map[string]string{"value": "Level " + strconv.Itoa(n)}

	plugin.SetFeedback(c, payload)
}
