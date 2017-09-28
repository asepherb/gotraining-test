package main


func main() {

  count := 10
  println("before :", count, &count)

  increment(count)

  println("after :", count, &count)
}

func increment(inc int) {

  inc++
  println("inc :", inc, &inc)
}
