package storeutil

type simpleMem struct {
	db map[string][]byte
}

func (m *simpleMem) open() {}

func (m simpleMem) get(key []byte) ([]byte, error) {
	sk := string(key)
	data := m.db
	return data[sk], nil
}

func (m simpleMem) set(key, value []byte) error {
	sk := string(key)
	data := m.db
	data[sk] = value
	return nil
}

func (m simpleMem) has(key []byte) (bool, error) {
	sk := string(key)
	data := m.db
	_, ok := data[sk]
	return ok, nil
}
