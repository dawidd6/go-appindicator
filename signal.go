package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	SignalNewIcon           Signal = C.APP_INDICATOR_SIGNAL_NEW_ICON
	SignalNewAttentionIcon  Signal = C.APP_INDICATOR_SIGNAL_NEW_ATTENTION_ICON
	SignalNewStatus         Signal = C.APP_INDICATOR_SIGNAL_NEW_STATUS
	SignalNewLabel          Signal = C.APP_INDICATOR_SIGNAL_NEW_LABEL
	SignalConnectionChanged Signal = C.APP_INDICATOR_SIGNAL_CONNECTION_CHANGED
	SignalNewIconThemePath  Signal = C.APP_INDICATOR_SIGNAL_NEW_ICON_THEME_PATH
	SignalScrollEvent       Signal = C.APP_INDICATOR_SIGNAL_SCROLL_EVENT
)

type Signal string
