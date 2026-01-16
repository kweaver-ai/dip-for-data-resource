package errorcode

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agcodes"
	"github.com/kweaver-ai/idrm-go-frame/core/errorx/agerrors"
)

type ErrorCodeFullInfo struct {
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Cause       string      `json:"cause"`
	Solution    string      `json:"solution"`
	Detail      interface{} `json:"detail,omitempty"`
}

type errorCodeInfo struct {
	description string
	cause       string
	solution    string
}

func New(errorCode, description, cause, solution string, detail interface{}, errLink string) error {
	coder := agcodes.New(errorCode, description, cause, solution, detail, errLink)
	return agerrors.NewCode(coder)
}

type errorCode map[string]errorCodeInfo

var errorCodeMap errorCode

func Contains(err error, msg string) bool {
	errCoder, ok := err.(*agerrors.Error)
	if !ok {
		return strings.Contains(err.Error(), msg)
	}
	coder := errCoder.Code()
	return strings.Contains(fmt.Sprintf("%v", coder.GetErrorDetails()), msg) ||
		strings.Contains(coder.GetDescription(), msg) ||
		strings.Contains(coder.GetErrorCode(), msg)
}

func IsSameErrorCode(err error, codeString string) bool {
	code, ok := err.(*agerrors.Error)
	if !ok {
		return false
	}
	return code.Code().GetErrorCode() == codeString
}

func registerErrorCode(errCodes ...errorCode) {
	if errorCodeMap == nil {
		// errorCodeMap init
		errorCodeMap = errorCode{}
	}

	for _, m := range errCodes {
		for k := range m {
			if _, ok := errorCodeMap[k]; ok {
				// error code is not allowed to repeat
				panic(fmt.Sprintf("error code is not allowed to repeat, code: %s", k))
			}

			errorCodeMap[k] = m[k]
		}
	}
}

func init() {
	registerErrorCode(publicErrorMap)
	registerErrorCode(knowledgeNetworkErrorMap)
	registerErrorCode(drivenErrorMap)
}

func Desc(errCode string, args ...any) error {
	return newCoder(errCode, nil, args...)
}

func Detail(errCode string, err any, args ...any) error {
	return newCoder(errCode, err, args...)
}

func newCoder(errCode string, err any, args ...any) error {
	errInfo, ok := errorCodeMap[errCode]
	if !ok {
		errInfo = errorCodeMap[PublicInternalError]
		errCode = PublicInternalError
	}

	desc := errInfo.description
	if len(args) > 0 {
		desc = FormatDescription(desc, args...)
	}
	if err == nil {
		err = struct{}{}
	}

	coder := agcodes.New(errCode, desc, errInfo.cause, errInfo.solution, err, "")
	return agerrors.NewCode(coder)
}

// FormatDescription replace the placeholder in coder.Description
// Example:
// Description: call service [service_name] api [api_name] error,
// args:  af-sailor-service, create
// =>
// Description: call service [af-sailor-service] api [create] error,
func FormatDescription(s string, args ...interface{}) string {
	if len(args) <= 0 {
		return s
	}
	re, _ := regexp.Compile("\\[\\w+\\]")
	result := re.ReplaceAll([]byte(s), []byte("[%v]"))
	return fmt.Sprintf(string(result), args...)
}
