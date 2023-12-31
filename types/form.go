package types

import (
	"github.com/google/uuid"
	"github.com/thuongtruong1009/gouse/constants"
	"github.com/thuongtruong1009/gouse/regex"
	"github.com/thuongtruong1009/gouse/shared"
	"github.com/thuongtruong1009/gouse/strings"
)

func IsUUID(input string) (bool, error) {
	if input == "" {
		return false, constants.ErrRequiredUUID
	}

	_, err := uuid.Parse(input)
	if err != nil {
		return false, constants.ErrInvalidUUID
	}

	return true, nil
}

func isMail(emailStr, domain string) (bool, error) {
	if strings.Includes(emailStr, "@") {
		if !regex.IsMatch(shared.EmailLenReg, emailStr) {
			return false, constants.ErrEmailLen
		}

		split := strings.Split(emailStr, "@")
		if len(split) != 2 || !strings.Includes(split[1], ".") {
			return false, constants.ErrInvalidEmail
		} else if !strings.Includes(split[1], domain) {
			return false, constants.ErrInvalidEmail
		} else {
			return true, nil
		}
	} else {
		return false, constants.ErrInvalidEmail
	}
}

func IsGmail(emailStr string) (bool, error) {
	return isMail(emailStr, "gmail.com")
}

func IsYahoo(emailStr string) (bool, error) {
	return isMail(emailStr, "yahoo.com")
}

func IsOutlook(emailStr string) (bool, error) {
	return isMail(emailStr, "outlook.com")
}

func IsEdu(emailStr string) (bool, error) {
	return isMail(emailStr, "edu")
}

func IsEmail(emailStr, customDomain string) (bool, error) {
	return isMail(emailStr, customDomain)
}

func IsUsername(username string) (bool, error) {
	if !regex.IsMatch(shared.UsernameReg, username) {
		return false, constants.ErrInvalidUsername
	}

	return true, nil
}

func IsPassword(password string) (bool, error) {
	if !regex.IsMatch(shared.PasswordLenReg, password) {
		return false, constants.ErrPasswordLen
	} else if !regex.IsMatch(shared.PasswordLowerReg, password) {
		return false, constants.ErrPasswordEmptyLower
	} else if !regex.IsMatch(shared.PasswordUpperReg, password) {
		return false, constants.ErrPasswordEmptyUpper
	} else if !regex.IsMatch(shared.PasswordDigitReg, password) {
		return false, constants.ErrPasswordEmptyDigit
	} else if !regex.IsMatch(shared.PasswordSpecialReg, password) {
		return false, constants.ErrPasswordEmptySpecial
	} else {
		return true, nil
	}
}

func IsPhone(phone string) (bool, error) {
	if !regex.IsMatch(shared.PhoneReg, phone) {
		return false, constants.ErrInvalidPhone
	}

	return true, nil
}
