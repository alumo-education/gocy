package gocy

func Match() clause {
	return clause{
		queryString: "\nMATCH ",
		returns:     map[string]struct{}{},
	}
}
