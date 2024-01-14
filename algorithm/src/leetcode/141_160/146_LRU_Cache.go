package _41_160

/**
  @author: CodeWater
  @since: 2024/1/14
  @desc: 146. LRU缓存机制
**/

type LRUCache struct {
	Key, Val    int
	Left, Right *LRUCache
}

// //map记录键值是否存在，双链表保证O（1）删除（最左边是最近用过的，最右边是不常用的值）
var (
	hash = make(map[int]*LRUCache)
	L, R = &LRUCache{Key: -1, Val: -1}, &LRUCache{Key: -1, Val: -1}
	n    int
)

func Constructor(capacity int) LRUCache {
	n = capacity
	/*flag:如果不加下面三行 ，只在全局初始化，力扣的多个用例之间会互相影响，用上一个案例留下来的值
	  ["LRUCache","put","put","get","put","get","put","get","get","get"]
	  [[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	*/
	hash = map[int]*LRUCache{}
	L = &LRUCache{}
	R = &LRUCache{}
	L.Right = R
	R.Left = L
	return LRUCache{}
}

func remove(p *LRUCache) {
	p.Right.Left = p.Left
	p.Left.Right = p.Right
}

// 插入到链表的最左边
func insert(p *LRUCache) {
	p.Right = L.Right
	p.Left = L
	L.Right.Left = p
	L.Right = p
}

func (this *LRUCache) Get(key int) int {
	if _, ok := hash[key]; !ok {
		return -1
	}
	p := hash[key]
	//从双链表中删除，插入到队头
	remove(p)
	insert(p) //到最左边，表示最近用过
	return p.Val
}

func (this *LRUCache) Put(key int, value int) {
	p, ok := hash[key]
	if ok {
		p.Val = value
		remove(p)
		insert(p)
	} else {
		if len(hash) == n {
			//cache容量满了，删除最
			p = R.Left
			remove(p)
			delete(hash, p.Key)
		}
		p = &LRUCache{Key: key, Val: value}
		hash[key] = p
		insert(p)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
