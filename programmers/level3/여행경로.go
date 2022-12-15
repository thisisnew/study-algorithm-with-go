package main

func main() {

}

var (
	isVisitedTicket []bool
	graphTickets    [][]string
)

func 여행경로(tickets [][]string) []string {
	isVisitedTicket = make([]bool, len(tickets)+1)
	graphTickets = make([][]string, len(tickets)+1)

	for i := range graphTickets {
		graphTickets[i] = make([]string, len(tickets)+1)
	}

	//for _, ticket := range tickets {
	//	graphTickets[]
	//}

	return []string{}
}

func dfsTickets(v int) {
	isVisitedTicket[v] = true
}
