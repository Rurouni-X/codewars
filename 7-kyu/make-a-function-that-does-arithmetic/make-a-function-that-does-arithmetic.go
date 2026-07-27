package kata
​
func Arithmetic(a int, b int, operator string) (res int){
  
  switch operator {
    case "add":
      res = a + b
    case "subtract":
      res = a - b
    case "multiply":
      res = a * b
    case "divide": {
      if b == 0 {
        panic("no divode by zero")
      }
      res = a / b
    }
  }
  return
}