package main

//TODO 效率待提升
type MyHashMap struct {
	s []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	ss := make([]int, 1000000)
	for i := 0; i < len(ss); i++ {
		ss[i] = -1
	}
	return MyHashMap{
		s: ss,
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.s[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if this.s[key] != -1 {
		return this.s[key]
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.s[key] = -1
}
