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
	index := 0
	for i, key := range keys {
		g := curr.GetNode(key)
		if g != nil {
			curr = g
			continue
		}
		index = i
		break
	}
	toAdd := keys[index:]
	for _, key := range toAdd {
		_ = curr.AddKey(key)
		curr = curr.GetNode(key)
	}
	curr.Data = data
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

func (r *CheckTree) GetNode(key string) *CheckTree {
	if r == nil {
		return nil
	}
	c, ok := r.Children[key]
	if !ok {
		return nil
	}
	return c
}

func (r *CheckTree) Next(key string) *CheckTree {
	c, ok := r.Children[key]
	if ok {
		return c
	}
	if r.Key == "$$" {
		return r
	}
	c, ok = r.Children["$$"]
	if ok {
		return c
	}
	c, ok = r.Children["$"]
	return c
}
