package pstring

import "strings"

// ConvertToPhoneNumber ...
func ConvertToPhoneNumber(phone string) string {
	// Return if there is no phone
	if phone == "" {
		return ""
	}

	// Split all spaces
	phone = strings.Replace(phone, " ", "", -1)

	// If first character is "0", replace to "+84", become "+84....."
	if string(phone[0]) == "0" {
		phone = strings.Replace(phone, "0", "+84", 1)
	}

	// If 2 first characters is "84"
	if phone[0:2] == "84" {
		// phone format is 84x xxx xxx, must add prefix "+84" before
		// else, phone format is 84 xxx xxx xxx add "+" before
		if len(phone) == 9 {
			phone = "+84" + phone
		} else {
			phone = "+" + phone
		}
	}

	return phone
}
