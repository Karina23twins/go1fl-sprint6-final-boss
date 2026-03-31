package service

import (
	"errors"
	"log"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertString(s []byte) ([]byte, error) {
	// Преобразуем в строку
	str := string(s)
	if str == "" || len(strings.TrimSpace(str)) == 0 {
		log.Println("error: an empty string was passed")
		return nil, errors.New("error: an empty string was passed")
	}

	// Проверяем, Морзе или текст
	isMorse := true
	for _, r := range str {
		if r != '.' && r != '-' && r != ' ' {
			isMorse = false
			break
		}
	}

	var result string

	if isMorse {
		result = morse.ToText(str) // Морзе -> текст
	} else {
		result = morse.ToMorse(str) // Текст -> Морзе
	}

	return []byte(result), nil
}
