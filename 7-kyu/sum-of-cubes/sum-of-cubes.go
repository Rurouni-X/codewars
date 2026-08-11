package kata
​
func SumCubes(n int) (res int) {
  
  for i := 1; i <= n; i++ {
    cube := i * i * i
    res += cube
  }
  return
}