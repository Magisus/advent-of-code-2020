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

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)

    nums := strings.Split(strings.Trim(string(data), "\n"), "\n")
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            for k := j+1; k < len(nums); k++ {
                num1, _ := strconv.Atoi(nums[i])
                num2, _ := strconv.Atoi(nums[j])
                num3, _ := strconv.Atoi(nums[k])
                sum := num1 + num2 + num3
                if sum == 2020 {
                    fmt.Println(sum)
                    fmt.Println(strings.Join([]string{nums[i], nums[j], nums[k]}, " "))
                    fmt.Println(num1 * num2 * num3)
                    return;
                }
            }
        }
    }
}
