package functions

func Graphic(splitArgs, banner []string) string {
	lines := 0
	result := ""

	for _, word := range splitArgs {
		if word == "" {
			lines = 1
		} else {
			lines = 8
		}
		for i := 1; i <= lines; i++ {
			for _, runWord := range word {
				index := (runWord-32)*9 + rune(i)
				result += banner[index]
			}
			result += "\n"
		}
	}
	return result
}
