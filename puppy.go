package puppy

import (
	"github.com/AmirH21/dog"
)

func Woof() string {
	return "Woof!"
}

func Woofs() string {
	return "Woof Woof Woof!"
}

func Dog() (string, string) {
	return dog.Dog(Woof())
}

func Dogs() (string, string) {
	return dog.Dog(Woofs())
}
