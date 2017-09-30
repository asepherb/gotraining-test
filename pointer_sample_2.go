package main


const size = 1002

func main() {

  s := "hello"
  stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {

  println(c, s, *s)

  c++
  if c == 10 {
    return
  }

  stackCopy(s, c, a)
}
