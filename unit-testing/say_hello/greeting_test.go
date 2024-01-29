package say_hello

import (
	"fmt"
	"testing"
)

func Test_SayGoodBye(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Bye Bye %s", name)
	result := sayGoodBye(name)

	if (result != expected) {
		t.Errorf("\"sayGoodBye('%s')\" FAILED, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayGoodBye('%s')\" SUCCEED, expected -> %v, got -> %v", name, expected, result)
	}
}