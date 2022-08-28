package rule

import (
	"fmt"
	"strconv"
	"strings"
)

type RuleStruct struct {
	Born    []int
	Survive []int
}

func (rule RuleStruct) String() string {
	var buf strings.Builder
	buf.WriteString("B")
	for i, v := range rule.Born {
		buf.WriteString(fmt.Sprintf("%d", v))
		if i < len(rule.Born)-1 {
			buf.WriteString(",")
		}
	}

	buf.WriteString("/S")
	for i, v := range rule.Survive {
		buf.WriteString(fmt.Sprintf("%d", v))
		if i < len(rule.Survive)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func ParseRule(s string) (RuleStruct, error) {
	born := make([]int, 0)
	survive := make([]int, 0)

	bits := strings.Split(s, "/")

	for _, v := range bits[0][1:] {
		x, err := strconv.Atoi(string(v))
		if err != nil {
			return RuleStruct{}, err
		}
		born = append(born, x)
	}

	for _, v := range bits[1][1:] {
		x, err := strconv.Atoi(string(v))
		if err != nil {
			return RuleStruct{}, err
		}
		survive = append(survive, x)
	}

	return RuleStruct{Born: born, Survive: survive}, nil
}

var ReplicatorRule = RuleStruct{Born: []int{1, 3, 5, 7}, Survive: []int{1, 3, 5, 7}}
var ConwaywayRule = RuleStruct{Born: []int{3}, Survive: []int{2, 3}}
var PedestrianRule = RuleStruct{Born: []int{3, 8}, Survive: []int{2, 3}}
var FlockRule = RuleStruct{Born: []int{3}, Survive: []int{1, 2}}

var Patterns = map[string]RuleStruct{
	"conway":     ConwaywayRule,
	"replicator": ReplicatorRule,
	"pedestrian": PedestrianRule,
	"flock":      FlockRule,
}
