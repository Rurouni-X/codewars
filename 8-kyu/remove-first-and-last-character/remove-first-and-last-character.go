package kata
​
func RemoveChar(word string) string {
  
  if len(word) == 2 {
    return ""
  }
  res := word[1: len(word)- 1]
  return res
}