package httpwrapper

import (
	"errors"
)

var (
	ErrorPayloadField     = errors.New("Payload missing the required key, or invalid values")
	ErrorQueryParams      = errors.New("Params missing the required query values, or invalid values")
	ErrorAuthHeader       = errors.New("Authorization Header missing token from Request")
	ErrorAuthToken        = errors.New("The token is invalid or expired")
	ErrorUserLogin        = errors.New("Invalid Email or Password")
	ErrorUserRegister     = errors.New("User already exists")
	ErrorUserNotFound     = errors.New("User not found")
	ErrorUserEmailInvalid = errors.New("Invalid email address or format")
	ErrorUserResetError   = errors.New("Email or password is incorrect")
	ErrorUserNotActive    = errors.New("User not activated")
	ErrorUserPermission   = errors.New("The user might not have the necessary permissions")
	ErrorFormActived      = errors.New("This form has been activated, you can't update it")
	ErrorResourceExist    = errors.New("The resource already exists")
	ErrorResourceNotFound = errors.New("The resource not found")
	ErrorResourceRefer    = errors.New("This resource has a reference object using it")
	ErrorDeposit          = errors.New("There was an error in your deposit. You have not enough points")
)
