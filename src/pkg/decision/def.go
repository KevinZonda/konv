package decision

import (
	"errors"
)

type CheckTree struct {
	Key      string
	Data     *string
	Children map[string]*CheckTree
}

func (t *CheckTree) AddKey(key string) error {
	return t.Add(&CheckTree{Key: key})
}

func (t *CheckTree) Add(v *CheckTree) error {
	if t == nil {
		return errors.New("empty tree")
	}
	if t.Children == nil {
		t.Children = make(map[string]*CheckTree)
	}
	t.Children[v.Key] = v
	return nil
}

func (t *CheckTree) AddToSub(keys []string, data *string) {
	curr := t
	maxIndex := 0
	for i, key := range keys {
		now, ok := t.Children[key]
		if ok {
			curr = now
			maxIndex = i
			continue
		}
	}
	keysN := keys[maxIndex+1:]
	if len(keysN) < 1 {
		return
	}
	cons := curr
	for _, key := range keysN {
		_ = cons.AddKey(key)
		cons = curr.Children[key]
	}
	cons.Data = data
}

func (t *CheckTree) Get(keys []string) *CheckTree {
	return t.get(keys, true)
}

func (t *CheckTree) get(keys []string, isfirst bool) *CheckTree {
	if t == nil {
		return nil
	}
	if len(keys) == 0 {
		return t
	}
	if isfirst {
		if t.Key != keys[0] {
			return nil
		}
		keys = keys[1:]
	}
	child, ok := t.Children[keys[0]]
	if !ok {
		return nil
	}
	return child.get(keys[1:], false)
}
