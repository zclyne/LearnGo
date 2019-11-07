package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int

	// 读入行数和列数，注意传参时必须取地址
	fmt.Fscanf(file, "%d %d", &row, &col)

	// 创建迷宫slice，注意必须用下面这种写法
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 迷宫中的某一点，作为一个结构体
type point struct {
	i, j int
}

// 4个可以走的方向
var dirs = [4]point {
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0}, // 下
	{0, 1}, // 右
}

// 点坐标和方向做加法，用于获得下一个点的位置
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 获取某个点在迷宫上为1还是0，第二个返回值bool表示是否存在越界
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 走迷宫
func walk(maze [][]int, start, end point) [][]int {
	// 创建二维切片steps，每个元素记录走的步数，在遍历完成后，要使用根据步数从大到小来重建路径
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// BFS用的队列
	Q := []point{start}

	for len(Q) > 0 {
		// 取出队首节点
		cur := Q[0]
		Q = Q[1:]

		// 当前节点就是终点
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir) // 下一个节点位置
			// 保证maze at next为0，且steps at next也为0，表示之前没有访问过这个点
			// 并且要保证next不是整个迷宫的起点
			val, ok := next.at(maze)
			if !ok || val == 1 { // next越界或撞墙，不进行下一步操作
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 { // next已经被访问过
				continue
			}
			if next == start { // next就是迷宫起点
				continue
			}
			// 此时next是一个合法的点，设置next的步数，并将其加入队列中
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	// 从文件中读入迷宫
	maze := readMaze("maze/maze.in")

	// 打印迷宫内容，检查是否有错误
	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}

	// 走迷宫
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	// 打印走完之后的steps结果
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

}