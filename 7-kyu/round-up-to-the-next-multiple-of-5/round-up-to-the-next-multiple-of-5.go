package kata
​
func RoundToNext5(n int) int {
  
  remainder := n % 5
  
  if remainder > 0 {
    return n - remainder + 5
  }
  
  return n - remainder
}
​