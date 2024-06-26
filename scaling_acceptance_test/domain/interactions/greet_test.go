package interactions

import (
	"go-specs-greet/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(Greet))
}
