package kata
​
func GetSize(w,h,d int) [2]int {
  square := 2 * (w * h + w * d + h * d)
  deep := w * h * d
  return [2]int{square, deep}
}