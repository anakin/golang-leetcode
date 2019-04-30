package main

/**
这里实现了一个trie数
*/
type Trie struct {
	char     byte
	end      bool
	value    int
	children map[byte]*Trie
}

func (t *Trie) Insert(key string, val, del int) {
	if key == "" {
		return
	}
	t.value += val - del
	if len(key) == 1 {
		t.end = true
		return
	}
	ch, ok := t.children[key[1]]
	if !ok {
		ch = &Trie{char: key[1], children: map[byte]*Trie{}}
		t.children[key[1]] = ch
	}
	(*ch).Insert(key[1:], val, del)
}
func (t *Trie) Find(key string) (int, bool) {
	if key == "" {
		return 0, false
	}
	if len(key) == 1 {
		return t.value, t.end
	}
	ch, ok := t.children[key[1]]
	if !ok {
		return 0, false
	}
	return ch.Find(key[1:])
}

type MapSum struct {
	tries map[byte]*Trie
}

func Constructor() MapSum {
	return MapSum{map[byte]*Trie{}}
}
func (this *MapSum) Insert(key string, val int) {
	if t, ok := this.tries[key[0]]; ok {
		prev, ok := (*t).Find(key)
		if !ok {
			prev = 0
		}
		(*t).Insert(key, val, prev)
	} else {
		t = &Trie{char: key[0], children: map[byte]*Trie{}}
		t.char = key[0]
		this.tries[key[0]] = t
		(*t).Insert(key, val, 0)
	}
}

func (this *MapSum) Sum(prefix string) int {
	if prefix == "" {
		return 0
	}
	t, ok := this.tries[prefix[0]]
	if !ok {
		return 0
	}
	val, _ := t.Find(prefix)
	return val
}
