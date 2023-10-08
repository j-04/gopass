package searcher

type Searcher interface {
	find(line string, cname string) (bool, error)
}

type CustomSearcher struct{}

func NewCustomSearcher() *CustomSearcher {
	return &CustomSearcher{}
}

func find(line string, cname string) (bool, error) {
	return false, nil
}
