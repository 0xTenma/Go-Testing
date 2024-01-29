package math

import (
	"testing"
)

func TestAbs(t *testing.T) {
	// t.Log("Similar to fmt.Println()")
	// t.Fail() // Will show a test cas has failed in the results
	// t.FailNow() // t.Fail() + safely exit without continuing
	// t.Error() // t.Log() + t.Fail();
	// t.Fatal() // t.Log() + t.FailNow();

	if Abs(-1) < 0 {
		t.Error("Negative value found in abs() with", -1);
	}
	if Abs(0) < 0 {
		t.Error("Negative value found in abs() with", 0);
	}
	if Abs(1) < 0 {
		t.Error("Negative value found in abs() with", 1);
	}
}