package appindicator

// #cgo pkg-config: appindicator3-0.1
// #include <libappindicator/app-indicator.h>
import "C"
import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"unsafe"
)

// Indicator represents a single indicator object.
//
// It holds only a pointer to C counterpart.
type Indicator struct {
	indicator *C.AppIndicator
}

// Object returns a pointer to glib.Object of indicator.
//
// Useful for connecting signals to indicator for example.
func (indicator *Indicator) Object() *glib.Object {
	return glib.Take(
		unsafe.Pointer(
			indicator.indicator,
		),
	)
}

/* Create stuff */

// New creates a fresh indicator.
func New(id, iconName string, category Category) *Indicator {
	return &Indicator{
		indicator: C.app_indicator_new(
			C.CString(id),
			C.CString(iconName),
			category.native(),
		),
	}
}

// NewWithPath creates a fresh indicator with custom icon theme path.
func NewWithPath(id, iconName string, category Category, iconThemePath string) *Indicator {
	return &Indicator{
		indicator: C.app_indicator_new_with_path(
			C.CString(id),
			C.CString(iconName),
			category.native(),
			C.CString(iconThemePath),
		),
	}
}

/* Set properties */

// SetStatus sets status of indicator.
func (indicator *Indicator) SetStatus(status Status) {
	C.app_indicator_set_status(
		indicator.indicator,
		status.native(),
	)
}

// SetAttentionIcon sets attention icon of indicator.
func (indicator *Indicator) SetAttentionIcon(iconName string) {
	C.app_indicator_set_attention_icon(
		indicator.indicator,
		C.CString(iconName),
	)
}

// SetAttentionIconFull sets attention icon of indicator with description.
func (indicator *Indicator) SetAttentionIconFull(iconName, iconDesc string) {
	C.app_indicator_set_attention_icon_full(
		indicator.indicator,
		C.CString(iconName),
		C.CString(iconDesc),
	)
}

// SetMenu sets gtk.Menu for indicator.
//
// Note that this is mandatory for indicator to show and work properly.
func (indicator *Indicator) SetMenu(menu *gtk.Menu) {
	C.app_indicator_set_menu(
		indicator.indicator,
		(*C.GtkMenu)(unsafe.Pointer(menu.Native())),
	)
}

// SetIcon sets icon of indicator.
func (indicator *Indicator) SetIcon(iconName string) {
	C.app_indicator_set_icon(
		indicator.indicator,
		C.CString(iconName),
	)
}

// SetIconFull sets icon of indicator with description.
func (indicator *Indicator) SetIconFull(iconName, iconDesc string) {
	C.app_indicator_set_icon_full(
		indicator.indicator,
		C.CString(iconName),
		C.CString(iconDesc),
	)
}

// SetLabel sets label of indicator.
//
// Second parameter "guide" is used to set maximum width for label.
// Don't know if it works. Feel free to pass empty string.
func (indicator *Indicator) SetLabel(label, guide string) {
	C.app_indicator_set_label(
		indicator.indicator,
		C.CString(label),
		C.CString(guide),
	)
}

// SetIconThemePath sets icon theme path of indicator.
func (indicator *Indicator) SetIconThemePath(iconThemePath string) {
	C.app_indicator_set_icon_theme_path(
		indicator.indicator,
		C.CString(iconThemePath),
	)
}

// SetOrderingIndex sets ordering index of indicator.
//
// It may or may not work.
func (indicator *Indicator) SetOrderingIndex(orderingIndex uint) {
	C.app_indicator_set_ordering_index(
		indicator.indicator,
		C.guint(orderingIndex),
	)
}

// SetSecondaryActivateTarget sets which menu item's callback
// will be activated when middle clicked indicator.
func (indicator *Indicator) SetSecondaryActivateTarget(menuItem *gtk.MenuItem) {
	C.app_indicator_set_secondary_activate_target(
		indicator.indicator,
		(*C.GtkWidget)(unsafe.Pointer(menuItem.Native())),
	)
}

// SetTitle sets title of indicator.
func (indicator *Indicator) SetTitle(title string) {
	C.app_indicator_set_title(
		indicator.indicator,
		C.CString(title),
	)
}

/* Get properties */

// GetId returns the ID of indicator specified in New() or NewWithPath().
func (indicator *Indicator) GetId() string {
	return C.GoString(
		C.app_indicator_get_id(
			indicator.indicator,
		),
	)
}

// GetCategory returns the category of indicator specified in New() or NewWithPath().
func (indicator *Indicator) GetCategory() Category {
	return Category(
		C.app_indicator_get_category(
			indicator.indicator,
		),
	)
}

// GetStatus returns current status of indicator.
func (indicator *Indicator) GetStatus() Status {
	return Status(
		C.app_indicator_get_status(
			indicator.indicator,
		),
	)
}

// GetIcon returns current icon of indicator.
func (indicator *Indicator) GetIcon() string {
	return C.GoString(
		C.app_indicator_get_icon(
			indicator.indicator,
		),
	)
}

// GetIconDesc returns current icon description of indicator.
func (indicator *Indicator) GetIconDesc() string {
	return C.GoString(
		C.app_indicator_get_icon_desc(
			indicator.indicator,
		),
	)
}

// GetIconThemePath returns current icon theme path of indicator.
func (indicator *Indicator) GetIconThemePath() string {
	return C.GoString(
		C.app_indicator_get_icon_theme_path(
			indicator.indicator,
		),
	)
}

// GetAttentionIcon returns current attention icon of indicator.
func (indicator *Indicator) GetAttentionIcon() string {
	return C.GoString(
		C.app_indicator_get_attention_icon(
			indicator.indicator,
		),
	)
}

// GetAttentionIconDesc returns current attention icon description of indicator.
func (indicator *Indicator) GetAttentionIconDesc() string {
	return C.GoString(
		C.app_indicator_get_attention_icon_desc(
			indicator.indicator,
		),
	)
}

// GetTitle returns current title of indicator.
func (indicator *Indicator) GetTitle() string {
	return C.GoString(
		C.app_indicator_get_title(
			indicator.indicator,
		),
	)
}

// GetMenu returns pointer to gtk.Menu.
func (indicator *Indicator) GetMenu() *gtk.Menu {
	object := glib.Take(unsafe.Pointer(
		C.app_indicator_get_menu(
			indicator.indicator,
		),
	))

	fn := gtk.WrapMap["GtkMenu"].(func(*glib.Object) *gtk.Menu)
	return fn(object)
}

// GetLabel returns current label of indicator.
func (indicator *Indicator) GetLabel() string {
	return C.GoString(
		C.app_indicator_get_label(
			indicator.indicator,
		),
	)
}

// GetLabelGuide returns current label guide of indicator.
func (indicator *Indicator) GetLabelGuide() string {
	return C.GoString(
		C.app_indicator_get_label_guide(
			indicator.indicator,
		),
	)
}

// GetOrderingIndex returns current ordering index of indicator.
func (indicator *Indicator) GetOrderingIndex() uint {
	return uint(
		C.app_indicator_get_ordering_index(
			indicator.indicator,
		),
	)
}

// GetSecondaryActivateTarget returns pointer to gtk.MenuItem which
// connected callback is activated when middle clicked indicator.
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

// BuildMenuFromDesktop builds whole menu with items from desktop file.
//
// To be honest I don't know how it works.
func (indicator *Indicator) BuildMenuFromDesktop(desktopFile, desktopProfile string) {
	C.app_indicator_build_menu_from_desktop(
		indicator.indicator,
		C.CString(desktopFile),
		C.CString(desktopProfile),
	)
}
