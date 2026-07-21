package kata
​
​
func Invert(arr []int) (res []int) {
  
  for _, v := range arr {
    res = append(res, -v)
  }
  return 
}
​