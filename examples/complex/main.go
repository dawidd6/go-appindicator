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
	// GTK init and loop main at the end.
	gtk.Init(nil)
	defer gtk.Main()

	// Example MenuItem attached to indicator.
	item, err := gtk.MenuItemNewWithLabel(label)
	if err != nil {
		log.Fatal(err)
	}

	// Mandatory menu for indicator.
	//
	// ShowAll() will run before gtk.Main().
	menu, err := gtk.MenuNew()
	if err != nil {
		log.Fatal(err)
	}
	defer menu.ShowAll()

	// Indicator creation and checking functionality.
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

	// Fancy adding item to menu through indicator.
	//
	// It could be done by just calling menu.Add(item),
	// but I want to check here if GetMenu() works correctly.
	indicator.GetMenu().Add(item)
	// Same as above.
	indicator.GetSecondaryActivateTarget().SetLabel(label + "-changed")

	// Connect callback to item.
	_, err = item.Connect("activate", func() {
		indicator.SetLabel(label+"-changed", "")
	})
	if err != nil {
		log.Fatal(err)
	}

	// Connect callback to indicator
	_, err = indicator.Object().Connect(appindicator.SignalNewLabel, func() {
		fmt.Println("NewLabel: ", indicator.GetLabel())
	})
	if err != nil {
		log.Fatal(err)
	}
}

func check(got, set interface{}, what string) {
	if !reflect.DeepEqual(got, set) {
		log.Println(what)
		log.Fatal("got: ", got, " | ", "set: ", set)
	}
}
