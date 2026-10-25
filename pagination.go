package nextmail

import (
	"context"
	"iter"
)

type pageFetcher func(ctx context.Context, cursor string) (items []Object, hasMore bool, nextCursor string, err error)

type Iterator struct {
	ctx      context.Context
	fetch    pageFetcher
	cursor   string
	followed map[string]struct{}
	items    []Object
	index    int
	current  Object
	started  bool
	more     bool
	next     string
	done     bool
	err      error
}

func newIterator(ctx context.Context, cursor string, fetch pageFetcher) *Iterator {
	followed := map[string]struct{}{}

	if cursor != "" {
		followed[cursor] = struct{}{}
	}

	return &Iterator{ctx: ctx, fetch: fetch, cursor: cursor, followed: followed}
}

func (i *Iterator) Next() bool {
	if i == nil || i.done {
		return false
	}

	for i.index >= len(i.items) {
		if !i.advance() {
			i.current = nil
			i.done = true

			return false
		}
	}

	i.current = i.items[i.index]
	i.index++

	return true
}

func (i *Iterator) advance() bool {
	if i.started {
		if _, seen := i.followed[i.next]; !i.more || i.next == "" || seen {
			return false
		}

		i.followed[i.next] = struct{}{}
		i.cursor = i.next
	}

	items, more, next, err := i.fetch(i.ctx, i.cursor)

	if err != nil {
		i.err = err

		return false
	}

	i.started = true
	i.items = items
	i.index = 0
	i.more = more
	i.next = next

	return true
}

func (i *Iterator) Item() Object {
	if i == nil {
		return nil
	}

	return i.current
}

func (i *Iterator) Err() error {
	if i == nil {
		return nil
	}

	return i.err
}

func (i *Iterator) All() iter.Seq2[Object, error] {
	return func(yield func(Object, error) bool) {
		for i.Next() {
			if !yield(i.Item(), nil) {
				return
			}
		}

		if err := i.Err(); err != nil {
			yield(nil, err)
		}
	}
}

func (i *Iterator) collect() ([]Object, error) {
	all := []Object{}

	for i.Next() {
		all = append(all, i.Item())
	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return all, nil
}
