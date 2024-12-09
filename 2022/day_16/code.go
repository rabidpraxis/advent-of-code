package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rabidpraxis/advent-of-code/utils"
)

// type run struct {
// 	score   int
// 	visited utils.OccurrenceSet[string]
// 	opened  utils.Set[string]
// }

// func newRun() run {
// 	return run{
// 		visited: *utils.NewOccurrenceSet[string](),
// 		opened:  *utils.NewSet[string](),
// 	}
// }

// func (r *run) clone() run {
// 	return run{
// 		score:   r.score,
// 		visited: *r.visited.Clone(),
// 		opened:  *r.opened.Clone(),
// 	}
// }

// func (r *run) canVisit(tunnel string) bool {
// 	ct, _ := r.visited.Get(tunnel)
// 	if ct > 2 {
// 		return false
// 	}

// 	return true
// }

// func checkValve(t int, currTunnel string, r run) run {
// 	if t == 0 {
// 		return r
// 	}

// 	r.visited.Add(currTunnel)
// 	nt := t - 1

// 	// fmt.Println(r.path)

// 	v, _ := valves[currTunnel]
// 	localRuns := []run{}

// 	// With opening
// 	if !r.opened.Has(currTunnel) && nt > 0 {
// 		nr := r.clone()
// 		ntt := nt - 1
// 		nr.score += v.rate * ntt
// 		nr.opened.Add(currTunnel)

// 		for _, tunnel := range v.tunnels {
// 			if nr.canVisit(tunnel) {
// 				localRuns = append(localRuns, checkValve(ntt, tunnel, nr))
// 			}
// 		}
// 	}

// 	for _, tunnel := range v.tunnels {
// 		nr := r.clone()
// 		// Without opening
// 		if nr.canVisit(tunnel) {
// 			localRuns = append(localRuns, checkValve(nt, tunnel, nr))
// 		}
// 	}

// 	if len(localRuns) > 0 {
// 		sort.Slice(localRuns, func(i, j int) bool {
// 			return localRuns[i].score > localRuns[j].score
// 		})

// 		// fmt.Println(localRuns[0].score)

// 		return localRuns[0]
// 	} else {
// 		return r
// 	}
// }

var (
	re     = regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (.*)`)
	valves = map[string]valve{}
)

func loadValves(lines []string) {
	for _, line := range lines {
		m := re.FindStringSubmatch(line)
		rate, _ := strconv.Atoi(m[2])

		valves[m[1]] = valve{
			name:    m[1],
			rate:    rate,
			tunnels: strings.Split(m[3], ", "),
		}
	}
}

type valve struct {
	id    string
	rate  int
	paths []string
}

func (n node) PathNeighbors() []utils.Pather {
	ps := []utils.Pather{}

	if !n.open.Has(n.valve) && !n.secondVisit {
		ps = append(ps, node{
			valve:       n.valve,
			minute:      n.minute - 1,
			secondVisit: true,
			open:        n.open,
		})
	}

	v, _ := valves[n.valve]

	for _, t := range v.tunnels {
		ps = append(ps, node{
			valve:       t,
			minute:      n.minute - 1,
			secondVisit: false,
			open:        n.open,
		})
	}

	return ps
}

func (n node) PathNeighborCost(to utils.Pather) float64 {
	if !n.open.Has(n.valve) {
		v, _ := valves[n.valve]

		score := float64(v.rate * n.minute)
		fmt.Println("here", v, n.valve, n.minute, score)
		n.open.Add(n.valve)
		return score
	} else {
		return 0
	}
}

func (n node) PathEstimatedCost(to utils.Pather) float64 {
	return 0
}

func (n node) PathComplete(to utils.Pather) bool {
	return n.minute == 0
}

type node struct {
	valveId      string
	minute       int
	secondVisit  bool
	cost         float64
	parent       *node
	open         bool
	closed       bool
	openedValves *utils.Set[string]
}

func part1() {
	start := node{
		valve:        "AA",
		minute:       30,
		secondVisit:  false,
		openedValves: utils.NewSet[string](),
	}

	heap := utils.NewHeap[node](func(a, b node) bool {
		return a.cost > b.cost
	})

	fmt.Println(utils.Path(start, node{}))
}

func main() {
	loadValves(utils.FileLines(os.Args[1]))

	part1()
}
