package configs

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

type OauthConfig struct {
	*oauth2.Config
	oauthStateStringGl oauth2.AuthCodeOption
}

func NewOauthConfGl() *OauthConfig {
	return &OauthConfig{
		&oauth2.Config{
			ClientID:     viper.GetString("google.clientID"),
			ClientSecret: viper.GetString("google.clientSecret"),
			RedirectURL:  "http://localhost:8080/callback-gl",
			Scopes:       []string{calendar.CalendarScope},
			Endpoint:     google.Endpoint,
		},
		oauth2.AccessTypeOffline,
	}
}

//func InitializeOAuthGoogle(oauthConfGl *OauthConfig) {
//	oauthConfGl.ClientID = viper.GetString("google.clientID")
//	oauthConfGl.ClientSecret = viper.GetString("google.clientSecret")
//	oauthConfGl.oauthStateStringGl = oauth2.AccessTypeOffline(viper.GetString("oauthStateString"))
//}
