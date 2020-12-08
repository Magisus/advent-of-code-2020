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

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)

    rows := strings.Split(strings.Trim(string(data), "\n"), "\n")
    width := len(rows[0])
    height := len(rows)

    pos := [2]int{0, 0}
    treeCount := 0
    for i := 0; i < height; {
        if IsTree(pos, rows) {
            treeCount++
        }
        i++
        pos[0] = i
        pos[1] = (pos[1] + 3) % width
    }
    fmt.Println("Tree count: " + strconv.Itoa(treeCount))
}
