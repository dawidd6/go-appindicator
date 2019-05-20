# go-appindicator

[![Build Status](https://api.cirrus-ci.com/github/dawidd6/go-appindicator.svg)](https://cirrus-ci.com/github/dawidd6/go-appindicator)
[![Go Report Card](https://goreportcard.com/badge/github.com/dawidd6/go-appindicator)](https://goreportcard.com/report/github.com/dawidd6/go-appindicator)
[![GoDoc](https://godoc.org/github.com/dawidd6/go-appindicator?status.svg)](https://godoc.org/github.com/dawidd6/go-appindicator)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

Go bindings for [libappindicator3](https://launchpad.net/libappindicator) C library.

Libappindicator is a library to allow applications to export a menu into the Unity Menu bar.
Based on KSNI it also works in KDE and will fallback to generic Systray support if none of those are available.

Also it works in:
 - Unity
 - GNOME Shell with [this extension](https://github.com/ubuntu/gnome-shell-extension-appindicator)
 - Budgie Desktop with [this applet](https://github.com/UbuntuBudgie/budgie-indicator-applet)
 - ...probably others that I haven't tested

This package aims to be interoperable with [Go gtk3 bindings](https://github.com/gotk3/gotk3).

`3` in the name means that it's GTK3 version.

 ## Dependencies
 
On Debian-based distributions:

```bash
apt install libappindicator3-dev libgtk-3-dev
```

And of course `go` with `cgo` is required.

## Building

Refer to [gotk3 wiki](https://github.com/gotk3/gotk3/wiki)

...or simply run [`build.sh`](./build.sh) script that will try to detect
currently installed version of GTK, pass along given `go build` flags
and execute it.

For example to build one of examples:

```bash
./build.sh -v examples/simple/main.go
```

## Examples

Examples are located in [examples](./examples) directory