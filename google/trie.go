type node struct {
    child []*node
    val bool
}

func NewNode() *node {
    return &node{
        child: make([]*node, 26),
    }
}

type Trie struct {
    root *node
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    this.root = this.insert(this.root, word, 0)
}

func (this *Trie) insert(x *node, word string, d int) *node {
    if x == nil {
        x = NewNode()
    }
    if d == len(word) {
        x.val = true
        return x
    }
    
    c := word[d]-'a'
    x.child[c] = this.insert(x.child[c], word, d+1)
    return x
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    x := this.search(this.root, word, 0)
    return x != nil && x.val 
}

func (this *Trie) search(x *node, word string, d int) *node {
    if x == nil {
        return nil
    }
    
    if d == len(word) {
        return x
    }
    
    c := word[d]-'a'
    return this.search(x.child[c], word, d+1)
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    x := this.search(this.root, prefix, 0)
    return x != nil
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */