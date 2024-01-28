package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	d.p.Speak()
}

func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T) {
	var dog = new(Dog)
	dog.SpeakTo("Chao")
	fmt.Printf("%T", dog)
}
