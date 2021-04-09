package main

func setZeroes(matrix [][]int)  {
	r0, c0 := false, false
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					r0 = true
				}
				if j == 0 {
					c0 = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if r0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if c0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
