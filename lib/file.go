package lib

func ValidateExtension(extension string) bool {
	switch extension {
	case "jpg", "jpeg", "gif", "png":
		return true
	default:
		return false
	}
}
