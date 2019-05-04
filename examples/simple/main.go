package main

import (
	"fmt"
	"github.com/dawidd6/go-appindicator"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	gtk.Init(nil)
	defer gtk.Main()

	menu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal(err)
	}

	item, err := gtk.MenuItemNewWithLabel("item-label")
	if err != nil {
		log.Fatal(err)
	}

	_, err = item.Connect("activate", onActivate)
	if err != nil {
		log.Fatal(err)
	}

	indicator := appindicator.New("indicator-xyz", "network-transmit-receive", appindicator.CategoryApplicationStatus)
	indicator.SetTitle("indi-title")
	indicator.SetLabel("indi-label")
	indicator.SetStatus(appindicator.StatusActive)
	indicator.SetMenu(menu)

	menu.Add(item)
	menu.ShowAll()
}

func onActivate() {
	fmt.Println("activated")
}
