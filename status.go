package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	StatusActive    Status = C.APP_INDICATOR_STATUS_ACTIVE
	StatusPassive   Status = C.APP_INDICATOR_STATUS_PASSIVE
	StatusAttention Status = C.APP_INDICATOR_STATUS_ATTENTION
)

type Status int

func (status Status) native() C.AppIndicatorStatus {
	return (C.AppIndicatorStatus)(status)
}
