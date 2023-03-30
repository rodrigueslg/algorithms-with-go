package module01

func GCD(a, b int) int {
	// var minor = a
	// if b < a {
	// 	minor = b
	// }

	// for i := minor; i > 1; i-- {
	// 	modA := a % i
	// 	modB := b % i
	// 	if modA == 0 && modB == 0 {
	// 		return i
	// 	}
	// }

	// return 1

	return GCDEuclideanAlgorithm(a, b)
}

func GCDEuclideanAlgorithm(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GCD Euclidean Algorithm
