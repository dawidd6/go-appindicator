package appindicator

// #include <libappindicator/app-indicator.h>
import "C"

const (
	// CategoryApplicationStatus defines that indicator's category is
	// an application status.
	CategoryApplicationStatus Category = C.APP_INDICATOR_CATEGORY_APPLICATION_STATUS
	// CategoryCommunications defines that indicator's category is
	// communications.
	CategoryCommunications Category = C.APP_INDICATOR_CATEGORY_COMMUNICATIONS
	// CategorySystemServices defines that indicator's category is
	// system services.
	CategorySystemServices Category = C.APP_INDICATOR_CATEGORY_SYSTEM_SERVICES
	// CategoryHardware defines that indicator's category is
	// hardware.
	CategoryHardware Category = C.APP_INDICATOR_CATEGORY_HARDWARE
	// CategoryOther defines that indicator's category is
	// something else.
	CategoryOther Category = C.APP_INDICATOR_CATEGORY_OTHER
)

// Category represents a category of indicator.
type Category int
