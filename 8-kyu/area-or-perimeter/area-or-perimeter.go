package kata
​
func AreaOrPerimeter(l, w int) int {
  
  if l == w {
    return l * w
  }
  
  return (l * 2) + (w * 2)
}