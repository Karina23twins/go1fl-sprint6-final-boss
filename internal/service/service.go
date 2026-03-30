package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertString(s string) (string, error) {
	if s == "" || len(strings.TrimSpace(s)) == 0 {
		log.Println("error: an empty string was passed")
		return "", errors.New("error: an empty string was passed")
	}

	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' {
			morseString := morse.ToMorse(s)
			return morseString, nil
		}
	}

	textString := morse.ToText(s)
	return textString, nil
}
