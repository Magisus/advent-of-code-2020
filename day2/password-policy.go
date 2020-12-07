package main

import(
    "fmt"
    "regexp"
    "io/ioutil"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)

    entries := strings.Split(strings.Trim(string(data), "\n"), "\n")

    positionsRegex := regexp.MustCompile("^([0-9]+)-([0-9]+)")
    requiredLetterRegex := regexp.MustCompile("([a-z]):")
    passwordRegex := regexp.MustCompile("[a-z]+$")

    validCount := 0

    for _, entry := range entries {
        positions := positionsRegex.FindStringSubmatch(entry)
        fmt.Println(positions[1] + " or " + positions[2])
        // Policy is 1-indexed
        pos1, _ := strconv.Atoi(positions[1])
        pos2, _ := strconv.Atoi(positions[2])
        pos1 = pos1 - 1
        pos2 = pos2 - 1

        requiredLetter := requiredLetterRegex.FindStringSubmatch(entry)[1]
        fmt.Println(requiredLetter)

        password := passwordRegex.FindString(entry)
        fmt.Println(password)

        oneValidLetter := false
        if pos1 < len(password) && string(password[pos1]) == requiredLetter {
            oneValidLetter = !oneValidLetter
        }
        if pos2 < len(password) && string(password[pos2]) == requiredLetter {
            oneValidLetter = !oneValidLetter
        }

        if oneValidLetter {
            fmt.Println(true)
            validCount++
        } else {
            fmt.Println(false)
        }
    }
    fmt.Println(validCount)
}
