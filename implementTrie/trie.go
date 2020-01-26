package implementTrie

type Trie struct {
	Data   byte
	isWord bool
	Child  []*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	var cur *Trie
	cur = this
	for i := 0; i < len(word); i++ {
		child := cur.Child
		w := word[i]
		temp := &Trie{
			Data:  w,
			Child: []*Trie{},
		}
		for _, c := range child {
			if c.Data == w {
				cur = c
				if i == len(word)-1 {
					cur.isWord = true
				}
				goto NEXT
			}
		}
		if i == len(word)-1 {
			temp.isWord = true
		}
		cur.Child = append(child, temp)
		cur = temp
	NEXT:
	}
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		child := cur.Child
		w := word[i]
		for _, c := range child {
			if c.Data == w {
				cur = c
				goto NEXT
			}
		}
		return false
	NEXT:
	}
	temp := cur.Child
	if len(temp) == 0 {
		return true
	}
	if cur.isWord {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		child := cur.Child
		w := prefix[i]
		for _, c := range child {
			if c.Data == w {
				cur = c
				goto NEXT
			}
		}
		return false
	NEXT:
	}
	return true
}
