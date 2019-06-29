package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type entry struct {
	team    string
	matches int
	won     int
	draw    int
	lost    int
	points  int
}

// Tally computes the results of a small football competition.
func Tally(r io.Reader, w io.Writer) error {
	teams := map[string]entry{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		chunks := strings.Split(line, ";")
		if len(chunks) != 3 {
			return fmt.Errorf("wrong line: %s", line)
		}

		t1 := teams[chunks[0]]
		t1.team = chunks[0]
		t1.matches++

		t2 := teams[chunks[1]]
		t2.team = chunks[1]
		t2.matches++

		switch chunks[2] {
		case "win":
			t1.won++
			t1.points += 3
			t2.lost++
		case "loss":
			t1.lost++
			t2.won++
			t2.points += 3
		case "draw":
			t1.draw++
			t1.points++
			t2.draw++
			t2.points++
		default:
			return fmt.Errorf("wrong result")
		}

		teams[chunks[0]] = t1
		teams[chunks[1]] = t2
	}

	var a []entry
	for _, v := range teams {
		a = append(a, v)
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i].points == a[j].points {
			return a[i].team < a[j].team
		}
		return a[i].points > a[j].points
	})

	fmt.Fprintf(w, "%-31s| %s |  %s |  %s |  %s |  %s\n", "Team", "MP", "W", "D", "L", "P")

	for _, v := range a {
		fmt.Fprintf(w, "%-31s|  %d |  %d |  %d |  %d |  %d\n", v.team, v.matches, v.won, v.draw, v.lost, v.points)
	}

	return nil
}
