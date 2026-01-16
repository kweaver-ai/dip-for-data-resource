package constant

// Deprecated: Use middleware/v1.NewContextWithUser to create a new context
// including given user. User middleware/v1.UserFromContext to get user from
// context.
const InfoName = "info"

// Deprecated: To distinguish HTTP Header Authorization from Bearer Token, use
// NewContextWithAuth and NewContextWithBearerToken to create context.Context,
// use SetGinContextWithAuth and SetGinContextWithBearerToken to set
// gin.Context, use AuthFromContext to get HTTP Header Authorization from
// context.Context, use BearerTokenFromContext to get Bearer Token from
// context.Context.
const Token = "token"
const TokenType = "token_type" //TokenTypeUser or TokenTypeClient

const StatusCode = "StatusCode"

const (
	TokenTypeUser = iota
	TokenTypeClient
)
