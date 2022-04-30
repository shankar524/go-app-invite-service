package lib

import (
	"github.com/shankar524/password-generator/src/text"
)

func NewTextGenerator() text.Generator {
	textBuilder := text.TextBuilder{}
	textBuilder.AddRule(text.TextRule{For: text.LOWERCASE, Length: 5})
	textBuilder.AddRule(text.TextRule{For: text.NUMBERS, Length: 1})
	textBuilder.AddRule(text.TextRule{For: text.UPPERCASE, Length: 3})

	generator, _ := textBuilder.Build()
	return generator
}
