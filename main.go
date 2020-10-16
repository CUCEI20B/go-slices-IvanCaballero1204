package main

import "fmt"

func main()  {
  var s []int;
  var tam, temp, total int;

  fmt.Scan(&tam);
  
  for i:= 0; i < tam; i++{
    fmt.Scan(&temp)
    s = append(s, temp)
  }
  
  for _, v := range s {
    total += v
  }
  
  fmt.Println(total)
}