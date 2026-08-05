package kata
​
func GetMiddle(s string) string {
  
  n := len(s)
  
  mid := n / 2
  
  if n%2 == 0 {
    return s[mid-1 : mid+1]
  }
  return s[mid : mid+1]
}