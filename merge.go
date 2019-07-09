package gocy

// Starts a Merge clause
func Merge() clause {
	return clause{
		queryString: "\nMERGE ",
		returns:     map[string]struct{}{},
	}
}
