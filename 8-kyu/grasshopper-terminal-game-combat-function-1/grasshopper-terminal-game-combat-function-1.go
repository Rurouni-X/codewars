package kata
​
func combat(health, damage float64) float64 {
​
  res := health - damage
  
  if res <= 0 {
    return 0
  }
  return res
}