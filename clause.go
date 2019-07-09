package gocy

import (
	"fmt"
	"strings"
)

type clause struct {
	queryString string
	returns     map[string]struct{}
}

func (m clause) Node(i interface{}) clause {
	query, structName := parseStruct(i, "(%s)")
	m.queryString += query
	m.returns[structName] = struct{}{}
	return m
}

func (m clause) Rel(i interface{}) clause {
	query, structName := parseStruct(i, "[%s]")
	m.queryString += query
	m.returns[structName] = struct{}{}
	return m
}

func (m clause) L() clause {
	m.queryString += "-"
	return m
}

func (m clause) LV() clause {
	m.queryString += "->"
	return m
}

func (m clause) VL() clause {
	m.queryString += "<-"
	return m
}

func (m clause) LL() clause {
	m.queryString += "--"
	return m
}

func (m clause) VLL() clause {
	m.queryString += "<--"
	return m
}

func (m clause) LLV() clause {
	m.queryString += "-->"
	return m
}

func (m clause) Return() clause {
	var returns []string
	for key := range m.returns {
		returns = append(returns, key)
	}
	m.queryString += fmt.Sprintf("\nRETURN %s", strings.Join(returns, ","))

	return m
}

func (m clause) Limit(i int64) clause {
	m.queryString += fmt.Sprintf("\nLIMIT %d", i)
	return m
}

func (m clause) String() string {
	return m.queryString
}
