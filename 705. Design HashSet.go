package main

/**
用map实现
*/
type MyHashSet struct {
	s map[int]int
}

func Constructor() MyHashSet {
	ss := make(map[int]int)
	h := MyHashSet{
		s: ss,
	}
	return h
}

func (this *MyHashSet) Add(key int) {
	this.s[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	delete(this.s, key)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.s[key]
	return ok
}
