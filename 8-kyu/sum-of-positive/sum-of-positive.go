package kata
​
func PositiveSum(numbers []int) (res int) {
​
  for _, numb := range numbers {
    
    if numb > 0 {
      res += numb
    }
    
  }
  return
}
​