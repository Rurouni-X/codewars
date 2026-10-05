package kata
​
func Add(n int) func(int)int {
    return func (j int) int {
      return n + j
    }
}