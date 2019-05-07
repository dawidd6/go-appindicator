package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	SignalNewIcon           = C.APP_INDICATOR_SIGNAL_NEW_ICON
	SignalNewAttentionIcon  = C.APP_INDICATOR_SIGNAL_NEW_ATTENTION_ICON
	SignalNewStatus         = C.APP_INDICATOR_SIGNAL_NEW_STATUS
	SignalNewLabel          = C.APP_INDICATOR_SIGNAL_NEW_LABEL
	SignalConnectionChanged = C.APP_INDICATOR_SIGNAL_CONNECTION_CHANGED
	SignalNewIconThemePath  = C.APP_INDICATOR_SIGNAL_NEW_ICON_THEME_PATH
	SignalScrollEvent       = C.APP_INDICATOR_SIGNAL_SCROLL_EVENT
)
