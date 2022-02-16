package mapdb

type MapDB struct {
	Data map[string]interface{}
}

func New() *MapDB {
	return &MapDB{Data: make(map[string]interface{})}
}

