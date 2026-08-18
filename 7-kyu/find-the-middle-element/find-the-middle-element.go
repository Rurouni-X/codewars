package kata
​
​
func Gimme(array [3]int) int {
  
  min := array[0]
  max := array[0]
  
  for _, v := range array {
    
    if v < min {
      min = v
    }
    if v > max {
      max = v
    }
    
  }
  
  for idx, val := range array {
    
    if val != min && val != max {
      return idx
    }
  }
  return 0
}
​