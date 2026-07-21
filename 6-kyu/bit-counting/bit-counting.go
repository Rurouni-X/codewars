package kata
​
import "fmt"
​
func CountBits(n uint) (res int) {
    
  value := fmt.Sprintf("%b", n)
  
  for _, v := range value {
    
    if v == '1' {
      res++
    }
  }
  return
}