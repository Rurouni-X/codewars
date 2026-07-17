package kata
​
​
func GetSum(a, b int) (res int) {
  
  if b == a {
    return a
  }
  
  var x int
  var z int
  
  if a < b {
    x, z = a, b
  } else {
    x, z = b, a
  }
  
  for i := x; i <= z; i++ {
    res += i
  }
  return
}