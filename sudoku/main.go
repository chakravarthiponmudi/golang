package main

import "fmt"

type grid int16
type blockLocation int8

const gridsize = 9
const maxCellValue = 9
const blocksize = 3
const cubesize = 81

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
	// printSolution(sol)
	solveSudoku(sol, 0, 1)
}
