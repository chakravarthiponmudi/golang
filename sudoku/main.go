package main

import "fmt"

type grid int16
type blockLocation int8

const gridsize = 9
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
		column = 9
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
}

func main() {
	indexToBlockMap = make(map[grid]blockLocation)
	blockToIndexes = make(map[blockLocation][]grid)
	for i := 0; i < cubesize; i++ {
		determineBlock(grid(i))
	}
	fmt.Println(indexToBlockMap)
	fmt.Println(blockToIndexes)
}
