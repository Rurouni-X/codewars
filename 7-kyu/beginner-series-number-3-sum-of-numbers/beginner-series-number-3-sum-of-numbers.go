package kata
​
​
func GetSum(a, b int) (res int) {
  
  if a > b {
    a, b = b, a
  }
  
  for i := a; i <= b; i++ {
    res += i
  }
  return
}