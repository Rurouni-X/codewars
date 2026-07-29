package kata
​
func FindMultiples(integer, limit int) (res []int) {
  
  for i := 1; i <= limit; i++ {
    
    if i % integer == 0 {
      res = append(res, i)
    }
    
  }
  return
}
​