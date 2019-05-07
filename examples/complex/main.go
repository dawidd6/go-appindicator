package main

import (
	"fmt"
	"github.com/dawidd6/go-appindicator"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"reflect"
)

const (
	label             = "label"
	guide             = "guide"
	id                = "indicator-id"
	iconName          = "network-transmit-receive"
	attentionIconName = "network-transmit"
	iconThemePath     = "/usr/share/icons"
	iconDesc          = "desc123"
	category          = appindicator.CategoryApplicationStatus
	status            = appindicator.StatusAttention
	orderingIndex     = uint(3)
	title             = "indicator-title"
)

func main() {
	gtk.Init(nil)
	defer gtk.Main()

	item, err := gtk.MenuItemNewWithLabel(label)
	if err != nil {
		log.Fatal(err)
	}

	_, err = item.Connect("activate", onActivate)
	if err != nil {
		log.Fatal(err)
	}

	menu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal(err)
	}
	defer menu.ShowAll()

	indicator := appindicator.New(id, iconName, category)
	check(indicator.GetId(), id, "id")
	check(indicator.GetCategory(), category, "category")

	indicator.SetStatus(status)
	check(indicator.GetStatus(), status, "status")

	indicator.SetAttentionIcon(attentionIconName)
	check(indicator.GetAttentionIcon(), attentionIconName, "attention_icon")

	indicator.SetAttentionIconFull(attentionIconName, iconDesc)
	check(indicator.GetAttentionIconDesc(), iconDesc, "attention_icon_desc")

	indicator.SetMenu(menu)
	check(indicator.GetMenu(), menu, "menu")

	indicator.SetIcon(iconName)
	check(indicator.GetIcon(), iconName, "icon")

	indicator.SetIconFull(iconName, iconDesc)
	check(indicator.GetIconDesc(), iconDesc, "icon_desc")

	indicator.SetLabel(label, guide)
	check(indicator.GetLabel(), label, "label")
	check(indicator.GetLabelGuide(), guide, "label_guide")

	indicator.SetIconThemePath(iconThemePath)
	check(indicator.GetIconThemePath(), iconThemePath, "icon_theme_path")

	indicator.SetOrderingIndex(orderingIndex)
	check(indicator.GetOrderingIndex(), orderingIndex, "ordering_index")

	indicator.SetSecondaryActivateTarget(item)
	check(indicator.GetSecondaryActivateTarget(), item, "secondary_activate_target")

	indicator.SetTitle(title)
	check(indicator.GetTitle(), title, "title")

	indicator.GetMenu().Add(item)
	indicator.GetSecondaryActivateTarget().SetLabel("changed")
}

func check(got, set interface{}, what string) {
	if !reflect.DeepEqual(got, set) {
		log.Println(what)
		log.Fatal("got: ", got, " | ", "set: ", set)
	}
}

func onActivate() {
	fmt.Println("activated")
}
