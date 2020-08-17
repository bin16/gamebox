package storeutil

var data map[string][]byte = make(map[string][]byte)

func Get(key string) []byte {
	return data[key]
}

func Set(key string, value []byte) {
	data[key] = value
}

func Has(key string) bool {
	return data[key] != nil
}

func Clear(key string) {
	data[key] = nil
}
