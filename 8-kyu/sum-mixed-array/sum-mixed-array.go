package kata
​
import "strconv"
​
func SumMix(arr []any) (res int) {
  
  for _, v := range arr {
    
    switch num := v.(type) {
      
      case string:
        a, _ := strconv.Atoi(num)
        res += a
      case int:
        res += num
    }
  }
  return
}
​