package kata
​
func PowersOfTwo(n int) []uint64 {
  
  res := make([]uint64, 0, n+1)
  
  for i := 0; i <= n; i++ {
      res = append(res, uint64(1)<<i)
  }
  
  return res
}