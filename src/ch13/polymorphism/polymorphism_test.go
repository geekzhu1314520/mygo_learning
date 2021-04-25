package polymorphism

import (
	"fmt"
	"testing"
)

type Code string
type Programer interface {
	WriteHelloWorld() Code
}

type GoProgramer struct {
}

func (p *GoProgramer) WriteHelloWorld() Code {
	return "fmt.Println(\"HelloWorld\")"
}

type JavaProgramer struct {
}

func (p *JavaProgramer) WriteHelloWorld() Code {
	return "System.out.println(\"HelloWorld\")"
}

func WriteProgram(p Programer) {
	fmt.Printf("%T, %s\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	// goProm := new(GoProgramer)
	goProm := &GoProgramer{}
	javaProm := new(JavaProgramer)
	WriteProgram(goProm)
	WriteProgram(javaProm)
}
