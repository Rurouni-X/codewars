package kata
​
import "strings"
​
func IsPalindrome(str string) bool {
  
  arr := []rune(strings.ToLower(str))
  
  left, right := 0, len(arr) - 1
  
  for left <= right {
    
    if arr[left] != arr[right] {
      return false
    }
    left++
    right--
  }
  return true
}