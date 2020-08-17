package eventutil

import (
	"log"
	"testing"
)

func next(ch Channal) {
	select {
	case m := <-ch:
		log.Println(m)
	}
}

func TestIf(t *testing.T) {
	ch := Subscribe("k1")
	go func() {
		Post("k1", "aaa")
		Post("k1", "bbb")
		Post("k1", "ccc")
		Post("k1", "ddd")
	}()
	next(ch)
}

func TestSingle(t *testing.T) {
	m := "Hello World!"
	ch := Subscribe("k1")
	go func() {
		Post("k1", m)
	}()
	m1 := <-ch
	if m1 != m {
		t.Errorf("Failed to get message; got %s; want %s", m1, m)
	}
}

func TestDouble(t *testing.T) {
	m := "Hello World!"
	n := "Biu Biu Biu!"
	ch := Subscribe("k1")
	go func() {
		Post("k1", m)
		Post("k1", n)
	}()
	m1 := <-ch
	n1 := <-ch
	if m1 != m {
		t.Errorf("Failed to get message; got %s; want %s", m1, m)
	}
	if n1 != n {
		t.Errorf("Failed to get nessage; got %s; want %s", n1, n)
	}
}

func TestAll(t *testing.T) {
	ml := []string{"Hello", "World"}
	ch := Subscribe("k1")
	i := 0

	go func() {
		for _, m := range ml {
			Post("k1", m)
		}
		close(ch)
	}()

	for m := range ch {
		if m != ml[i] {
			t.Errorf("Failed to get message %d; got %s; want %s", i, m, ml[i])
		}
		i++
	}
}
