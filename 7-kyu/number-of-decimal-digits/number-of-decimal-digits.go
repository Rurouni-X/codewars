package kata
​
import "fmt"
​
func Digits(n uint64) int {
  
  return len(fmt.Sprint(n))
}
​