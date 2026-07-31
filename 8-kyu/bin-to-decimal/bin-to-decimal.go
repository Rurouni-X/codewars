package kata
​
import "strconv"
​
func BinToDec(bin string) int {
      
  i, err := strconv.ParseInt(bin, 2, 64)
  
  if err != nil {
    return 0
  }
  return int(i)
}
​