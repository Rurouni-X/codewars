package kata
​
func HowMuchILoveYou(i int) (res string) {
  
  
  switch {
    case i % 6 == 0:
      res = "not at all"
    case i % 6 == 1:
      res = "I love you"
    case i % 6 == 2:
      res = "a little"
    case i % 6 == 3:
      res = "a lot"
    case i % 6 == 4:
      res = "passionately"
    case i % 6 == 5:
      res = "madly"
  }
  return
}