// generated by yagi. Don't modify this file!
// Any changes will be lost if this file is regenerated.

package list

type ListInt64 struct{ items []int64 }

func (l ListInt64) Items() []int64 {
	return l.items
}

func (l *ListInt64) Add(item int64) {
	l.items = append(l.items, item)
}

func (l *ListInt64) Len() int {
	return len(l.items)
}

type ListInt32 struct{ items []int32 }

func (l ListInt32) Items() []int32 {
	return l.items
}

func (l *ListInt32) Add(item int32) {
	l.items = append(l.items, item)
}

func (l *ListInt32) Len() int {
	return len(l.items)
}
