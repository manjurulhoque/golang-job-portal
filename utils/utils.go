package utils

import (
	"encoding/json"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"io/ioutil"
	"net/http"
	"strings"
)

type IError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var (
	uni *ut.UniversalTranslator
	vl  *validator.Validate
)


func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// Translate errors
func TranslateError(s interface{}) (errs []IError) {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	vl = validator.New()
	_ = enTranslations.RegisterDefaultTranslations(vl, trans)

	_ = vl.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	err := vl.Struct(s)

	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)

	for _, e := range validatorErrs {
		//translatedErr := fmt.Errorf(e.Translate(trans))
		translatedErr := IError{
			Field:   strings.ToLower(e.Field()),
			Message: e.Translate(trans),
		}
		errs = append(errs, translatedErr)
	}
	return errs
}
