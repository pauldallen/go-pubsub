package datastore

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestLoadDatastore(t *testing.T) {
	cases := []struct {
		input  *Config
		expect string
	}{
		{
			&Config{},
			"*datastore.Memory",
		},
		{
			// WARNING: Require connect redis
			&Config{
				Redis: &RedisConfig{
					Addr: "localhost:6379",
					DB:   0,
				},
			},
			"*datastore.Redis",
		},
	}
	for i, c := range cases {
		d, err := LoadDatastore(c.input)
		if err != nil {
			t.Fatalf("#%d: failed to load datastore, got err %v", i, err)
		}
		if got := reflect.TypeOf(d); got.String() != c.expect {
			t.Errorf("#%d: datastore type want %s, got %s", i, c.expect, got)
		}
	}
}

func TestMemorySet(t *testing.T) {
	msgA := dummy{ID: "a"}
	msgB := dummy{ID: "b"}

	cases := []struct {
		inputMsgs []dummy
		expect    map[interface{}]interface{}
	}{
		{
			[]dummy{
				msgA, msgA, msgB,
			},
			map[interface{}]interface{}{
				"a": msgA, "b": msgB,
			},
		},
	}
	for i, c := range cases {
		m := NewMemory(nil)
		for _, v := range c.inputMsgs {
			m.Set(v.ID, v)
		}
		if got := m.Store; !reflect.DeepEqual(got, c.expect) {
			t.Errorf("#%d: want %v, got %v", i, c.expect, got)
		}
	}
}

func TestMemoryGet(t *testing.T) {
	msgA := dummy{ID: "a"}
	msgB := dummy{ID: "b"}
	baseStore := Memory{
		Store: map[interface{}]interface{}{"a": msgA, "b": msgB},
	}

	cases := []struct {
		input     string
		expectMsg interface{}
		expectErr error
	}{
		{
			"a",
			msgA,
			nil,
		},
		{
			"c",
			nil,
			ErrNotFoundEntry,
		},
	}
	for i, c := range cases {
		got, err := baseStore.Get(c.input)
		if errors.Cause(err) != c.expectErr {
			t.Fatalf("#%d: want %v, got %v", i, c.expectErr, err)
		}
		if !reflect.DeepEqual(got, c.expectMsg) {
			t.Errorf("#%d: want %v, got %v", i, c.expectMsg, got)
		}
	}
}

func TestMemoryDelete(t *testing.T) {
	msgA := dummy{ID: "a"}
	msgB := dummy{ID: "b"}
	baseStore := Memory{
		Store: map[interface{}]interface{}{"a": msgA, "b": msgB},
	}

	cases := []struct {
		input       string
		expectStore map[interface{}]interface{}
		expectErr   error
	}{
		{
			"a",
			map[interface{}]interface{}{"b": msgB},
			nil,
		},
		{
			// delete to non exist key
			"a",
			map[interface{}]interface{}{"b": msgB},
			nil,
		},
	}
	for i, c := range cases {
		err := baseStore.Delete(c.input)
		if errors.Cause(err) != c.expectErr {
			t.Errorf("#%d: want %v, got %v", i, c.expectErr, err)
		}
		if !reflect.DeepEqual(baseStore.Store, c.expectStore) {
			t.Errorf("#%d: want %v, got %v", i, c.expectStore, baseStore.Store)
		}
	}
}
