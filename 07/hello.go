// Package hello provides greeting messages.
package hello

import "fmt"

// ReturnGreeting returns different greeting message depending on the given greeting parameter.
func ReturnGreeting(greeting string) string {
	return fmt.Sprintf("%s yourself!", greeting)
}
