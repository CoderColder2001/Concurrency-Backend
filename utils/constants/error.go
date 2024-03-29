package constants

import (
	"Concurrency-Backend/api"
	"errors"
)

var (
	InvalidTokenErr      = errors.New(api.ErrorCodeToMsg[api.TokenInvalidErr])
	NoVideoErr           = errors.New(api.ErrorCodeToMsg[api.NoVideoErr])
	UnKnownActionTypeErr = errors.New(api.ErrorCodeToMsg[api.UnKnownActionType])

	UserNotExistErr     = errors.New(api.ErrorCodeToMsg[api.UserNotExistErr])
	UserAlreadyExistErr = errors.New(api.ErrorCodeToMsg[api.UserAlreadyExistErr])
	RecordNotExistErr   = errors.New(api.ErrorCodeToMsg[api.RecordNotExistErr])
	RecordNotMatchErr   = errors.New(api.ErrorCodeToMsg[api.RecordNotMatchErr])
	InnerDataBaseErr    = errors.New(api.ErrorCodeToMsg[api.InnerDataBaseErr])
	CreateDataErr       = errors.New(api.ErrorCodeToMsg[api.CreateDataErr])

	VideoFormatErr = errors.New(api.ErrorCodeToMsg[api.VideoFormationErr])
	VideoSizeErr   = errors.New(api.ErrorCodeToMsg[api.VideoSizeErr])
	SavingFailErr  = errors.New(api.ErrorCodeToMsg[api.SavingFailErr])
	UploadFailErr  = errors.New(api.ErrorCodeToMsg[api.UploadFailErr])
)
