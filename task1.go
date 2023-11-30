package main

import (
	"fmt"
	"math"
)

const (
	epsilon = 0.00005
	a       = 2.0
	b       = 10.0
	order   = 4 // 4 is the order of Simpson's rule
)

func f(x float64) float64 {
	return 1 / (1 + math.Pow(x, 3))
}

// http://vingar.ho.ua/for_students/chm/integration_new_2.pdf - page 16
func simpsonsRule(a, b float64, n int) float64 {
	h := (b - a) / float64(n)
	sum := f(a) + f(b)

	for i := 1; i < n; i += 2 {
		sum += 4 * f(a+float64(i)*h)
	}

	for i := 2; i < n-1; i += 2 {
		sum += 2 * f(a+float64(i)*h)
	}

	return h / 3 * sum
}

func rungesRule(I, I2, order float64) float64 {
	return math.Abs(I-I2) / (math.Pow(2, order) - 1)
}

func main() {
	n := 2

	// Initial approximation
	I := simpsonsRule(a, b, n)
	fmt.Printf("Initial approximation using %d intervals: %.6f\n", n, I)

	// Refinement loop using Runge's rule
	for {
		n *= 2
		I2 := simpsonsRule(a, b, n)
		deviation := rungesRule(I, I2, order)

		fmt.Printf("Approximation using %d intervals: %.6f, Error: %.6f\n", n, I2, deviation)

		I = I2

		if deviation < epsilon {
			break
		}
	}

	fmt.Printf("Final approximation: %.6f\n", I)
}
