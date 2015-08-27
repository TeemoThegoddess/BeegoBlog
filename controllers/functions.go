package controllers

func GetStrLength(content string) bool {
	if len(content) < 30 {
		return false
	} else {
		return true
	}
}
