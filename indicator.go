package appindicator

// #cgo pkg-config: appindicator3-0.1
// #include <libappindicator/app-indicator.h>
import "C"
import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

type Indicator struct {
	indicator *C.AppIndicator
}

/* Create stuff */

func New(id, iconName string, category Category) *Indicator {
	return &Indicator{
		indicator: C.app_indicator_new(
			C.CString(id),
			C.CString(iconName),
			category.Native(),
		),
	}
}

func NewWithPath(id, iconName string, category Category, iconThemePath string) *Indicator {
	return &Indicator{
		indicator: C.app_indicator_new_with_path(
			C.CString(id),
			C.CString(iconName),
			category.Native(),
			C.CString(iconThemePath),
		),
	}
}

/* Set properties */

func (indicator *Indicator) SetStatus(status Status) {
	C.app_indicator_set_status(
		indicator.indicator,
		status.Native(),
	)
}

func (indicator *Indicator) SetAttentionIcon(iconName string) {
	C.app_indicator_set_attention_icon(
		indicator.indicator,
		C.CString(iconName),
	)
}

func (indicator *Indicator) SetAttentionIconFull(iconName, iconDesc string) {
	C.app_indicator_set_attention_icon_full(
		indicator.indicator,
		C.CString(iconName),
		C.CString(iconDesc),
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
		C.CString(iconName),
	)
}

func (indicator *Indicator) SetIconFull(iconName, iconDesc string) {
	C.app_indicator_set_icon_full(
		indicator.indicator,
		C.CString(iconName),
		C.CString(iconDesc),
	)
}

func (indicator *Indicator) SetLabel(label, guide string) {
	C.app_indicator_set_label(
		indicator.indicator,
		C.CString(label),
		C.CString(guide),
	)
}

func (indicator *Indicator) SetIconThemePath(iconThemePath string) {
	C.app_indicator_set_icon_theme_path(
		indicator.indicator,
		C.CString(iconThemePath),
	)
}

func (indicator *Indicator) SetOrderingIndex(orderingIndex uint) {
	C.app_indicator_set_ordering_index(
		indicator.indicator,
		C.guint(orderingIndex),
	)
}

func (indicator *Indicator) SetSecondaryActivateTarget(menuItem *gtk.MenuItem) {
	C.app_indicator_set_secondary_activate_target(
		indicator.indicator,
		(*C.GtkWidget)(unsafe.Pointer(menuItem.Native())),
	)
}

func (indicator *Indicator) SetTitle(title string) {
	C.app_indicator_set_title(
		indicator.indicator,
		C.CString(title),
	)
}

/* Get properties */

func (indicator *Indicator) GetId() string {
	return C.GoString(
		C.app_indicator_get_id(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetCategory() Category {
	return Category(
		C.app_indicator_get_category(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetStatus() Status {
	return Status(
		C.app_indicator_get_status(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetIcon() string {
	return C.GoString(
		C.app_indicator_get_icon(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetIconDesc() string {
	return C.GoString(
		C.app_indicator_get_icon_desc(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetIconThemePath() string {
	return C.GoString(
		C.app_indicator_get_icon_theme_path(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetAttentionIcon() string {
	return C.GoString(
		C.app_indicator_get_attention_icon(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetAttentionIconDesc() string {
	return C.GoString(
		C.app_indicator_get_attention_icon_desc(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetTitle() string {
	return C.GoString(
		C.app_indicator_get_title(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetMenu() *gtk.Menu {
	object := glib.Take(unsafe.Pointer(
		C.app_indicator_get_menu(
			indicator.indicator,
		),
	))

	fn := gtk.WrapMap["GtkMenu"].(func(*glib.Object) *gtk.Menu)
	return fn(object)
}

func (indicator *Indicator) GetLabel() string {
	return C.GoString(
		C.app_indicator_get_label(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetLabelGuide() string {
	return C.GoString(
		C.app_indicator_get_label_guide(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetOrderingIndex() uint {
	return uint(
		C.app_indicator_get_ordering_index(
			indicator.indicator,
		),
	)
}

func (indicator *Indicator) GetSecondaryActivateTarget() *gtk.MenuItem {
	object := glib.Take(unsafe.Pointer(
		C.app_indicator_get_secondary_activate_target(
			indicator.indicator,
		),
	))

	fn := gtk.WrapMap["GtkMenuItem"].(func(*glib.Object) *gtk.MenuItem)
	return fn(object)
}

/* Helpers */

func (indicator *Indicator) BuildMenuFromDesktop(desktopFile, desktopProfile string) {
	C.app_indicator_build_menu_from_desktop(
		indicator.indicator,
		C.CString(desktopFile),
		C.CString(desktopProfile),
	)
}
