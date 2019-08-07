package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type matrix struct {
	row int
	col int
}

func generateMatrixDimensions(chainSize int) []matrix {
	matrixDimensions := make([]matrix, chainSize)
	generator := rand.New(rand.NewSource(92311))
	seed := chainSize
	matrixDimensions[0].row = generator.Intn(seed) + 1
	for i := range matrixDimensions {

		matrixDimensions[i].col = generator.Intn(seed) + 1
		if i < len(matrixDimensions)-1 {
			matrixDimensions[i+1].row = matrixDimensions[i].col
		}
	}
	return matrixDimensions
}

type dArray = []int

func matrixChainMultiplication(p []matrix, n int) {
	m := make([]dArray, n)
	s := make([]dArray, n)
	for i := range m {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
	}

	for l := 2; l < n; l++ {
		for i := 0; i < n-l; i++ {
			j := i + l - 1
			m[i][j] = math.MaxUint32
			for k := i; k < j-1; k++ {
				q := m[i][k] + m[k+1][j] //+ ???
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

}

func main() {
	chainLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic("Error in Chain length", err)
	}

	p := generateMatrixDimensions(chainLength)
	fmt.Println(p)

	matrixChainMultiplication(p, chainLength)
}
