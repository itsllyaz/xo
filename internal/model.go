package internal

import (

	"net/http"
	"os"
)
var (
	BearerToken = os.Getenv("BEARER_TOKEN")
	BaseURL = "https://api.x.com/2"
	
)

func BearerOAUTH(httpRequest *http.Request) {
	httpRequest.Header.Add("Authorization", "Bearer "+BearerToken)
	httpRequest.Header.Add("User-Agent", "v2UserTweetsGo")

}




