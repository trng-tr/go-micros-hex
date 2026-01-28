package usecase

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

func checkInputFields(fileds map[string]string) error {
	for key, value := range fileds {
		value = strings.TrimSpace(value)
		if value == "" {
			return fmt.Errorf("%w %s", errEmptyFields, key)
		} else if len(value) < 2 {
			return fmt.Errorf("%w %s", errTooShort, key)
		} else if len(value) > 255 {
			return fmt.Errorf("%w %s", errTooLong, key)
		}

	}

	return nil
}

func checkEmailValid(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	} else if len(email) < 5 {
		return false
	}
	return true
}

func checkInputGenda(genda domain.Genda) error {
	switch genda {
	case domain.Female, domain.Male:
		return nil
	default:
		{
			return fmt.Errorf("%w %s", errInvalidGenda, genda)
		}
	}
}

func checkInputId(id int64) error {
	if id > 0 {
		return nil
	}
	return fmt.Errorf("%w %d", errInvalidId, id)
}

func checkPhoneValid(phone string) error {
	var regex = regexp.MustCompile(`^\+?[0-9]{8,20}$`)
	if regex.MatchString(phone) {
		return nil
	}

	return fmt.Errorf("%w %s", errInvalidPhone, phone)
}

func generateDate() time.Time {
	return time.Now()
}

func checkZipCode(zip string) error {
	if len(zip) < 4 || len(zip) > 10 {
		return fmt.Errorf("%w", errInvalidZipCode)
	}

	return nil
}
