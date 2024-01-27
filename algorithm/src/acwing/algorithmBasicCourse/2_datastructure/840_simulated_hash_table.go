package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2024/1/27
  @desc: 模拟哈希表
**/

const N, null = 100010, 0x3f3f3f3f

var (
	n              int
	h              [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

// find 开放寻址法
func find(x int) int {
	t := (x%N + N) % N
	// 开放寻址法,当前位置已经有元素在了，往后一个位置找空位
	for h[t] != null && h[t] != x {
		t++
		if t == N {
			t = 0
		}
	}
	return t
}

func main() {
	defer writer.Flush()
	for i := 0; i < N; i++ {
		h[i] = null
	}
	fmt.Fscan(reader, &n)

	for ; n > 0; n-- {
		var op string
		var x int
		fmt.Fscan(reader, &op, &x)
		if op[0] == 'I' {
			h[find(x)] = x
		} else {
			if h[find(x)] == null {
				fmt.Fprintln(writer, "No")
			} else {
				fmt.Fprintln(writer, "Yes")
			}
		}
	}
}

/* ====================================================拉链法==================================================== */
/*
package main

import (
    "fmt"
    "os"
    "bufio"
)

const N = 100003
var (
    n , idx int
    h , e , ne [N]int
    reader , writer = bufio.NewReader(os.Stdin) , bufio.NewWriter(os.Stdout)
)

// insert 拉链法
func insert(x int) {
    k := (x % N + N) % N
    // 头插，先存下当前节点和指向h[k]所指的，然后更新h[k]
    e[idx] = x
    ne[idx] = h[k]
    h[k] = idx
    idx++
}

func find(x int) bool {
    k := (x % N + N) % N
    for i := h[k] ; i != -1 ; i = ne[i] {
        if e[i] == x {
            return true
        }
    }
    return false
}

func main() {
    defer writer.Flush()
    fmt.Fscan(reader , &n)
    for i := 0 ; i < N ; i++ {
        h[i] = -1
    }

    for ; n > 0 ; n-- {
        var op string
        var x int
        fmt.Fscan(reader , &op , &x)
        if op[0] == 'I' {
            insert(x)
        }else {
            if find(x) {
                fmt.Fprintln(writer , "Yes")
            }else {
                fmt.Fprintln(writer , "No")
            }
        }
    }
}

*/
