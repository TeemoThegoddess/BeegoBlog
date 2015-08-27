package controllers

func GetStrLength(content string) string {
	if len(content) < 30 {
		return "false"
	} else {
		return "true"
	}
}
