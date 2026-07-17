package kata
​
​
func FakeBin(x string) (res string) {
  
  for _, s := range x {
    
    if s < '5' {
      res += "0"
    } else {
      res += "1"
    }
    
  }
  return
}
​