package usecase

import (
	"fmt"
)

func checkValue(values map[string]int64) error {
	for key, value := range values {
		if value <= 0 {
			return fmt.Errorf("%w:%v", errInvalidVaue, key)
		}
	}

	return nil
}

/*func checkId(id int64) error {
	if id <= 0 {
		return errors.New("invalid id value !!!!!!!!!!!!!!!!")
	}
	return nil
}*/
