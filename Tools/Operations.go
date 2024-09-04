package tools

func PushTop(arrA, arrB *[]int) {
	*arrB = append(*arrB, (*arrA)[0])
	*arrA = (*arrA)[1:]
}

func SwapFirstTwo(arrA *[]int) {
	(*arrA)[0], (*arrA)[1] = (*arrA)[1], (*arrA)[0]
}

func Rotate(arrA *[]int) {
	*arrA = append((*arrA)[1:], (*arrA)[0])
}

func RevRotate(arrA *[]int) {
	*arrA = append((*arrA)[len(*arrA)-1:], (*arrA)[:len(*arrA)-1]...)
}
