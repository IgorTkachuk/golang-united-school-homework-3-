package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	l := len(input)
	for i := 0; i < l; i++ {
		result += input[i]
	}
	result /= float32(l)
	return
}
