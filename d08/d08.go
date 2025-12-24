package main

import (
	"aoc/utils"
	"cmp"
	"fmt"
	"log"
	"maps"
	"math"
	"slices"
	"strings"
	"time"
)

type coord struct {
	x int
	y int
	z int
}

type box struct {
	coord     coord
	networkId int
}

type boxDist struct {
	b1   *box
	b2   *box
	dist float64
}

func (b box) String() string {
	return fmt.Sprintf("%d,%d,%d", b.coord.x, b.coord.y, b.coord.z)
}

func distance(a, b coord) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

func p1(data []string, test bool) {

	start := time.Now()
	ans := 0
	connections := 1000
	if test {
		connections = 10
	}
	networkId := 0
	boxes := make([]*box, len(data))
	for ind, line := range data {
		parts := strings.Split(line, ",")
		c := coord{
			x: utils.Atoi(parts[0]),
			y: utils.Atoi(parts[1]),
			z: utils.Atoi(parts[2]),
		}
		b := box{
			coord:     c,
			networkId: -1,
		}
		boxes[ind] = &b
	}
	combs := utils.Combinations(boxes)
	distances := make([]boxDist, len(combs))
	for ind, c := range combs {
		distances[ind] = boxDist{
			b1:   c[0],
			b2:   c[1],
			dist: distance(c[0].coord, c[1].coord),
		}
	}
	slices.SortStableFunc(distances, func(a, b boxDist) int { return cmp.Compare(a.dist, b.dist) })

	connectionCount := 0
	connected := map[int][]*box{}
	for _, dist := range distances {
		if connectionCount >= connections {
			break
		}
		b1 := dist.b1
		b2 := dist.b2
		if b1.networkId == -1 && b2.networkId == -1 {
			log.Printf("creating network between %v and %v network %v", b1, b2, networkId)
			b1.networkId = networkId
			b2.networkId = networkId
			connected[networkId] = []*box{b1, b2}
			networkId++
		} else if b1.networkId == -1 || b2.networkId == -1 {
			connectedBox := b1
			otherBox := b2
			if b1.networkId == -1 {
				connectedBox = b2
				otherBox = b1
			}
			log.Printf("adding %v to %v's existing network %v", otherBox, connectedBox, connectedBox.networkId)
			otherBox.networkId = connectedBox.networkId
			connected[connectedBox.networkId] = append(connected[connectedBox.networkId], otherBox)
		} else if b1.networkId != b2.networkId {
			log.Printf("updating %v's network %v to %v's network %v", b1, b1.networkId, b2, b2.networkId)
			newId := b2.networkId
			oldId := b1.networkId
			for _, update := range connected[oldId] {
				log.Printf("updating %v", update)
				update.networkId = newId
			}
			connected[newId] = append(connected[newId], connected[oldId]...)
			connected[oldId] = nil
		} else {
			if (b1.networkId != b2.networkId) || b1.networkId == -1 || b2.networkId == -1 {
				panic("panik")
			}
			log.Printf("skipping %v and %v, already connected", b1, b2)
		}
		connectionCount++
	}

	sizes := []int{}
	for _, c := range connected {
		sizes = append(sizes, len(c))
	}
	slices.SortStableFunc(sizes, func(a, b int) int { return cmp.Compare(b, a) })
	ans = sizes[0] * sizes[1] * sizes[2]

	log.Printf("p1 ans %v", ans)
	log.Printf("p1 took %s", time.Since(start))
}

func p2(data []string) {

	start := time.Now()
	ans := 0
	networkId := 0
	boxes := make([]*box, len(data))
	for ind, line := range data {
		parts := strings.Split(line, ",")
		c := coord{
			x: utils.Atoi(parts[0]),
			y: utils.Atoi(parts[1]),
			z: utils.Atoi(parts[2]),
		}
		b := box{
			coord:     c,
			networkId: -1,
		}
		boxes[ind] = &b
	}
	combs := utils.Combinations(boxes)
	distances := make([]boxDist, len(combs))
	for ind, c := range combs {
		distances[ind] = boxDist{
			b1:   c[0],
			b2:   c[1],
			dist: distance(c[0].coord, c[1].coord),
		}
	}
	slices.SortStableFunc(distances, func(a, b boxDist) int { return cmp.Compare(a.dist, b.dist) })

	connected := map[int][]*box{}
	var b1, b2 *box

	for _, dist := range distances {
		b1 = dist.b1
		b2 = dist.b2
		if b1.networkId == -1 && b2.networkId == -1 {
			log.Printf("creating network between %v and %v network %v", b1, b2, networkId)
			b1.networkId = networkId
			b2.networkId = networkId
			connected[networkId] = []*box{b1, b2}
			networkId++
		} else if b1.networkId == -1 || b2.networkId == -1 {
			connectedBox := b1
			otherBox := b2
			if b1.networkId == -1 {
				connectedBox = b2
				otherBox = b1
			}
			log.Printf("adding %v to %v's existing network %v", otherBox, connectedBox, connectedBox.networkId)
			otherBox.networkId = connectedBox.networkId
			connected[connectedBox.networkId] = append(connected[connectedBox.networkId], otherBox)
		} else if b1.networkId != b2.networkId {
			log.Printf("updating %v's network %v to %v's network %v", b1, b1.networkId, b2, b2.networkId)
			newId := b2.networkId
			oldId := b1.networkId
			for _, update := range connected[oldId] {
				log.Printf("updating %v", update)
				update.networkId = newId
			}
			connected[newId] = append(connected[newId], connected[oldId]...)
			connected[oldId] = nil
			delete(connected, oldId)
		} else {
			if (b1.networkId != b2.networkId) || b1.networkId == -1 || b2.networkId == -1 {
				panic("panik")
			}
			log.Printf("skipping %v and %v, already connected", b1, b2)
		}
		if len(connected) == 1 && len(connected[slices.Collect(maps.Keys(connected))[0]]) == len(data) {
			break
		}
	}
	ans = b1.coord.x * b2.coord.x
	log.Printf("p2 ans %v", ans)
	log.Printf("p2 took %s", time.Since(start))
}

func main() {
	test := false
	day := "d08"

	data, err := utils.ReadFile(day, test)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	p1(data, test)
	p2(data)
}
