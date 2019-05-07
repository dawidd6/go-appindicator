package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	// SignalNewIcon defines that indicator has new icon.
	SignalNewIcon = C.APP_INDICATOR_SIGNAL_NEW_ICON
	// SignalNewAttentionIcon defines that indicator has new attention icon.
	SignalNewAttentionIcon = C.APP_INDICATOR_SIGNAL_NEW_ATTENTION_ICON
	// SignalNewStatus defines that indicator has new status.
	SignalNewStatus = C.APP_INDICATOR_SIGNAL_NEW_STATUS
	// SignalNewLabel defines that indicator has new label.
	SignalNewLabel = C.APP_INDICATOR_SIGNAL_NEW_LABEL
	// SignalConnectionChanged defines that indicator's connection has changed.
	SignalConnectionChanged = C.APP_INDICATOR_SIGNAL_CONNECTION_CHANGED
	// SignalNewIconThemePath defines that indicator has new icon theme path.
	SignalNewIconThemePath = C.APP_INDICATOR_SIGNAL_NEW_ICON_THEME_PATH
	// SignalScrollEvent defines that user scrolled on indicator.
	SignalScrollEvent = C.APP_INDICATOR_SIGNAL_SCROLL_EVENT
)
