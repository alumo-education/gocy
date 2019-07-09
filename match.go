package gocy

// Starts a Match clause
func Match() clause {
	return clause{
		queryString: "\nMATCH ",
		returns:     map[string]struct{}{},
	}
}
