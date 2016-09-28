package common

import (
	"errors"
	"reflect"
	"strconv"

	"gopkg.in/validator.v2"
)

func init() {
	validator.SetValidationFunc("strlen", strlen)
}

func strlen(v interface{}, param string) error {
	st := reflect.ValueOf(v)

	if st.Kind() != reflect.String {
		return errors.New("notZZ only validates strings")
	}

	length, err := strconv.Atoi(param)
	if err != nil {
		return err
	}

	val := st.String()

	if len(val) > length {
		return errors.New("value length fail.")
	}

	return nil
}

func ValidateStruct(obj interface{}) (bool, map[string]string) {
	keys := make(map[string]string)

	if errs := validator.Validate(obj); errs != nil {

		errar := errs.(validator.ErrorMap)

		for i, v := range errar {
			key := LowerFirst(i)
			var msg string
			for _, err := range v {
				msg = msg + err.Error()
			}
			keys[key] = msg
		}

		return false, keys
	}

	return true, keys
}
