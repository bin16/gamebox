package eventutil

type Channal chan interface{}

var channalMap map[string]Channal

func init() {
	channalMap = make(map[string]Channal)
}

func Post(key string, data interface{}) {
	if channalMap[key] == nil {
		channalMap[key] = make(Channal)
	}
	m := channalMap[key]

	m <- data
}

func Subscribe(key string) Channal {
	if channalMap[key] == nil {
		channalMap[key] = make(Channal)
	}

	return channalMap[key]
}
