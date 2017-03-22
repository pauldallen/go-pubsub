package queue

import (
	"reflect"
	"testing"
)

type DummyDatastore struct{}

func setupGlobal() {
	GlobalTopics = newTopics()
}

func dummyTopic(name string) *Topic {
	return &Topic{
		name: name,
	}
}

func dummyTopics(args ...string) *topics {
	m := newTopics()
	for _, a := range args {
		m.Set(dummyTopic(a))
	}
	return m
}

func TestNewTopic(t *testing.T) {
	cases := []struct {
		inputs       []string
		expectErr    error
		expectTopics *topics
	}{
		{
			[]string{"a", "b"},
			nil,
			dummyTopics("a", "b"),
		},
		{
			[]string{"a", "a"},
			ErrAlreadyExistTopic,
			dummyTopics("a"),
		},
	}
	for i, c := range cases {
		setupGlobal()
		var err error
		for _, s := range c.inputs {
			// expect last input return value equal expectErr
			_, err = NewTopic(s, nil)
		}
		if err != c.expectErr {
			t.Errorf("%#d: want %v, got %v", i, c.expectErr, err)
		}
		if !reflect.DeepEqual(GlobalTopics, c.expectTopics) {
			t.Errorf("%#d: want %v, got %v", i, c.expectTopics, GlobalTopics)
		}
	}
}

func TestGetTopic(t *testing.T) {
	// make test topics
	setupGlobal()
	GlobalTopics.Set(dummyTopic("a"))
	GlobalTopics.Set(dummyTopic("b"))

	cases := []struct {
		input       string
		expectTopic *Topic
		expectErr   error
	}{
		{"a", dummyTopic("a"), nil},
		{"c", nil, ErrNotFoundTopic},
	}
	for i, c := range cases {
		got, err := GetTopic(c.input)
		if err != c.expectErr {
			t.Errorf("%#d: want %v, got %v", i, c.expectErr, err)
		}
		if !reflect.DeepEqual(got, c.expectTopic) {
			t.Errorf("%#d: want %v, got %v", i, c.expectTopic, got)
		}
	}
}

func TestDelete(t *testing.T) {
	cases := []struct {
		baseTopics *topics
		input      *Topic
		expect     *topics
	}{
		{
			dummyTopics("a", "b"),
			dummyTopic("a"),
			dummyTopics("b"),
		},
		{
			dummyTopics("a", "b"),
			dummyTopic("c"),
			dummyTopics("a", "b"),
		},
	}
	for i, c := range cases {
		GlobalTopics = c.baseTopics
		// delete depends topic.name
		c.input.Delete()
		if !reflect.DeepEqual(GlobalTopics, c.expect) {
			t.Errorf("%#d: want %v, got %v", i, c.expect, GlobalTopics)
		}
	}
}

// TODO: integration datastore and subscription
func TestPublish(t *testing.T) {
	cases := []struct {
		input     Message
		expectErr error
	}{
		{
			Message{},
			nil,
		},
	}
	for i, c := range cases {
		topic := dummyTopic("a")
		topic.store = newTestDatastore()
		topic.subscriptions = []Subscription{}
		got := topic.Publish(c.input)
		if got != c.expectErr {
			t.Errorf("%#d: want %v, got %v", i, c.expectErr, got)
		}
	}
}
