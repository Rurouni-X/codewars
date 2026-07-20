package kata
​
import "unicode"
​
func ToAlternatingCase(str string) (res string) {
  
  for _, char := range str {
    
    if unicode.IsUpper(char) {
      res += string(unicode.ToLower(char))
    } else {
      res += string(unicode.ToUpper(char))
    }
  }
  return
}