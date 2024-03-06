package iteration

const defualtRepeatCount = 5

func Repeat(character string, count int) string {
	var repeated string

	if count == 0 {
		count = defualtRepeatCount
	}

	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
