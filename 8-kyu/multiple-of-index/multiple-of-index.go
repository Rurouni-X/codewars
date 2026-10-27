package kata
​
func multipleOfIndex (ints []int) []int {
  res := make([]int, 0)
  
  for idx, num := range ints {
    
    if idx == 0 {
      continue
    }
    
    if num % idx == 0 {
      res = append(res, num)
    }
  }
  return res
}