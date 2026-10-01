package kata
​
import "slices"
​
func SortNumbers(numbers []int) []int {
  slices.Sort(numbers)
  return numbers
}