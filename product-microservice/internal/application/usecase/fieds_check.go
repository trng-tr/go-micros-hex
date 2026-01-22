package usecase

import (
	"fmt"
	"strings"

	"github.com/trng-tr/product-microservice/internal/domain"
)

func checkInputs(fileds map[string]string) error {
	for key, value := range fileds {
		value = strings.TrimSpace(value)
		if value == "" {
			return fmt.Errorf("%w %s", errEmptyFields, key)
		} else if len(value) < 2 {
			return fmt.Errorf("%w, %s", errTooShort, key)
		} else if len(value) > 255 {
			return fmt.Errorf("%w %s", errTooLong, key)
		}
	}
	return nil
}

func checkInputId(id int64) error {
	if id > 0 {
		return nil
	}
	return fmt.Errorf("%w %d", errInvalidId, id)
}

func checkPrice(price domain.Price) error {
	if price.UnitPrice <= 0 {
		return fmt.Errorf("%w", errInvalidUnitPrice)
	}
	switch price.Currency {
	case domain.Dollar, domain.Euro:
		return nil
	default:
		return fmt.Errorf("%w %s", errInvalidCurrency, price.Currency)
	}
}

func checkProdCategory(cat domain.Category) error {
	switch cat {
	case domain.Book, domain.Clothing, domain.Shoes:
		return nil
	default:
		return fmt.Errorf("%w", errInvalidProductCat)
	}
}

func checkStockName(name string) error {
	if len(name) < 2 {
		return fmt.Errorf("%w", errTooShort)
	}
	return nil
}
func checkStockInputs(fields map[string]int64) error {
	for key, v := range fields {
		if v <= 0 {
			return fmt.Errorf("%w:%v", errInvalidStockField, key)
		}
	}
	return nil
}

func checkInputStockQty(q int64) error {
	if q < 0 {
		return fmt.Errorf("%w", errInvalidStockQty)
	}
	return nil
}
