package kata
​
import "fmt"
​
func MultiTable(number int) (res string) {
  
  for i := 1; i <= 10; i++ {
    
    if i != 10 {
     res += fmt.Sprintf("%d * %d = %d\n", i, number, i*number) 
    } else {
      res += fmt.Sprintf("%d * %d = %d", i, number, i*number)
    }
  }
  return
}