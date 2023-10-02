package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/google/uuid"
)

func (r ErrorResponse) ErrorMessage(e ErrorBody, code string, message interface{}, source string, details string) ErrorResponse {
	r.ErrorReference = uuid.New()
	r.TimeStamp = time.Now().Format(time.RFC3339)
	e.Code = code
	e.Message = message
	e.Source = source
	e.Details = details
	r.Errors = append(r.Errors, e)
	return r
}
   func (r SuccessBody) SuccessMessage(code, message string, reference string) SuccessBody {
	r.Code = code
	r.Message = message
	r.Reference = reference
	return r
}
func (r AllSuccessBody) AllSuccessMessage(code, message string, service interface{}) AllSuccessBody {
	r.Code = code
	r.Message = message
	r.Service = service
	return r
}
func validate() (*validator.Validate, ut.Translator) {
	validate := validator.New()
	readableErr := en.New()
	uni := ut.New(readableErr, readableErr)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}
	return validate, trans
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

func ValidateStruct(data interface{}) ([]error, error) {
	validate, trans := validate()
	err := validate.Struct(data)
	errorValue := translateError(err, trans)

	return errorValue, err
}
