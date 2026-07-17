package kata
​
func SimpleMultiplication(n int) (res int) {
  
  switch {
    case n % 2 == 0:
      res = n * 8
    default:
    res = n * 9
  }
  return
}