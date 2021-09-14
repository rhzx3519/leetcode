package main

import "fmt"

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	vpre := [][]int{} // vertical pre sum
	for c := 0; c < n; c++ {
		tmp := make([]int, m+1)
		for r := 1; r <= m; r++ {
			tmp[r] = grid[r-1][c] + tmp[r-1]
		}
		vpre = append(vpre, tmp)
	}

	hpre := [][]int{} // horizon pre sum
	for r := 0; r < m; r++ {
		tmp := make([]int, n+1)
		for c := 1; c <= n; c++ {
			tmp[c] = grid[r][c-1] + tmp[c-1]
		}
		hpre = append(hpre, tmp)
	}
	//var radius1 int
	//for r := 0; r < m; r++ {
	//	for c := 0; c < n; c++ {
	//		if r == 2 && c == 2 {
	//			fmt.Println(r, c)
	//		}
	//		for k := 1; r+k < m && r-k >= 0 && c+k < n && c-k >= 0; k++ {
	//			//lb := []int{r+k, c-k} // left, bottom
	//			//rb := []int{r+k, c+k} // right, bottom
	//			//lt := []int{r-k, c-k} // left, top
	//			//rt := []int{r-k, c+k} // right, top
	//			//bottom := hpre[r+k][c+k+1] - hpre[r+k][c-k]
	//			//top := hpre[r-k][c+k+1] - hpre[r-k][c-k]
	//			//left := vpre[c-k][r+k+1] - vpre[c-k][r-k]
	//			//right := vpre[c+k][r+k+1] - vpre[c+k][r-k]
	//			//
	//			r1 := hpre[r-k][c+k+1] - hpre[r-k][c-k]
	//			var i int
	//			for i = r-k+1; i <= r+k; i++ {
	//				tmp := hpre[i][c+k+1] - hpre[i][c-k]
	//				if tmp != r1 {
	//					break
	//				}
	//			}
	//			if i != r+k+1 {
	//				continue
	//			}
	//
	//			c1 := vpre[c-k][r+k+1] - vpre[c-k][r-k]
	//			if c1 != r1 {
	//				continue
	//			}
	//			i = 0
	//			for i = c-k+1; i <= c+k; i++ {
	//				tmp := vpre[i][r+k+1] - vpre[i][r-k]
	//				if tmp != c1 {
	//					break
	//				}
	//			}
	//			if i != c+k+1 {
	//				continue
	//			}
	//
	//			var t1, t2 int
	//			for i, j := r-k, c-k; i <= r+k; i, j = i+1, j+1 {
	//				t1 += grid[i][j]
	//			}
	//			for i, j := r-k, c+k; i <= r+k; i, j = i+1, j-1 {
	//				t2 += grid[i][j]
	//			}
	//			if r1 != t1 || r1 != t2 {
	//				continue
	//			}
	//
	//			if k > radius1 {
	//				radius1 = k
	//			}
	//		}
	//	}
	//}
	//
	//fmt.Println(1 + radius1<<1)

	var radius2 int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			for k := 1; r+k < m && c+k < n; k++ {
				//bottom := hpre[r+k][c+k+1] - hpre[r+k][c]
				//top := hpre[r][c+k+1] - hpre[r][c]
				//left := vpre[c][r+k+1] - vpre[c][r]
				//right := vpre[c+k][r+k+1] - vpre[c+k][r]

				r1 := hpre[r][c+k+1] - hpre[r][c]
				var i int
				for i = r+1; i <= r+k; i++ {
					tmp := hpre[i][c+k+1] - hpre[i][c]
					if tmp != r1 {
						break
					}
				}
				if i != r+k+1 {
					continue
				}

				c1 := vpre[c][r+k+1] - vpre[c][r]
				for i = c+1; i <= c+k; i++ {
					tmp := vpre[i][r+k+1] - vpre[i][r]
					if tmp != c1 {
						break
					}
				}
				if i != c+k+1 {
					continue
				}

				var t1, t2 int
				for i, j := r, c; i <= r+k; i, j = i+1, j+1 {
					t1 += grid[i][j]
				}
				for i, j := r, c+k; i <= r+k; i, j = i+1, j-1 {
					t2 += grid[i][j]
				}
				if r1 != t1 || r1 != t2 {
					continue
				}

				if k > radius2 {
					radius2 = k
				}
			}
		}
	}

	fmt.Println(1 + radius2)
	return 1 + radius2
}


func main() {
	largestMagicSquare([][]int{{7,1,4,5,6}, {2,5,1,6,4}, {1,5,4,3,2}, {1,2,7,3,4}})
	largestMagicSquare([][]int{{5,1,3,1}, {9,3,3,1}, {1,3,3,8}})
}