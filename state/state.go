package state

import "sync"

type State struct {
	mu        sync.Mutex `json:"-"`
	currentTS int64
	Datasets  []*Series
	View      []*Series
}
