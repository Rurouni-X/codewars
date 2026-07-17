package kata
​
func QuarterOf(month int) (res int) {
  
  switch month {
    case 1,2,3:
      res = 1
    case 4,5,6:
      res = 2
    case 7,8,9:
      res = 3
    case 10,11,12:
      res = 4
    default:
    res = 0
  }
  return
}