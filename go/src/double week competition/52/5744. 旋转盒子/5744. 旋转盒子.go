package main

import "fmt"

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	b := make([][]byte, n)
	for i := 0; i < n; i++ {
		b[i] = make([]byte, m)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; {
			count := 0
			k := j
			for k < n && box[i][k] != '*' {
				if box[i][k] == '#' {
					count++
				}
				k++
			}

			for t := j; t < k; t++ {
				if t < k-count {
					box[i][t] = '.'
				} else {
					box[i][t] = '#'
				}
			}

			j = k+1
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b[j][m-1-i] = box[i][j]
		}
	}

	//fmt.Println(b)

	return b
}

func main() {
	fmt.Println('.', ".", '#', "#", '*', "*")
	rotateTheBox([][]byte{
		{'#', '.', '#'},
		})

	rotateTheBox([][]byte{
		{'#','.','*','.'},
		{'#','#','*','.'},
	})

	rotateTheBox([][]byte{
		{'#','#','*','.','*','.'},
		{'#','#','#','*','.','.'},
		{'#','#','#','.','#','.'},
	})
}