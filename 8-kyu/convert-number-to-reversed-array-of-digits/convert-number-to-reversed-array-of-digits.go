package kata
​
import (
  "strconv"
)
​
func Digitize(n int) []int {
  arr := strconv.Itoa(n)
  res := make([]int, 0, len(arr))
  
  for i := len(arr) - 1; i >= 0; i-- {
    
    numb, _ := strconv.Atoi(string(arr[i]))
    res = append(res, numb)
  }
  
  return res
}
​