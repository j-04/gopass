package core

type Searcher interface {
	Find(pattern, source string) (bool, error)
}

type LevenshteinSearcher struct {
	threshold int
}

func NewLevenshteinSearcher(threshold int) *LevenshteinSearcher {
	return &LevenshteinSearcher{
		threshold: threshold,
	}
}

func (searcher *LevenshteinSearcher) Find(pattern, source string) (bool, error) {
	distance := levenshtein(pattern, source)
	if distance <= searcher.threshold {
		return true, nil
	}
	return false, nil
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

func levenshtein(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	matrix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 {
				matrix[i][j] = j
			} else if j == 0 {
				matrix[i][j] = i
			} else {
				cost := 0
				if s1[i-1] != s2[j-1] {
					cost = 1
				}
				matrix[i][j] = min(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1]+cost)
			}
		}
	}

	return matrix[m][n]
}
