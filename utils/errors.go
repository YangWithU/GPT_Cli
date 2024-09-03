package GPT_cli_errors

import "errors"

var (
	ErrAPIKeyRequired = errors.New("API Key is required")

	// ErrInvalidModel 输入的Model有误
	ErrInvalidModel = errors.New("invalid model")

	// ErrNoMessages 输入Message为空
	ErrNoMessages = errors.New("no messages provided")

	// ErrInvalidRole Role错误，仅支持user,system,assistant
	ErrInvalidRole = errors.New("invalid role. Only `user`, `system` and `assistant` are supported")

	// ErrInvalidTemperature Temperature错误，temp应在区间 [0, 2]
	ErrInvalidTemperature = errors.New("invalid temperature. 0 <= temp <= 2")

	// ErrInvalidPresencePenalty presence penalty错误，应在区间 [-2, 2]
	ErrInvalidPresencePenalty = errors.New("invalid presence penalty. -2<= presence penalty <= 2")

	// ErrInvalidFrequencyPenalty frequency penalty错误，应在区间 [-2, 2]
	ErrInvalidFrequencyPenalty = errors.New("invalid frequency penalty. -2<= frequency penalty <= 2")
)
