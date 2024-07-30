package addresses

import "strings"

// AddressType verifies if an address has a valid type and it returns it
func AddressType(address string) string { // Capitalized to allow to export
	validTypes := []string{"street", "avenue", "road", "highway", "freeway"}
	addressMinorCase := strings.ToLower(address)

	firstAdrressWord := strings.Split(addressMinorCase, " ")[0]

	addressHasAValidType := false
	for _, kind := range validTypes {
		if kind == firstAdrressWord {
			addressHasAValidType = true
		}
	}

	if addressHasAValidType {
		return strings.Title(firstAdrressWord) 
	}

	return "Invalid Type"
}