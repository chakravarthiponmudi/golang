package main

import (
	"fmt"
	"sync"
)

type grid int16
type blockLocation int8

const gridsize = 9
const maxCellValue = 9
const blocksize = 3
const cubesize = 81

var cellvalues = []int8{1, 2, 3, 4, 5, 6, 7, 8, 9}

func getRow(index grid) grid {
	lIndex := index + 1 //offset adjuster
	row := lIndex / gridsize
	if lIndex%gridsize > 0 {
		row++
	}
	return row
}

func getColumn(index grid) grid {

	column := (index + 1) % gridsize
	if column == 0 {
		column = gridsize
	}
	return column
}

func isDiagonal(row grid, column grid) bool {
	return row == column
}

func getIndexFromRowAndColumn(row grid, column grid) grid {
	return (row-1)*gridsize + column - 1
}

var indexToBlockMap map[grid]blockLocation
var blockToIndexes map[blockLocation][]grid
var rowToIndexMap map[grid][]grid
var columnToIndexMap map[grid][]grid

func getBlockFromRowAndColumn(row grid, column grid) blockLocation {
	x := row / blocksize
	if row%blocksize > 0 {
		x++
	}
	y := column / blocksize
	if column%blocksize > 0 {
		y++
	}
	return blockLocation((x-1)*blocksize + y)

}

func determineBlock(index grid) {
	row := getRow(index)
	column := getColumn(index)
	block := getBlockFromRowAndColumn(row, column)
	indexToBlockMap[index] = block
	blockToIndexes[block] = append(blockToIndexes[block], index)
	rowToIndexMap[row] = append(rowToIndexMap[row], index)
	columnToIndexMap[column] = append(columnToIndexMap[column], index)
}

/**
	Bounding function:_
		Rules :-
			1) no number is repeated betwen 1 - gridsize in a row
			2) no number is repeated betwen 1 - gridsize in a column

**/
func BoundingFunction(elem node, solution []node) bool {
	return rowCheck(elem, solution) &&
		columnCheck(elem, solution) &&
		blockCheck(elem, solution)
}

func rowCheck(elem node, solution []node) bool {
	row := getRow(elem.index)
	for _, v := range rowToIndexMap[row] {
		if v != elem.index && solution[v].value == elem.value {
			return false
		}
	}
	return true
}

func columnCheck(elem node, solution []node) bool {
	column := getColumn(elem.index)
	for _, v := range columnToIndexMap[column] {
		if v != elem.index && solution[v].value == elem.value {
			return false
		}
	}
	return true
}

func blockCheck(elem node, solution []node) bool {
	block := indexToBlockMap[elem.index]
	for _, v := range blockToIndexes[block] {
		if v != elem.index && solution[v].value == elem.value {
			return false
		}
	}
	return true
}

type node struct {
	index grid
	block blockLocation
	fixed bool
	value int8
}

type solution struct {
	elems []node
}

func printSolution(s *solution) {
	fmt.Println()
	for i, n := range s.elems {
		if i%gridsize == 0 {
			fmt.Println()
		}
		if n.value == -1 {
			fmt.Print("\t_")
		} else {
			fmt.Printf("\t%d", n.value)
		}
	}
	fmt.Println()
}
func guessValue(g int8) int8 {
	g++
	return g
}

var solutionFound = false

func cloneSolution(s solution) solution {
	cloneSol := solution{}
	cloneSol.elems = make([]node, cap(s.elems))
	for i, v := range s.elems {
		cloneSol.elems[i] = v
	}
	return cloneSol
}

var goroutineCounter int
var printcounter int
var mutex sync.Mutex
var threads = 16
var parallelThread grid = 0

func chooseValue(s solution, index grid) {

	if goroutineCounter < threads && index > parallelThread {
		mutex.Lock()
		goroutineCounter++
		mutex.Unlock()
	}

	printcounter++
	if printcounter%1000000 == 0 {
		printcounter = 0
		printSolution(&s)
	}

	// printSolution(&s)
	// fmt.Println(":@:", index)
	if solutionFound {
		return
	}
	if s.elems[index].fixed {
		//a new function to be called
		if index == cubesize-1 {
			solutionFound = true
			printSolution(&s)
			waitgroup.Done()
		}
		cloneSol := cloneSolution(s)
		go chooseValue(cloneSol, index+1)
		return
	}

	for _, v := range cellvalues {
		s.elems[index].value = v
		if BoundingFunction(s.elems[index], s.elems) {
			if index == cubesize-1 {
				solutionFound = true
				printSolution(&s)
				waitgroup.Done()
			}
			cloneSol := cloneSolution(s)
			if (goroutineCounter < threads) && (index > parallelThread) {
				go chooseValue(cloneSol, index+1)
			} else {
				chooseValue(cloneSol, index+1)
			}
		}
		if solutionFound {
			break
		}
	}
	return

}

func createSolution() *solution {
	sol := new(solution)
	sol.elems = make([]node, cubesize)
	for i := range sol.elems {
		sol.elems[i].index = grid(i)
		sol.elems[i].block = indexToBlockMap[sol.elems[i].index]
		sol.elems[i].value = -1
	}

	return sol

}

var prefillValue map[grid]int8

func getPrefillSolution() {
	prefillValue = make(map[grid]int8)
	prefillValue[0] = 5
	prefillValue[5] = 4
	prefillValue[7] = 7
	prefillValue[8] = 2
	prefillValue[21] = 6
	prefillValue[22] = 1
	prefillValue[23] = 3
	prefillValue[30] = 3
	prefillValue[31] = 7
	prefillValue[32] = 8
	prefillValue[39] = 4
	prefillValue[43] = 6
	prefillValue[44] = 3
	prefillValue[47] = 4
	prefillValue[51] = 9
	prefillValue[55] = 9
	prefillValue[56] = 6
	prefillValue[60] = 1
	prefillValue[71] = 4
	prefillValue[73] = 1
	prefillValue[77] = 5
	prefillValue[78] = 3
	prefillValue[79] = 9

}

func populateSolution(s *solution) {
	for k, v := range prefillValue {
		s.elems[k].fixed = true
		s.elems[k].value = v
	}
}

var waitgroup sync.WaitGroup

func main() {

	waitgroup.Add(1)
	indexToBlockMap = make(map[grid]blockLocation)
	blockToIndexes = make(map[blockLocation][]grid)
	rowToIndexMap = make(map[grid][]grid)
	columnToIndexMap = make(map[grid][]grid)
	for i := 0; i < cubesize; i++ {
		determineBlock(grid(i))
	}
	// fmt.Println(indexToBlockMap)
	// fmt.Println(blockToIndexes)
	// fmt.Println(rowToIndexMap)
	// fmt.Println(columnToIndexMap)
	sol := createSolution()
	// getPrefillSolution()
	// populateSolution(sol)
	printSolution(sol)
	chooseValue(*sol, 0)
	// fmt.Println("Waiting for the go routines to complete")
	waitgroup.Wait()
}
