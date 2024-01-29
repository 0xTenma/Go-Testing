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

// table driven testing
func Test_SayHello_ValidArgument(t *testing.T) {
	inputs := []struct {
		name string
		result string 
	}{
		{name: "Yemeksepeti", result: "Hello Yemeksepeti"},
		{name: "Banabi", result: "Hello Banabi"},
		{name: "Yemek", result: "Hello Yemek"},
	}

	for _, item := range inputs {
		result := sayHello(item.name)
		if (result != item.result) {
			t.Errorf("Failed")
		} else {
			t.Logf("Succeed")
		}
	}
}