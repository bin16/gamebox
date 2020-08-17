package storeutil

import (
	"testing"

	"github.com/bin16/Reversi/reversi"
)

func TestSet(t *testing.T) {
	g := reversi.NewGame()
	id := g.ID
	st := g.Status

	if Has(id) {
		t.Errorf("storeutil.Has failed, should be false")
	}

	Set(id, g.Save())
	if !Has(id) {
		t.Errorf("storeutil.Has failed, should be true")
	}

	g1data := Get(id)
	g1, _ := reversi.Load(g1data)
	if g1.ID != id || g1.Status != st {
		t.Errorf("storeutil.Get failed")
	}
}
