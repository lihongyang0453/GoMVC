package RegexpHelper

import (
	"regexp"
)

func IsPhone(phone string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|6|7|8][0-9]\d{4,8})$`, phone); !m {
		return false
	}
	return true
}
func IsEmail(email string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}

func IsEnglish(str string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", str); !m {
		return false
	}
	return true
}
func IsChinese(str string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", str); !m {
		return false
	}
	return true
}

func IsIDCard(str string) bool {
	if len(str) == 18 {
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, str); !m {
			return false
		}
		return true
	} else if len(str) == 15 {
		if m, _ := regexp.MatchString(`^(\d{15})$`, str); !m {
			return false
		}
		return true
	} else {
		return false
	}
}
