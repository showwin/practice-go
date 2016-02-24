package main

import (
  "fmt"
  "time"
  "log"
)

func listenChannel(queue chan int, b bool) {
  for out := range queue {
    time.Sleep(10000)
    result = append(result, out)
    fmt.Println(b)
  }
  fmt.Println("Finish!!!!")
}

var result = make([]int, 0, 100)

var fin = make(chan int)

func main() {

  finishTime := time.Now().Add(1 * time.Second)

  queue := make(chan int)

  go listenChannel(queue, true)
  go listenChannel(queue, false)

  //conc := 5

 // 時間になるまで、リクエストをqueue にいれて、
 // listenChannnel の方でリクエストを送る
 i := 0
 log.Print("Start")
  for {
    i++
    if time.Now().Before(finishTime) {
      queue <- i
    } else {
      fmt.Println(i)
      break
    }
  }

  fmt.Println(result[len(result)-1])
  log.Print("Finish")
}

func boom(ch chan int, i int) {
  time.Sleep(1000)
  // レポート入れる
  ch <- i
}
