package main

import(
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func IsTree(pos [2]int, rows []string) bool {
    sym := rows[pos[0]][pos[1]]
    return string(sym) == "#"
}

func CountTrees(slope [2]int, rows []string) int {
    width := len(rows[0])
    height := len(rows)

    pos := [2]int{0, 0}
    treeCount := 0
    for i := 0; i < height; {
        if IsTree(pos, rows) {
            treeCount++
        }
        i += slope[0]
        pos[0] = i
        pos[1] = (pos[1] + slope[1]) % width
    }
    fmt.Println("Tree count: " + strconv.Itoa(treeCount))
    return treeCount
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)

    rows := strings.Split(strings.Trim(string(data), "\n"), "\n")
    treeCount1 := CountTrees([2]int{1, 1}, rows)
    treeCount2 := CountTrees([2]int{1, 3}, rows)
    treeCount3 := CountTrees([2]int{1, 5}, rows)
    treeCount4 := CountTrees([2]int{1, 7}, rows)
    treeCount5 := CountTrees([2]int{2, 1}, rows)
    total := treeCount1 * treeCount2 * treeCount3 * treeCount4 * treeCount5
    fmt.Println(total)
}
