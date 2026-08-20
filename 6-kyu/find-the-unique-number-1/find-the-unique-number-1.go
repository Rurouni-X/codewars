package kata
​
func FindUniq(arr []float32) float32 {
  
  mp := make(map[float32]float32)
  
  for _, val := range arr {
    mp[val]++
  }
  
  for key, value := range mp {
    
    if value == 1 {
      return key
    }
  }
  return 0
}