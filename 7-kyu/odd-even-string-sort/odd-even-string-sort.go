package kata
​
​
func SortMyString(s string) string {
  odd := ""
  even := ""
  
  for idx, char := range s {  
    if idx % 2 == 0 {
      odd += string(char)
    }else {
      even += string(char)
    }
  }
  
  return odd + " " + even
}
​