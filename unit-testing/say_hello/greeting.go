package say_hello

import "fmt"

func sayGoodBye(name string) string {
	if len(name) == 0 {
		return "Bye Bye Anon!"
	}

	return fmt.Sprintf("Bye Bye %s", name)
}

func sayHello(name string) string {
	if len(name) == 0 {
		return "Hello Anonymous"
	}

	return fmt.Sprintf("Hello %s", name)
}