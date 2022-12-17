package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type planet struct {
	name     string
	children []*planet
	parent   *planet
	distance float64
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	planets := map[string]*planet{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		planetInfo := strings.Split(scanner.Text(), ")")
		parentName, childName := planetInfo[0], planetInfo[1]

		child := planets[childName]
		if child == nil {
			child = newPlanet(childName)
		}

		parent := planets[parentName]
		if parent == nil {
			parent = newPlanet(parentName)
		}

		parent.addChild(child)

		child.parent = parent

		planets[parent.name] = parent
		planets[child.name] = child
	}

	fmt.Println("part 1:", planets["COM"].findAllOrbits())

	fmt.Println("part 2:", distance(planets))

}

func distance(planets map[string]*planet) int {
	start := planets["YOU"].parent
	start.distance = 0

	end := planets["SAN"].parent

	for len(planets) > 0 {
		lowest := min(planets)
		if lowest == end {
			return int(lowest.distance)
		}
		delete(planets, lowest.name)

		tmp := lowest.children
		if lowest.parent != nil {
			tmp = append(lowest.children, lowest.parent)
		}

		for _, x := range tmp {
			_, ok := planets[x.name]
			if ok {
				alt := lowest.distance + 1
				if alt < x.distance {
					x.distance = alt
				}
			}
		}
	}

	return 0
}

func min(planets map[string]*planet) *planet {
	lowest := math.Inf(1)
	planet := &planet{}

	for _, v := range planets {
		if v.distance < lowest {
			lowest = v.distance
			planet = v
		}
	}

	return planet
}

func newPlanet(name string) *planet {
	p := planet{
		name:     name,
		children: []*planet{},
		distance: math.Inf(1),
	}
	return &p
}

func (p *planet) findAllOrbits() int {
	total := p.findOrbits()
	for _, c := range p.children {
		total += c.findAllOrbits()
	}
	return total
}

func (p *planet) findOrbits() int {
	count := 0
	for _, c := range p.children {
		count += c.findOrbits()
		count += 1
	}
	return count
}

func (p *planet) addChild(c *planet) {
	p.children = append(p.children, c)
}
