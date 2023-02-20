PLUGIN_UUID   = com.andrewrynhard.knob
PLUGIN_BINARY = knob
PLUGIN_DIR    = $(PLUGIN_UUID).sdPlugin

all: build

.PHONY: build
build:
	mkdir -p $@/$(PLUGIN_DIR)
	cd src && GOARCH=amd64 CGO_ENABLED=1 go build -o ../$@/$(PLUGIN_BINARY)-amd64 .
	cd src && GOARCH=arm64 CGO_ENABLED=1 go build -o ../$@/$(PLUGIN_BINARY)-arm64 .
	cd $@/$(PLUGIN_DIR) && lipo -create -output knob ../$(PLUGIN_BINARY)-amd64 ../$(PLUGIN_BINARY)-arm64
	cp manifest.json $@/$(PLUGIN_DIR)
	cp -R assets/* $@/$(PLUGIN_DIR)
	cd $@ && zip -r $(PLUGIN_DIR).streamDeckPlugin $(PLUGIN_DIR)

install: build
	unzip $</$(PLUGIN_UUID).sdPlugin.streamDeckPlugin -d ~/Library/Application\ Support/com.elgato.StreamDeck/Plugins

uninstall:
	rm -rf ~/Library/Application\ Support/com.elgato.StreamDeck/Plugins/$(PLUGIN_UUID).sdPlugin

clean:
	rm -rf build
