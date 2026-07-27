package kata
​
func NoSpace(word string) string {
  
  arr := []rune(word)
  var str string
  
  for _, char := range arr {
    
    if char == ' ' {
      continue
    }
    str += string(char)
  }
  return str
}