package acmguide

import (
	"math"
)

// Hamilton
//
// 例题: ch0103/acwing91 https://ac.nowcoder.com/acm/contest/996/D
//
// 给定一张 n(n≤20)(n≤20) 个点的带权无向图
// 点从0∼n−10∼n−1标号，求起点 0 到终点 n-1 的最短Hamilton路径。
// Hamilton路径的定义是从 0 到 n-1 不重不漏地经过每个点恰好一次。
//
// 思路:
//
// 直观的思路是使用回溯算法穷举所有路径, 然后计算每个路径的长度.
func Hamilton(n int, weight [20][20]int) int {
	// TODO:
	return 0
}

// HamiltonWithBackTrack 使用回溯算法计算最短 Hamilton 路径
func HamiltonWithBackTrack(n int, weight [20][20]int) int {
	backtracker := NewHamiltonBacktrack(n, weight)
	return backtracker.Calc()
}

// NewHamiltonBacktrack 实例化一个 最短hamilton路径回溯计算器
func NewHamiltonBacktrack(n int, weight [20][20]int) *HamiltonBacktrack {
	return &HamiltonBacktrack{
		N:      n,
		Weight: weight,

		visited:       make(map[int]struct{}, n),
		latest:        0,
		visitedWeight: 0,

		result: math.MaxInt,
	}
}

// HamiltonBacktrack 回溯算法计算 Hamilton 最短路径问题
type HamiltonBacktrack struct {
	// 订单数量
	N int
	// 订单间权重
	Weight [20][20]int

	// 访问过的顶点
	visited map[int]struct{}
	// 最近访问的顶点
	latest int
	// 已访问顶点的总权重
	visitedWeight int

	// 计算结果
	result int
}

// Calc 计算 Hamilton 最短路径
func (h *HamiltonBacktrack) Calc() int {
	h.backtrack()
	return h.result
}

// backtrack 回溯
func (h *HamiltonBacktrack) backtrack() {
	// 计算结果
	if h.isSolution() {
		h.recordSolution()
		return
	}

	for choice := 0; choice < h.N; choice++ {
		// 剪枝
		if h.isValid(choice) {
			prev := h.latest
			// 尝试
			h.makeChoice(choice)
			h.backtrack()
			// 回溯
			h.undoChoice(prev)
		}
	}
}

// isSolution 判断是否为解
func (h *HamiltonBacktrack) isSolution() bool {
	return len(h.visited) == h.N && h.latest == h.N-1
}

// recordSolution 记录解
func (h *HamiltonBacktrack) recordSolution() {
	if h.visitedWeight < h.result {
		h.result = h.visitedWeight
	}
}

// isValid() 判断是否需要剪枝
func (h *HamiltonBacktrack) isValid(choice int) bool {
	if len(h.visited) == 0 && choice != 0 {
		return false
	}
	// 最近遍历的 vertex 时 h.N - 1 且 已经遍历了所有 vertices
	if choice == h.N-1 && len(h.visited) != h.N-1 {
		return false
	}
	// 即将遍历的节点没有被访问过
	_, ok := h.visited[choice]
	return !ok
}

// makeChoice 尝试
func (h *HamiltonBacktrack) makeChoice(choice int) {
	// 记录已访问顶点
	h.visited[choice] = struct{}{}
	pre := h.latest
	h.latest = choice
	// 更新总权重
	h.visitedWeight += h.Weight[pre][h.latest]
}

// undoChoice 回退
func (h *HamiltonBacktrack) undoChoice(prev int) {
	latest := h.latest
	delete(h.visited, h.latest)
	h.latest = prev

	h.visitedWeight -= h.Weight[prev][latest]
}
