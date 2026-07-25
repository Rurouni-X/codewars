package kata
​
​
func SumDigits(number int) int {
  
  if number < 0 {
    number = -number
  }
  sum := 0
  
  for number != 0 {
    sum += number % 10
    number /= 10
  }
  return sum
}
​