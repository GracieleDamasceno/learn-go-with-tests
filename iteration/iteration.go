package iteration

const repetitions = 5

func Iteration(character string) string {
	var generatedIteration string
	for i := 0; i < repetitions; i++ {
		generatedIteration += character
	}
	return generatedIteration
}
