package utils

func splitString(str string) []string {
	var splited []string
	var strPiece string

	for _, char := range str {
		if char == ' ' || char == '.' || char == ';' || char == ',' {
			splited = append(splited, strPiece)

			strPiece = string(char)
			splited = append(splited, strPiece)

			strPiece = ""
		} else {
			strPiece += string(char)
		}
	}

	splited = append(splited, strPiece)

	return splited
}

func tableMatch(target string, str interface{}) bool {
	if stringsSlice, ok := str.([]string); ok { // Check if str is of type []string
		for _, s := range stringsSlice {
			if s == target {
				return true
			}
		}
	}
	return false
}

func NormalizeQuery(query string, tables interface{}) string {
	querySplited := splitString(query)
	var res string

	for _, word := range querySplited {
		if tableMatch(word, tables) {
			res += "\"" + word + "\""
		} else {
			res += word
		}
	}
	return res
}
