package week4_assignment

func solve(board [][]byte) {
	height := len(board)
	width := len(board[0])

	//建立访问矩阵
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	//将边缘的O置为中间值A，便于后续bfs从边缘向内部遍历
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (i == 0 || i == height-1 || j == 0 || j == width-1) && board[i][j] == 'O' {
				board[i][j] = 'A'
			}
		}
	}

	//从边缘向内部遍历，把与边界相连的连通块都置为中间值A
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 'A' && visited[i][j] == false {
				board = bfs(board, visited, height, width, i, j)
			}
		}
	}

	//最后再遍历一遍，把剩下的O置为X，把A置回O
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func bfs(board [][]byte, visited [][]bool, height, width, h, w int) [][]byte {
	board[h][w] = 'A'
	visited[h][w] = true

	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, 1, 0, -1}

	point := map[string]int{
		"height": h,
		"width":  w,
	}
	queue := make([]map[string]int, 0)
	queue = append(queue, point)

	for len(queue) > 0 {
		x := queue[len(queue)-1]["width"]
		y := queue[len(queue)-1]["height"]
		queue = queue[:len(queue)-1]

		for index := 0; index < 4; index++ {
			newX := x + dirX[index]
			newY := y + dirY[index]

			if !checkChangeability(board, visited, height, width, newY, newX) {
				continue
			}

			point := map[string]int{
				"height": newY,
				"width":  newX,
			}
			queue = append(queue, point)
			board[newY][newX] = 'A'
			visited[newY][newX] = true
		}
	}

	return board
}

func checkChangeability(board [][]byte, visited [][]bool, height, width, h, w int) bool {
	if w < 0 || w > width-1 || h < 0 || h > height-1 {
		return false
	}

	if visited[h][w] == true {
		return false
	}

	if board[h][w] != 'O' {
		return false
	}

	return true
}
