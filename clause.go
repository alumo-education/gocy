package gocy

import (
	"fmt"
	"strings"
)

type clause struct {
	queryString string
	returns     map[string]struct{}
	// TODO improve var naming with local map (generate hash from struct)
}

func New() clause {
	return clause{}
}

func (m clause) Node(i interface{}) clause {
	query, structName := parseStructWithName(i, "(%s)")
	m.queryString += query
	m.returns[structName] = struct{}{}
	return m
}

func (m clause) Rel(name string) clause {
	m.queryString += fmt.Sprintf("[:%s]", name)
	return m
}

func (m clause) RelStruct(i interface{}) clause {
	query, structName := parseStructWithName(i, "[%s]")
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

func (m clause) Merge() clause {
	m.queryString += "\nMERGE "
	return m
}

func (m clause) Match() clause {
	m.queryString += "\nMATCH "
	return m
}

func (m clause) Set(i interface{}) clause {
	fields, structName := parseStructFields(i)
	m.queryString += fmt.Sprintf("\nSET %s += %s", structName, fields)
	return m
}

func (m clause) Limit(i int64) clause {
	m.queryString += fmt.Sprintf("\nLIMIT %d", i)
	return m
}

func (m clause) String() string {
	return m.queryString
}
