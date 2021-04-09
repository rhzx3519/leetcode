package main

func rotate(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > i {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
		Reverse(&matrix[i])
	}
}

func Reverse(a *[]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}