package gocy

func Merge() clause {
	return clause{
		queryString: "\nMERGE ",
		returns: map[string]struct{}{},
	}
}
