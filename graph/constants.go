package graph

var adjacencyList map[string][]string = map[string][]string{
	"a": {"b", "c"},
	"b": {"d"},
	"c": {"e"},
	"d": {"f"},
	"e": {},
	"f": {},
}
