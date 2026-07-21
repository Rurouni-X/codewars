package kata
​
​
func Between(a, b int) (res []int) {
  
  for i := a; i <= b; i++ {
    res = append(res, i)
  }
  return
}
​