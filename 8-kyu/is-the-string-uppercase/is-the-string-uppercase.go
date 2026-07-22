package kata
​
import "unicode"
​
type MyString string
​
func (s MyString) IsUpperCase() bool {
​
  for _, v := range string(s) {
    
    if unicode.IsLower(v) {
      return false
    }
  }
  return true
}