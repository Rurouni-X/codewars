package kata
​
import "strings"
​
func AbbrevName(name string) string{
  
  str := strings.Split(name, " ")
  
  return strings.ToUpper(string(str[0][0])) +"."+strings.ToUpper(string(str[1][0]))
}