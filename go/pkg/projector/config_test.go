package projector_test

import (
	"reflect"
	"testing"

	"github.com/RyanDev-21/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	return &projector.Opts{
		Args:   args,
		Pwd:    "",
		Config: "",
	}
}

func testConfig(t *testing.T, args []string, expected []string, op projector.Operation) {
	opts := getOpts(args)
	config, err := projector.NewConfig(opts)
	if err != nil {
		t.Errorf("expected to get no error: %v", err)
	}
	if !reflect.DeepEqual(config.Operation, op) {
		t.Errorf("expected operation to be %v :got %v", op, config.Operation)
	}
	// if !reflect.DeepEqual(config.Config, configPath) {
	// 	t.Errorf("expected operation to be %v :got %v", configPath, config.Config)
	// }
	// if !reflect.DeepEqual(config.Pwd, pwd) {
	// 	t.Errorf("expected operation to be %v :got %v", pwd, config.Pwd)
	// }
	switch op {
	case projector.Print:
		if !reflect.DeepEqual(config.Args, args) {
			t.Errorf("expected args to be %v :got %v", expected, config.Args)
		}
	default:
		if !reflect.DeepEqual(config.Args, args[1:]) {
			t.Errorf("expected args to be %v :got %v", expected, config.Args)
		}
	}
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKeyValue(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKey(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"rm", "foo"}, []string{"foo"}, projector.Remove)
}
