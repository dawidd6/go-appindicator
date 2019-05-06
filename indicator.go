package appindicator

// #cgo pkg-config: appindicator3-0.1
// #include <libappindicator/app-indicator.h>
import "C"
import (
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

type Indicator struct {
	indicator *C.AppIndicator
}

func New(id, iconName string, category Category) *Indicator {
	return &Indicator{
		indicator: C.app_indicator_new(
			stringToGCharPointer(id),
			stringToGCharPointer(iconName),
			category.Native(),
		),
	}
}

func NewWithPath(id, iconName string, category Category, iconThemePath string) *Indicator {
	return &Indicator{
		indicator: C.app_indicator_new_with_path(
			stringToGCharPointer(id),
			stringToGCharPointer(iconName),
			category.Native(),
			stringToGCharPointer(iconThemePath),
		),
	}
}

func (indicator *Indicator) SetStatus(status Status) {
	C.app_indicator_set_status(
		indicator.indicator,
		status.Native(),
	)
}

func (indicator *Indicator) SetAttentionIcon(iconName string) {
	C.app_indicator_set_attention_icon(
		indicator.indicator,
		stringToGCharPointer(iconName),
	)
}

func (indicator *Indicator) SetAttentionIconFull(iconName, iconDesc string) {
	C.app_indicator_set_attention_icon_full(
		indicator.indicator,
		stringToGCharPointer(iconName),
		stringToGCharPointer(iconDesc),
	)
}

func (indicator *Indicator) SetMenu(menu *gtk.Menu) {
	C.app_indicator_set_menu(
		indicator.indicator,
		(*C.GtkMenu)(unsafe.Pointer(menu.Native())),
	)
}

func (indicator *Indicator) SetIcon(iconName string) {
	C.app_indicator_set_icon(
		indicator.indicator,
		stringToGCharPointer(iconName),
	)
}

func (indicator *Indicator) SetIconFull(iconName, iconDesc string) {
	C.app_indicator_set_icon_full(
		indicator.indicator,
		stringToGCharPointer(iconName),
		stringToGCharPointer(iconDesc),
	)
}

func (indicator *Indicator) SetLabel(label string) {
	C.app_indicator_set_label(
		indicator.indicator,
		stringToGCharPointer(label),
		nil,
	)
}

func (indicator *Indicator) SetIconThemePath(iconThemePath string) {
	C.app_indicator_set_icon_theme_path(
		indicator.indicator,
		stringToGCharPointer(iconThemePath),
	)
}

func (indicator *Indicator) SetOrderingIndex(orderingIndex uint) {
	C.app_indicator_set_ordering_index(
		indicator.indicator,
		uIntToGUInt(orderingIndex),
	)
}

func (indicator *Indicator) SetTitle(title string) {
	C.app_indicator_set_title(
		indicator.indicator,
		stringToGCharPointer(title),
	)
}
