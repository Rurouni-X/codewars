package kata
​
func Evaporator(content float64, evapPerDay int, threshold int) (res int) {
  
  currentGas := 100.0
  th := float64(threshold)
  epd := float64(evapPerDay) / 100.0
  
  for currentGas > th {
    currentGas -= currentGas * epd
    res += 1
  }
  return 
}