package main
import (
   "bufio"
   "fmt"
   "os"
   "sort"
   "strings"
)
func main() {
   var cowsMap = map[string][]string{}
   var cows = []string{"Beatrice", "Sue", "Belinda", "Bessie", "Betsy", "Blue", "Bella", "Buttercup"}
   sort.Slice(cows, func(i, j int) bool {
      return cows[i] < cows[j]
   })
   for _, cow := range cows {
      cowsMap[cow] = []string{}
   }
   var read = bufio.NewReader(os.Stdin)
   var n int
   fmt.Fscanln(read, &n)
   for i := 0; i < n; i++ {
      text, _, _ := read.ReadLine()
      command := strings.Split(string(text), " ")
      proCow := command[0]
      postCow := command[5]
      val, _ := cowsMap[postCow]
      if len(val) > 0 {
         if proCow > val[0] {
            val = append(val, proCow)
         } else {
            val = append([]string{proCow}, val...)
         }
         cowsMap[postCow] = val
      } else {
         cowsMap[postCow] = append(cowsMap[postCow], proCow)
      }
      delete(cowsMap, proCow)
   }
   var result [][]string
   for key, val := range cowsMap {
      if len(val) == 0 {
         result = append(result, []string{key})
         continue
      }
      result = append(result, getCowsSlice(key, val))
   }
   sort.SliceStable(result, func(i, j int) bool {
      return result[i][0] < result[j][0]
   })
   for _, prop := range result {
      for _, p := range prop {
         fmt.Println(p)
      }
   }
}
func getCowsSlice(key string, val []string) []string {
   if len(val) == 2 {
      val = []string{val[0], key, val[1]}
   }
   if len(val) == 1 {
      if key > val[0] {
         val = []string{val[0], key}
      } else {
         val = []string{key, val[0]}
      }
   }
   return val

}
