package projector_test

import (
	"testing"

	"github.com/RyanDev-21/pkg/projector"
)

func getData() *projector.Data {
	return &projector.Data{
		Projector: map[string]map[string]string{
			"/": {
				"foo": "bar",
				"fem": "boy",
			},
			"/foo": {
				"foo": "bar2",
			},
			"/foo/bar": {
				"foo": "bar3",
			},
		},
	}
}

func createProjector(pwd string, data *projector.Data) *projector.Projector {
	return projector.CreateProjector(&projector.Config{
		Args:      []string{},
		Pwd:       pwd,
		Config:    "Something",
		Operation: projector.Print,
	}, data)
}

func test(t *testing.T, projector *projector.Projector, key, value string) {
	v, ok := projector.GetValue(key)
	if !ok {
		t.Error("expected to get a value")
	}
	if value != v {
		t.Errorf("expected to get %v , got %v ", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	projector := createProjector("/foo/bar", data)
	test(t, projector, "foo", "bar3")
	projector = createProjector("/foo", data)
	test(t, projector, "foo", "bar2")
	projector = createProjector("/", data)
	test(t, projector, "foo", "bar")
}

func TestSetValue(t *testing.T) {
	data := getData()
	projector := createProjector("/foo/bar", data)
	projector.SetValue("foo", "barzzzz")
	test(t, projector, "foo", "barzzzz")
	projector.SetValue("foo", "barzzzz2")
	test(t, projector, "foo", "barzzzz2")
	projector = createProjector("/foo", data)
	projector.SetValue("foo", "barzzz2")
	test(t, projector, "foo", "barzzz2")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	projector := createProjector("/foo/bar", data)
	projector.RemoveValue("foo")
	test(t, projector, "foo", "bar2")
	projector = createProjector("/foo", data)
	projector.RemoveValue("foo")
	test(t, projector, "foo", "bar")
}
