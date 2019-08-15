//rotate matrix by 90 degree

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func generateMatrix(n int, s int) [][]int {
	mat := make([][]int, n)
	generator := rand.New(rand.NewSource(int64(s)))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i] = append(mat[i], generator.Intn(s))
		}
	}

	return mat

}

func printMat(mat [][]int) {
	n := len(mat)
	fmt.Println()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d\t", mat[i][j])
		}
		fmt.Println()
	}
}

func doRotate(first int, last int, mat [][]int, w *sync.WaitGroup) {
	// fmt.Println("rotating layer for ", first, last)
	for i := first; i < last; i++ {

		offset := i - first
		temp := mat[first][i]                            // top -> temp
		mat[first][i] = mat[last-offset][first]          // left -> top
		mat[last-offset][first] = mat[last][last-offset] // bottom -> left
		mat[last][last-offset] = mat[i][last]            // right -> bottom
		mat[i][last] = temp                              //top -> right
	}
	w.Done()
	// fmt.Println("rotation completed for layer ", first, last)
}

func rotateRight(mat [][]int) [][]int {
	n := len(mat)
	layers := n / 2

	var waitgroup sync.WaitGroup
	waitgroup.Add(layers)

	first := 0
	last := n - 1
	for l := 0; l < layers; l++ {
		if parallel > 0 {
			go doRotate(first, last, mat, &waitgroup)
		} else {
			doRotate(first, last, mat, &waitgroup)
		}

		first = first + 1
		last = last - 1
		if first >= last {
			break
		}

	}
	if parallel > 0 {
		waitgroup.Wait()
	}

	return mat
}

var parallel int

func main() {

	matrixSize, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Panic(err)
	}

	source, err := strconv.Atoi(os.Args[2])
	if err != nil {
		source = 97
	}

	parallel, err = strconv.Atoi(os.Args[3])
	if err != nil {
		parallel = 0
	}

	mat := generateMatrix(matrixSize, source)
	for i := 0; i < 70; i++ {
		rotateRight(mat)
		fmt.Printf("\nCompleted rotation for %d", i)

	}
	fmt.Println()

	// printMat(mat)
	// printMat()

}
