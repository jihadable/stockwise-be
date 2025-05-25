package utils

import "github.com/go-playground/validator/v10"

func ParseValidationErrors(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		msg := ""
		for _, e := range errs {
			msg += e.Field() + " " + ValidationMessage(e) + ", "
		}
		return msg[:len(msg)-2]
	}
	return "Permintaan tidak valid"
}

func ValidationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "wajib diisi"
	case "gt":
		return "harus lebih besar dari " + e.Param()
	case "min":
		return "minimal bernilai " + e.Param()
	default:
		return "tidak valid"
	}
}
