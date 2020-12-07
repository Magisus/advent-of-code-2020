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

    minMaxRegex := regexp.MustCompile("^([0-9]+)-([0-9]+)")
    requiredLetterRegex := regexp.MustCompile("([a-z]):")
    passwordRegex := regexp.MustCompile("[a-z]+$")

    validCount := 0

    for _, entry := range entries {
        minMax := minMaxRegex.FindStringSubmatch(entry)
        min, _ := strconv.Atoi(minMax[1])

        max, _ := strconv.Atoi(minMax[2])

        requiredLetter := requiredLetterRegex.FindStringSubmatch(entry)[1]

        password := passwordRegex.FindString(entry)

        count := 0
        for _, c := range password {
            if string(c) == requiredLetter {
               count++
            }
        }
        if count >= min && count <= max {
            validCount++
        }
    }
    fmt.Println(validCount)
}
