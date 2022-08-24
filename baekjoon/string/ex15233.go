package main
import (
   "bufio"
   "fmt"
   "os"
   "strings"
)
var dictionary map[string]string
func main() {
   var read = bufio.NewReader(os.Stdin)
   var a, b, g int
   fmt.Fscanln(read, &a, &b, &g)
   var aTeam = getTeams(read)
   var bTeam = getTeams(read)
   aGoals, bGoals := countGoals(read, aTeam, bTeam)
   judgeTheWinner(aGoals, bGoals)
}
func getTeams(read *bufio.Reader) map[string]bool {
   var team = map[string]bool{}
   text, _, _ := read.ReadLine()
   players := strings.Split(string(text), " ")
   for _, ap := range players {
      team[ap] = true
   }
   return team
}
func countGoals(read *bufio.Reader, aTeam, bTeam map[string]bool) (int, int) {
   text, _, _ := read.ReadLine()
   strikers := strings.Split(string(text), " ")
   var aGoals int
   var bGoals int
   for _, st := range strikers {
      _, ok := aTeam[st]
      if ok {
         aGoals++
         continue
      }
      _, ok = bTeam[st]
      if ok {
         bGoals++
         continue
      }
   }
   return aGoals, bGoals
}
func judgeTheWinner(aGoals, bGoals int) {
   switch {
   case aGoals > bGoals:
      fmt.Println("A")
   case aGoals < bGoals:
      fmt.Println("B")
   default:
      fmt.Println("TIE")
   }

}
