package enderecos

import "strings"

// TypeAddress /*
func TypeAddress(address string) string {
	typeValid := []string{"rua", "avenida", "estrada", "rodovia"}

	addressFormatted := strings.ToLower(address)

	firstWorldAddress := strings.Split(addressFormatted, " ")[0]

	addressValid := false

	for _, typeAddress := range typeValid {
		if typeAddress == firstWorldAddress {
			addressValid = true
		}
	}

	if addressValid {
		return strings.ToTitle(firstWorldAddress)
	}

	return "Invalid Type"
}
