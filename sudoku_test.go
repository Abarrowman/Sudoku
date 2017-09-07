package main

import "testing"

func TestBoardFromString(t *testing.T) {
	arr := [dimQu]int{
		8, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 3, 6, 0, 0, 0, 0, 0,
		0, 7, 0, 0, 9, 0, 2, 0, 0,
		0, 5, 0, 0, 0, 7, 0, 0, 0,
		0, 0, 0, 0, 4, 5, 7, 0, 0,
		0, 0, 0, 1, 0, 0, 0, 3, 0,
		0, 0, 1, 0, 0, 0, 0, 6, 8,
		0, 0, 8, 5, 0, 0, 0, 1, 0,
		0, 9, 0, 0, 0, 0, 4, 0, 0}

	arrFromString := boardFromString(`
		8,0,0,0,0,0,0,0,0,
		0,0,3,6,0,0,0,0,0,
		0,7,0,0,9,0,2,0,0,
		0,5,0,0,0,7,0,0,0,
		0,0,0,0,4,5,7,0,0,
		0,0,0,1,0,0,0,3,0,
		0,0,1,0,0,0,0,6,8,
		0,0,8,5,0,0,0,1,0,
		0,9,0,0,0,0,4,0,0`).squares
	if arr != arrFromString {
		t.Error("The board parsed from the string is not what was expected.")
	}
}

func TestSolve(t *testing.T) {
	//1
	//"World's Hardest"
	//http://www.telegraph.co.uk/news/science/science-news/9359579/Worlds-hardest-sudoku-can-you-crack-it.html
	//http://www.telegraph.co.uk/news/science/science-news/9360022/Worlds-hardest-sudoku-the-answer.html
	input := boardFromString(`
		-------------
		|800|000|000|
		|003|600|000|
		|070|090|200|
		-------------
		|050|007|000|
		|000|045|700|
		|000|100|030|
		-------------
		|001|000|068|
		|008|500|010|
		|090|000|400|
		-------------`)
	solution := boardFromString(`
		-------------
		|812|753|649|
		|943|682|175|
		|675|491|283|
		-------------
		|154|237|896|
		|369|845|721|
		|287|169|534|
		-------------
		|521|974|368|
		|438|526|917|
		|796|318|452|
		-------------`).squares

	solved, _ := input.solve()
	if solved.squares != solution {
		t.Error("Wrong solution for 1.")
	}

	//2
	input = boardFromString(`
		-------------
		|120|400|300|
		|300|010|050|
		|006|000|100|
		-------------
		|700|090|000|
		|040|603|000|
		|003|002|000|
		-------------
		|500|080|700|
		|007|000|005|
		|000|000|098|
		-------------`)
	solution = boardFromString(`
		-------------
		|128|465|379|
		|374|219|856|
		|956|837|142|
		-------------
		|765|198|423|
		|249|673|581|
		|813|542|967|
		-------------
		|592|386|714|
		|487|921|635|
		|631|754|298|
		-------------`).squares

	solved, _ = input.solve()
	if solved.squares != solution {
		t.Error("Wrong solution for 2.")
	}

	//3
	input = boardFromString(`
		-------------
		|056|907|400|
		|081|040|000|
		|000|015|090|
		-------------
		|000|003|857|
		|840|060|023|
		|739|200|000|
		-------------
		|060|580|000|
		|000|070|360|
		|008|306|570|
		-------------`)
	solution = boardFromString(`
		-------------
		|256|937|481|
		|981|642|735|
		|473|815|692|
		-------------
		|612|493|857|
		|845|761|923|
		|739|258|146|
		-------------
		|367|584|219|
		|524|179|368|
		|198|326|574|
		-------------`).squares

	solved, _ = input.solve()
	if solved.squares != solution {
		t.Error("Wrong solution for 3.")
	}
}

func BenchmarkMain(b *testing.B) {
	main()
}
