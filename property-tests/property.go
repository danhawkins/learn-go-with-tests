package property_example

import "strings"

func ConvertToRoman(num int) string {
	var result strings.Builder

	for num > 0 {
		switch {
		case num > 9:
			result.WriteString("X")
			num -= 10
		case num > 8:
			result.WriteString("IX")
			num -= 9
		case num > 4:
			result.WriteString("V")
			num -= 5
		case num > 3:
			result.WriteString("IV")
			num -= 4
		default:
			result.WriteString("I")
			num--
		}
	}

	return result.String()
}
