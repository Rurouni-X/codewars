package kata
​
func solution(str, ending string) bool {
  
  if len(ending) > len(str) {
    return false
  }
  
  ln := len(str) - len(ending)
  return str[ln:] == ending
}