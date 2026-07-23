package kata
​
func SmallestIntegerFinder(numbers []int) int {
  
  res := numbers[0]
  
  for i := 1; i < len(numbers); i++ {
    
    if res > numbers[i] {
      res = numbers[i]
    }
  }
  return res
}