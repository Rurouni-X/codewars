package kata
​
func FindOdd(seq []int) int {
  mp := make(map[int]int)
  
  for _, num := range seq {
    mp[num]++
  }
  
  for key, val := range mp {
    if val % 2 != 0 {
      return key
    }
  }
  return 0
}