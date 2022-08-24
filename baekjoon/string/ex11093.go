package main
import (
   "bufio"
   "fmt"
   "math"
   "os"
   "strings"
)
func main() {
   var read = bufio.NewReader(os.Stdin)
   var n int
   fmt.Fscanln(read, &n)
   for i := 0; i < n; i++ {
      var s string
      fmt.Fscanln(read, &s)
      printTablePivoted(getCharacterTable(s))
   }
}
func getCharacterTable(s string) [][]string {
   var ln = len([]rune(s))
   p := math.Ceil(math.Sqrt(float64(ln)))
   var table = make([][]string, int(p))
   var idx = 0
   for i := 0; i < len(table); i++ {
      var t = make([]string, int(p))
      for j := 0; j < len(t); j++ {
         if idx < ln {
            t[j] = s[idx : idx+1]
            idx++
            continue
         }
         t[j] = "*"
      }
      table[i] = t
   }
   return table
}
func printTablePivoted(table [][]string) {
   var result strings.Builder
   for j := 0; j < len(table); j++ {
      for i := len(table) - 1; i >= 0; i-- {
         s := table[i][j]
         if s == "*" {
            continue
         }
         result.WriteString(s)
      }
   }
   fmt.Println(result.String())

}
