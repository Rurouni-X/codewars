package kata
​
func CountSheeps(numbers []bool) (res int) {
  
  for _, v := range numbers {
    if v {
      res++
    }
  }
  return
}
​