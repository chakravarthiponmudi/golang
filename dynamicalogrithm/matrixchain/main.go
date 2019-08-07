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

type pmat struct {
	i int
	j int
}

type _p = map[pmat]matrix

func generateMatrixDimensions(chainSize int) ([]matrix, _p) {
	matrixDimensions := make([]matrix, chainSize)
	p := make(map[pmat]matrix)
	generator := rand.New(rand.NewSource(92311))
	seed := chainSize
	matrixDimensions[0].row = generator.Intn(seed) + 1
	for i := range matrixDimensions {

		matrixDimensions[i].col = generator.Intn(seed) + 1
		if i < len(matrixDimensions)-1 {
			matrixDimensions[i+1].row = matrixDimensions[i].col
		}
	}
	for i, mat := range matrixDimensions {
		p[pmat{i, i}] = mat
	}
	return matrixDimensions, p
}

type dArray = []int

func calculateP(res bool, mat1 matrix, mat2 matrix) (int, *matrix) {
	cost := mat1.row * mat1.col * mat2.col
	if !res {
		mat := new(matrix)
		mat.row = mat1.row
		mat.col = mat2.col
		return cost, mat
	}

	return cost, new(matrix)

}

func matrixChainMultiplication(p _p, n int) []dArray {
	m := make([]dArray, n)
	s := make([]dArray, n)
	for i := range m {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			m[i][j] = math.MaxUint32
			for k := i; k < j; k++ {
				_, ok := p[pmat{i, j}]
				cost, mat := calculateP(ok, p[pmat{i, k}], p[pmat{k + 1, j}])
				q := m[i][k] + m[k+1][j] + cost
				if !ok {
					p[pmat{i, j}] = *mat
				}
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	// fmt.Println(m)
	// fmt.Println(s)
	// fmt.Println(p)
	return s

}

func printSolution(s []dArray, i int, j int) {
	if i == j {
		fmt.Printf("A%d", i)
	} else {
		fmt.Printf("(")
		printSolution(s, i, s[i][j])
		printSolution(s, s[i][j]+1, j)
		fmt.Printf(")")
	}

}

func main() {
	chainLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic("Error in Chain length", err)
	}

	mat, p := generateMatrixDimensions(chainLength)
	fmt.Println(p)

	s := matrixChainMultiplication(p, chainLength)
	fmt.Println(mat)
	printSolution(s, 0, chainLength-1)
}
