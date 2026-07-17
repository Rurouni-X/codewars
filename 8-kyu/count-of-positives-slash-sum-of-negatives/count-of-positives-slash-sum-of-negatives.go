package kata
​
func CountPositivesSumNegatives(numbers []int) []int {
  
  res := make([]int, 2)
  
  for _, numb := range numbers {
    
    if numb > 0 {
      res[0]++
    } else {
      res[1] += numb
    }
    
  }
  return res
}
​