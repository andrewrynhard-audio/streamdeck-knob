package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/andrewrynhard-audio/streamdeck-go-sdk/sdk"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	knobDirPath := path.Join(homeDir, ".knob")
	logFilePath := path.Join(knobDirPath, "com.andrewrynhard.knob.log")

	if err := os.MkdirAll(knobDirPath, 0700); err != nil {
		log.Fatal(err)
	}

	if err := os.RemoveAll(logFilePath); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	defer f.Close()

	plugin, err := sdk.NewPlugin()
	if err != nil {
		log.Fatal(err)
	}

	plugin.Handle(fmt.Sprintf("com.andrewrynhard.knob/%s", sdk.EventTouchTap), TouchTap)
	plugin.Handle(fmt.Sprintf("com.andrewrynhard.knob/%s", sdk.EventDialPress), DialPress)
	plugin.Handle(fmt.Sprintf("com.andrewrynhard.knob/%s", sdk.EventDialRotate), DialRotate)
	plugin.Handle(fmt.Sprintf("com.andrewrynhard.knob/%s", sdk.EventWillAppear), WillAppear)

	if err := plugin.Run(); err != nil {
		log.Fatal(err)
	}
}
