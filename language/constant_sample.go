package main

func main() {

  const ui = 12345
  const uf = 3.14159

  const ti int = 123456
  const tf float64 = 3.15159

  //constant 1000 overflows uint8
  //const myUint8 uint8 = 1000

  var answer = 3 * 0.333

  println(answer)

  const third = 1 / 3

  const one int8 = 1
  const two = 2 * one

  //this iota type
  const (
    CategoryBooks = iota
    CategoryHealth
    CategoryClothing
  )

  println("Category : ",CategoryBooks, CategoryHealth, CategoryClothing)

  type StereoType int

  const (
    TypicalNoob StereoType = iota
    TypicalHypster
  )

  println("StreoType : ", TypicalNoob, TypicalHypster)
}
