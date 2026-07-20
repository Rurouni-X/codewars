package kata
​
import "strconv"
​
func StringToNumber(str string) int {
  res, err := strconv.Atoi(str)
  
  if err != nil {
    return 0
  }
  return res
}
​