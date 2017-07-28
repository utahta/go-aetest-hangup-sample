package foo

import (
	"os"
	"testing"

	"google.golang.org/appengine/aetest"
)

var inst aetest.Instance

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func run(m *testing.M) int {
	var err error
	inst, err = aetest.NewInstance(nil)
	if err != nil {
		return 1
	}
	defer inst.Close()

	return m.Run()
}

func Test_Foo(t *testing.T) {
	t.Log("test: foo")

	panic("fail")
}
