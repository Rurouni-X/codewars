package kata
​
​
func SequenceSum(start, end, step int) (res int) {
  
  if start > end {
    res = 0
  }
  
  for i := start; i <= end; i += step {
    res += i
  }
  return
}
​