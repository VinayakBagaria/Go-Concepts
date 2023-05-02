package graph

var adjacencyList = map[string][]string{
	"a": {"b", "c"},
	"b": {"d"},
	"c": {"e"},
	"d": {"f"},
	"e": {},
	"f": {},
}

var edges = [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}
