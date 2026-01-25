package wordlist

import "math/rand"

type WordList struct {
	Name  string
	Words []string
}

func New(name string, words []string) *WordList {
	return &WordList{
		Name:  name,
		Words: words,
	}
}

func (wl *WordList) GetRandomWords(n int) []string {
	if n <= 0 || len(wl.Words) == 0 {
		return nil
	}

	result := make([]string, n)
	for i := range n {
		result[i] = wl.Words[rand.Intn(len(wl.Words))]
	}
	return result
}
