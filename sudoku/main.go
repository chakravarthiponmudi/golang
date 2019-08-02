package main

import (
	"fmt"
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
		fmt.Printf("\t%d", n.value)
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
func solveSudoku(s *solution, index grid, guess int8) {

	// printSolution(s)
	if s.elems[index].fixed {
		solveSudoku(s, index+1, 1)
		return
	}
	s.elems[index].value = guess
	if BoundingFunction(s.elems[index], s.elems) {
		fmt.Printf("\nindex %d, guess %d\n", index, guess)
		if index == cubesize-1 {
			solutionFound = true
			printSolution(s)
		} else {
			for i := 1; i <= maxCellValue && !solutionFound; i++ {
				solveSudoku(s, index+1, int8(i))
			}
		}
	} else {
		if guess == maxCellValue {
			s.elems[index].value = -1
		} else {
			solveSudoku(s, index, guessValue(guess))
		}
	}
}

func chooseValue(s solution, index grid) {

	printSolution(&s)
	fmt.Println(":@:", index)
	if solutionFound {
		return
	}
	if s.elems[index].fixed {
		//a new function to be called
		cloneSol := cloneSolution(s)
		chooseValue(cloneSol, index+1)
		return
	}

	for _, v := range cellvalues {
		s.elems[index].value = v
		if BoundingFunction(s.elems[index], s.elems) {
			if index == cubesize-1 {
				solutionFound = true
				printSolution(&s)
			}
			cloneSol := cloneSolution(s)
			chooseValue(cloneSol, index+1)
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
	prefillValue[6] = 3
	prefillValue[8] = 7
	prefillValue[9] = 9
	prefillValue[15] = 5
	prefillValue[16] = 1
	prefillValue[17] = 4
	prefillValue[18] = 3
	prefillValue[20] = 4
	prefillValue[22] = 1
	prefillValue[23] = 6
	prefillValue[25] = 2
	prefillValue[29] = 6
	prefillValue[34] = 5
	prefillValue[36] = 2
	prefillValue[41] = 4
	prefillValue[50] = 9
	prefillValue[51] = 4
	prefillValue[56] = 1
	prefillValue[57] = 9
	prefillValue[61] = 7
	prefillValue[62] = 6
	prefillValue[70] = 3
	prefillValue[74] = 7
	prefillValue[75] = 6
	prefillValue[80] = 5

}

func populateSolution(s *solution) {
	for k, v := range prefillValue {
		s.elems[k].fixed = true
		s.elems[k].value = v
	}
}

func main() {
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
	getPrefillSolution()
	populateSolution(sol)
	printSolution(sol)
	// solveSudoku(sol, 0, 1)
	chooseValue(*sol, 0)
}
