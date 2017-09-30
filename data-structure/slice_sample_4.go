package main

import "fmt"

func main() {

  var data []string

  lastcap := cap(data)

  for record := 1; record <= 102400; record++ {

    data = append(data, fmt.Sprint("rec : %d", record))

    if lastcap != cap(data) {

      //calculate change
      capChg := float64(cap(data) - lastcap) / float64(lastcap) * 100

      lastcap = cap(data)

      fmt.Printf("addr [%p]\tIndex[%d]\t\tcap[%d - %2.f%%]\n", &data[0], record, cap(data), capChg)
    }
  }
}
