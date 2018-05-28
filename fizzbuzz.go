package main

import (
    "fmt"
    "testing"
)

func FizzBuzz(amount int) string {
    results := ""
    for i := 1; i <= amount; i++ {
        result := ""
        if i%3 == 0 { result += "Fizz" }
        if i%5 == 0 { result += "Buzz" }
        if result != "" {
            results += result + "\n"
            continue
        }
        results += fmt.Sprintf("%d\n", i)
    }
    return results
}

func TestFizzBuzz(t *testing.T) {
    got := FizzBuzz(8)
    want := 
`1
2
Fizz
4
Buzz
Fizz
7
8
`
    if got != want {
        t.Errorf("FizzBuzz(8) \n got: \n%v \n want: \n%v", got, want)
    } 
}

func main() {
    tests := []testing.InternalTest{{"TestFizzBuzz", TestFizzBuzz}}
    matchAll := func(t string, pat string) (bool, error) { return true, nil }
    testing.Main(matchAll, tests, nil, nil)

}
