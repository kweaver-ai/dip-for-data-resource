package client

import "strings"

// Counter 计数器
type Counter struct {
	infos map[string]map[string]any
}

func NewCounter(words ...string) *Counter {
	return &Counter{
		infos: make(map[string]map[string]any),
	}
}

func (r *Counter) Record(key, value string) {
	f, ok := r.infos[key]
	if !ok {
		f = make(map[string]any)
	}
	f[value] = 1
	r.infos[key] = f
}

func (r *Counter) Count(key string) int {
	return len(r.infos[key])
}

func (r *Counter) ValueKeys(key string) []string {
	words := make([]string, 0)
	filter, ok := r.infos[key]
	if !ok {
		return words
	}
	for key := range filter {
		words = append(words, key)
	}
	return words
}

func (r *Counter) Keys() []string {
	words := make([]string, 0)
	for key := range r.infos {
		words = append(words, key)
	}
	return words
}

func (r *Counter) ContainsCount(word string) int {
	dict := make(map[string]any)
	for key, filter := range r.infos {
		if strings.Contains(word, key) && key != word {
			for key := range filter {
				dict[key] = 1
			}
		}
	}
	return len(dict)
}
