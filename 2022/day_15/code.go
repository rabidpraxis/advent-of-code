package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/rabidpraxis/advent-of-code/utils"
)

type sensor struct {
	pos     utils.Coord
	beacon  utils.Coord
	dist    utils.Coord
	distMag int
}

func (s *sensor) magDiff(c utils.Coord) int {
	return s.distMag - s.pos.Distance(c).Mag()
}

func (s *sensor) withinDist(c utils.Coord) bool {
	if s.pos.Distance(c).Mag() <= s.distMag {
		return true
	}

	return false
}

func part1(sensors []sensor, yCheck int) {
	posSet := utils.NewSet[int]()
	for _, s := range sensors {
		d := s.magDiff(utils.Coord{s.pos.X, yCheck})
		if d > 0 {
			for i := -d; i < d; i++ {
				posSet.Add(s.pos.X + i)
			}
		}
	}
	fmt.Println(posSet.Length())
}

func part2(sensors []sensor, r int) {
	edges := utils.NewSet[utils.Coord]()
	for _, s := range sensors {
		for i := 0; i <= (s.distMag + 1); i++ {
			inv := (s.distMag + 1) - i

			plusX := s.pos.X + i
			minX := s.pos.X - i
			plusY := s.pos.Y + inv
			minY := s.pos.Y - inv

			if plusX >= 0 && plusX <= r && minY >= 0 && minY <= r {
				edges.Add(utils.Coord{plusX, minY})
			}
			if plusX >= 0 && plusX <= r && plusY >= 0 && plusY <= r {
				edges.Add(utils.Coord{plusX, plusY})
			}
			if minX >= 0 && minX <= r && minY >= 0 && minY <= r {
				edges.Add(utils.Coord{minX, minY})
			}
			if minX >= 0 && minX <= r && plusY >= 0 && plusY <= r {
				edges.Add(utils.Coord{minX, plusY})
			}
		}
	}

	for _, edge := range edges.ToSlice() {
		okay := true
		for _, s := range sensors {
			if s.withinDist(edge) {
				okay = false
			}
		}

		if okay {
			fmt.Println((edge.X * 4000000) + edge.Y)
			return
		}
	}
}

func main() {
	lines := utils.FileLines(os.Args[1])

	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	sensors := []sensor{}

	for _, line := range lines {
		ms := re.FindStringSubmatch(line)
		senX, _ := strconv.Atoi(ms[1])
		senY, _ := strconv.Atoi(ms[2])
		beaX, _ := strconv.Atoi(ms[3])
		beaY, _ := strconv.Atoi(ms[4])

		senPos := utils.Coord{senX, senY}
		beaconPos := utils.Coord{beaX, beaY}
		dist := senPos.Distance(beaconPos)

		sensors = append(sensors, sensor{
			pos:     senPos,
			beacon:  beaconPos,
			dist:    dist,
			distMag: dist.Mag(),
		})
	}

	part1(sensors, 2000000)
	part2(sensors, 4000000)
}
