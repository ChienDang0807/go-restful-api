package utils

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf(" %s is required", fieldName)
	}

	return nil
}

func ValidationStringLength(fieldName, value string, max, min int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be betwenn %d and %d characters", fieldName, min, max)
	}

	return nil
}

func ValidationRegex(fieldName, value string, re *regexp.Regexp, errorMessage string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s %s", fieldName, errorMessage)
	}

	return nil
}

func ValidationPositiveInt(fieldName, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number", fieldName)
	}

	if v <= 0 {
		return 0, fmt.Errorf("%s must be positive", fieldName)
	}

	return v, nil
}

func ValidationUuid(fieldName, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", fieldName)
	}

	return uid, nil
}

func ValidationInList(fieldName, value string, allowed map[string]bool) error {
	if !allowed[value] {
		return fmt.Errorf("%s must be one of:%v", fieldName, keys(allowed))
	}
	return nil
}

func keys(m map[string]bool) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}

	return k
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(``)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.FieldName())
	})

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		minVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}

		return false
	})
	return nil
}
