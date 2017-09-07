package main

import "bufio"
import "fmt"
import "io"
import "os"

const dim int = 3
const dimSq int = dim * dim
const dimSqP int = dimSq + 1
const dimCu int = dim * dim * dim
const dimQu int = dimSq * dimSq

func main() {
	b := boardFromString(bufferStdin())
	fmt.Println("Input:")
	b.print()
	fmt.Println("Output:")

	solved, itterations := b.solve()
	if solved == nil {
		fmt.Println("No solution exists.")
	} else {
		solved.print()
		fmt.Printf("Solved in %v itterations.", itterations)
	}
}

func bufferStdin() string {
	var content string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		content += line
	}
	return content
}

func boardFromString(s string) *Board {
	var arr [dimQu]int
	zCount := 0
	idx := 0
	for _, r := range s {
		if idx == dimQu {
			break
		}
		switch r {
		case '0':
			zCount++
			idx++
		case '1':
			arr[idx] = 1
			idx++
		case '2':
			arr[idx] = 2
			idx++
		case '3':
			arr[idx] = 3
			idx++
		case '4':
			arr[idx] = 4
			idx++
		case '5':
			arr[idx] = 5
			idx++
		case '6':
			arr[idx] = 6
			idx++
		case '7':
			arr[idx] = 7
			idx++
		case '8':
			arr[idx] = 8
			idx++
		case '9':
			arr[idx] = 9
			idx++
		}
	}
	if idx < dimQu {
		panic(fmt.Sprintf("Only %v numbers were inputted %v required",
			idx, dimQu))
	}
	b := Board{dimQu - zCount, arr}
	return &b
}

func boardFromArray(arr [dimQu]int) *Board {
	count := 0
	for i := 0; i < dimQu; i++ {
		if arr[i] != 0 {
			count++
		}
	}
	b := Board{count, arr}
	return &b
}

func (b *Board) print() {
	for q := 0; q < dim+1+dimSq; q++ {
		fmt.Print("-")
	}
	fmt.Println("")
	for h := 0; h < dim; h++ {
		for i := 0; i < dim; i++ {
			fmt.Print("|")
			for j := 0; j < dim; j++ {
				for k := 0; k < dim; k++ {
					fmt.Print(b.squares[h*dimCu+i*dimSq+j*dim+k])
				}
				fmt.Print("|")
			}
			fmt.Println("")
		}
		for q := 0; q < dim+1+dimSq; q++ {
			fmt.Print("-")
		}
		fmt.Println("")
	}
}

func (b *Board) isSolved() bool {
	return b.placed == dimQu
}

func (b *Board) clone() *Board {
	clone := *b
	return &clone
}

func getSpotInfo(spot int) (int, int, int) {
	row := spot / dimSq
	col := spot % dimSq
	sq := (row/dim)*dim + (col / dim)
	return row, col, sq
}

func fullAvailMap() Avail {
	a := Avail{}
	a.count = dimSq
	for i := 1; i <= dimSq; i++ {
		a.vals[i] = true
	}
	return a
}

func emptyAvailMap() Avail {
	a := Avail{}
	for i := 1; i <= dimSq; i++ {
		a.vals[i] = false
	}
	return a
}

func (b *Board) colAvailable(col int, a *Avail) {
	for s := col; s < dimQu; s += dimSq {
		n := b.squares[s]
		if n != 0 {
			a.vals[n] = false
		}
	}
}

func (b *Board) rowAvailable(row int, a *Avail) {
	for s := row * dimSq; s < (row+1)*dimSq; s++ {
		n := b.squares[s]
		if n != 0 {
			a.vals[n] = false
		}
	}
}

func (b *Board) sqAvailable(sq int, a *Avail) {
	tl := (sq/dim)*dimCu + (sq%dim)*dim
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			s := tl + i*dimSq + j
			n := b.squares[s]
			if n != 0 {
				a.vals[n] = false
			}
		}
	}
}

func (b *Board) legalMoves(s int) Avail {
	row, col, sq := getSpotInfo(s)
	avail := fullAvailMap()
	a := &avail
	b.rowAvailable(row, a)
	b.colAvailable(col, a)
	b.sqAvailable(sq, a)
	a.calcCount()
	return avail
}

func (b *Board) bestMoves() (int, *Avail, bool) {
	min := dimSq + 1
	ok := false
	bestS := -1
	var bestA *Avail
	for s := 0; s < dimQu; s++ {
		if b.squares[s] == 0 {
			a := b.legalMoves(s)
			if a.count == 1 {
				return s, &a, true
			} else if a.count > 1 && a.count < min {
				min = a.count
				bestS = s
				bestA = &a
				ok = true
			}
		}
	}
	return bestS, bestA, ok
}

func (b *Board) solve() (*Board, int) {
	i := 0
	s := b.clone().innerSolve(&i)
	return s, i
}

func (b *Board) set(s int, v int) {
	//row, col, sq := getSpotInfo(s)
	//fmt.Printf("s=%v row=%v col=%v sq=%v v=%v\n", s, row, col, sq, v)
	b.squares[s] = v
	b.placed++
}

func (b *Board) innerSolve(itterations *int) *Board {
	for !b.isSolved() {
		*itterations++
		s, a, ok := b.bestMoves()

		//b.print()
		//fmt.Printf("placed %v\n\n", b.placed)

		if !ok {
			return nil
		} else if a.count == 1 {
			//fmt.Println("Forced")
			b.set(s, a.getFirst())
		} else {
			for i := 1; i <= dimSq; i++ {
				if a.vals[i] {
					cl := b.clone()
					cl.set(s, i)
					cl = cl.innerSolve(itterations)
					if cl != nil {
						return cl
					}
				}
			}
			return nil
		}
	}
	return b
}

//Board represents a partially or fully filled sudoku board
type Board struct {
	placed  int
	squares [dimQu]int
}

func (a *Avail) getFirst() int {
	for i := 1; i <= dimSq; i++ {
		if a.vals[i] {
			return i
		}
	}
	return -1
}

func (a *Avail) calcCount() {
	a.count = 0
	for i := 1; i <= dimSq; i++ {
		if a.vals[i] {
			a.count++
		}
	}
}

//Avail represents available moves in a sudoku square
type Avail struct {
	count int
	vals  [dimSqP]bool
}
