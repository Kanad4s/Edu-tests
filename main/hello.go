package main

import (
	"fmt"
	"math"
)

type figures int

const (
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
)

func area(f figures) (func(float64) float64, bool) {
    switch f {
    case square:
      return func(x float64) float64 {return x * x}, true
    case circle:
      return func(x float64) float64 {return math.Pi * x * x}, true
    case triangle:
      return func(x float64) float64 {return math.Sqrt(3) * x * x / 4}, true
    default:
      return nil, false
    }
}


func main() {
    fig := square
    ar, ok := area(fig)
    if !ok {
        fmt.Println("area error")
        return
    }
    println(ar(5))
    input := "The quick brown 狐 jumped over the lazy 犬"
    // Get Unicode code points. 
    n := 0
    // создаём слайс рун 
    runes := make([]rune, len(input))
    // добавляем руны в слайс
    for _, r := range input {
        runes[n] = r
        n++
    }
    // убираем лишние нулевые руны
    runes = runes[0:n]
    // разворачиваем
    for i := 0; i < n/2; i++ {
		runes[i], runes[n - 1 - i] = runes[n - 1 - i], runes[i]
    }
    // преобразуем в строку UTF-8. 
    output := string(runes)
    fmt.Println(output)
}