package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	// StatusActive defines that indicator is active.
	StatusActive Status = C.APP_INDICATOR_STATUS_ACTIVE
	// StatusPassive defines that indicator is not active.
	StatusPassive Status = C.APP_INDICATOR_STATUS_PASSIVE
	// StatusAttention defines that indicator wants a bag of attention.
	StatusAttention Status = C.APP_INDICATOR_STATUS_ATTENTION
)

// Status represents status of indicator.
type Status int
