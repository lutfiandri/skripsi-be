// https://developers.home.google.com/cloud-to-cloud/project/authorization
package oauth

// /authorize
type OAuthAuthorizeRequest struct {
	ClientId     string `query:"client_id" validate:"required"`
	RedirectUri  string `query:"redirect_uri" validate:"required"`
	State        string `query:"state" validate:"required"`
	ResponseType string `query:"response_type"`
	// Scope string
	// UserLocale string
}

type OAuthAuthorizeResponse struct {
	RedirectUri string `json:"redirect_uri"`
}

type OAuthTokenRequest struct {
	ClientId     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
	GrantType    string `query:"grant_type"` // "authorization_code" || "refresh_token"
	Code         string `query:"code"`
	RefreshToken string `query:"refresh_token"`
	RedirectUri  string `query:"redirect_uri"`
}

// /token
type OAuthTokenResponse struct {
	TokenType    string `json:"token_type"` // "Bearer"
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"` // null when grant_type is "refresh_token"
	ExpiresIn    uint   `json:"expires_in"`    // in seconds
}

type OAuthTokenErrorResponse struct {
	Error string `json:"error"` // "invalid_grant"
}
