# https://github.com/gioui-plugins/gio-plugins
# https://github.com/inkeliz

# overide his one...
GOGIO_HASH=1151eac07d84

print:
	@echo ""
	@echo "GOGIO_HASH:   $(GOGIO_HASH)"
	@echo ""


dep:

# The go.mods have the following so we need to get the right versions of gio 
# replace gioui.org => ../../../gio
# replace github.com/gioui-plugins/gio-plugins => ../../

	# pull the gio source code and put it one level above this repo.
	# gioui.org@v0.6.1
	# v0.6.1-0.20240607083507-1151eac07d84
	# 1151eac07d84
	# https://github.com/gioui/gio/releases/tag/v0.7.1
	cd ./.. && git clone https://github.com/gioui/gio -b $(GOGIO_HASH)

	# The sub make files assume gio cmd is installed
	# https://github.com/gioui/gio-cmd/blob/main/go.mod
	go install gioui.org/cmd/gogio@latest


dep-del:
	rm -rf ./../gio
	rm -f $(shell command -v gogio)


auth-run:
	# does not work...
	cd auth/demo && go mod tidy
	cd auth/demo && go run .

explorer-run:
	# works on mac
	cd explorer/demo && go mod tidy
	cd explorer/demo && go run .

hyperlink-run:
	# works on mac
	cd hyperlink/demo && go mod tidy
	cd hyperlink/demo && go run .

safedata-run-standalone:
	# works on mac really well.
	cd safedata/demo/gio-with-standalone && go mod tidy
	cd safedata/demo/gio-with-standalone && go run .
safedata-run:
	# works on mac really well.
	cd safedata/demo/gio-with-gioplugin && go mod tidy
	cd safedata/demo/gio-with-gioplugin && go run .

share-run:
	# works on mac really well and shares to all the normal stuff.
	cd share/demo/ && go mod tidy
	cd share/demo/ && go run .

webviewer-run:
	# opens but cant load https://www.google.com
	cd webviewer/demo/ && go mod tidy
	cd webviewer/demo/ && go run .
webviewer-build-macos:
	# works
	cd webviewer/demo/ && go get gioui.org/cmd/gogio
	cd webviewer/demo/ && $(MAKE) macos
webviewer-build-run-macos:
	# works
	cd webviewer/demo && open ./macos_arm64.app




