//package main
//
//import (
//	"fmt"
//)
//
//func reverse(s []int) []int {
//	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//		s[i], s[j] = s[j], s[i]
//	}
//	return s
//}
//
//func processRow(row []int) []int {
//	compressed := make([]int, 0)
//	for _, num := range row {
//		if num != 0 {
//			compressed = append(compressed, num)
//		}
//	}
//
//	merged := make([]int, 0)
//	i := 0
//	for i < len(compressed) {
//		if i+1 < len(compressed) && compressed[i] == compressed[i+1] {
//			merged = append(merged, compressed[i]*2)
//			i += 2
//		} else {
//			merged = append(merged, compressed[i])
//			i++
//		}
//	}
//
//	newRow := make([]int, len(row))
//	copy(newRow, merged)
//	return newRow
//}
//
//func transpose(matrix [][]int) [][]int {
//	if len(matrix) == 0 {
//		return nil
//	}
//	rows := len(matrix)
//	cols := len(matrix[0])
//	transposed := make([][]int, cols)
//	for i := range transposed {
//		transposed[i] = make([]int, rows)
//	}
//	for i := 0; i < rows; i++ {
//		for j := 0; j < cols; j++ {
//			transposed[j][i] = matrix[i][j]
//		}
//	}
//	return transposed
//}
//
//func main() {
//	var m, n int
//	fmt.Scan(&m, &n)
//
//	matrix := make([][]int, m)
//	for i := 0; i < m; i++ {
//		matrix[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			fmt.Scan(&matrix[i][j])
//		}
//	}
//
//	var k int
//	fmt.Scan(&k)
//	operations := make([]int, k)
//	for i := 0; i < k; i++ {
//		fmt.Scan(&operations[i])
//	}
//
//	for _, op := range operations {
//		switch op {
//		case 0: // 上移
//			transposed := transpose(matrix)
//			for i := 0; i < len(transposed); i++ {
//				transposed[i] = processRow(transposed[i])
//			}
//			matrix = transpose(transposed)
//		case 1: // 下移
//			transposed := transpose(matrix)
//			for i := 0; i < len(transposed); i++ {
//				reversed := reverse(transposed[i])
//				processed := processRow(reversed)
//				transposed[i] = reverse(processed)
//			}
//			matrix = transpose(transposed)
//		case 2: // 左移
//			for i := 0; i < len(matrix); i++ {
//				matrix[i] = processRow(matrix[i])
//			}
//		case 3: // 右移
//			for i := 0; i < len(matrix); i++ {
//				reversed := reverse(matrix[i])
//				processed := processRow(reversed)
//				matrix[i] = reverse(processed)
//			}
//		}
//	}
//
//	// 输出结果
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if j > 0 {
//				fmt.Print(" ")
//			}
//			fmt.Print(matrix[i][j])
//		}
//		fmt.Println()
//	}
//}

package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

func solve(root *Node, k int) int {
	if root == nil {
		fmt.Println("-1")
		return 0
	}
	var fast *Node = root
	var slow *Node = root

	for i := k; i >= 0; i-- {
		if fast.next == nil {
			fmt.Println("-1")
			return 0
		}
		fast = fast.next
	}
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	fmt.Println(slow.val)
	return slow.val
}

func main() {
	root := &Node{
		val:  1,
		next: nil,
	}
	solve(root, -1)
}