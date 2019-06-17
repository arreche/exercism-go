package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var used = map[string]bool{"": true}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// A Robot holds a robot name.
type Robot struct {
	name string
}

// Name returns robot name in the format of two uppercase letters followed by three digits, such as RX837 or BC811.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(used) > 26 * 26 * 10 * 10 * 10 {
		return "", errors.New("no more names available in AA000 - ZZ999")
	}
	for used[r.name] {
		r.name = generateRandomName()
	}
	used[r.name] = true
	return r.name, nil
}

// Reset resets a robot to its factory settings wipping its name.
func (r *Robot) Reset() {
	r.name = ""
}

func generateRandomName() string {
	r1 := string(rand.Intn(26) + 'A')
	r2 := string(rand.Intn(26) + 'A')
	num := rand.Intn(1000)

	return fmt.Sprintf("%s%s%03d", r1, r2, num)
}
