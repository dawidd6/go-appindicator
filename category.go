package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	CategoryApplicationStatus Category = C.APP_INDICATOR_CATEGORY_APPLICATION_STATUS
	CategoryCommunications    Category = C.APP_INDICATOR_CATEGORY_COMMUNICATIONS
	CategorySystemServices    Category = C.APP_INDICATOR_CATEGORY_SYSTEM_SERVICES
	CategoryHardware          Category = C.APP_INDICATOR_CATEGORY_HARDWARE
	CategoryOther             Category = C.APP_INDICATOR_CATEGORY_OTHER
)

type Category int

func (category Category) native() C.AppIndicatorCategory {
	return (C.AppIndicatorCategory)(category)
}
