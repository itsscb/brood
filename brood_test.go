package brood_test

import (
	"reflect"
	"testing"

	"github.com/itsscb/brood"
	"github.com/itsscb/brood/spore"
)

var testBrood = brood.New(spore.Channel(0), spore.Channel(1))

func TestPublish(t *testing.T) {
	hive := "test-hive"
	ch, _ := testBrood.Subscribe(hive)
	data := []byte("test-data")
	want := spore.New(hive, data)

	err := testBrood.Publish(want)
	if err != nil {
		t.Errorf("brood.Publish: want=<nil>, got=%#v", err)
	}

	got := <-ch
	if !reflect.DeepEqual(got, want) {
		t.Errorf("brood.Publish: want=%#v, got=%#v", got, want)
	}

}
