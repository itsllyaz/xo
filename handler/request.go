package httphandler

import (
	"fmt"
	"log"
	"net/http"
	"xo/internal"
)

func CreateGetRequest(url string) (*http.Request, error) {
	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//log.Fatal(err)
		return httpRequest, err
	}
	setupHeader(httpRequest)
	return httpRequest, nil
}

func setupHeader(httpReqeust *http.Request) {
	httpReqeust.Header.Set("Authorization", "Bearer "+internal.BearerToken)

}

func GetResponse(httpRequest *http.Request) (*http.Response, error) {
	client := http.Client{}
	res, err := client.Do(httpRequest)
	if err != nil {
		return res, err
	}
	return res, nil
}

func UsersLookup(usernames string, user_fields string) (*http.Request, error) {
	//usernames := "elonmusk,llyas__"
	//baseURL := "https://api.x.com/2/users/by"
//	subURL := "/users/by/username/"
	// user_fields := "user.fields=description,created_at" // optional

	// params := url.Values{}
	// params.Add("usernames", usernames)
	// params.Add("user.Fields", user_fields)

//	fullURL := fmt.Sprintf("%s?%s", internal.BaseURL, params.Encode())
		
	// data(username)
	// fields(user_fields)
	// subURL(/users/by/username/{%s}
	// parmas(&parmas)
//	httpRequest, err := CreateGetRequest(fullURL)
//	if err != nil {
//		return httpRequest, err
//	}

	return nil, nil

}

func GetTweetsById(ids string, tweet_fields string) *http.Response {
	subURL := "/tweets"
	// data(username)
	// fields(tweet_field)
	// subURL
	//
	
	url := fmt.Sprintf("%s/%s", internal.BaseURL, subURL)
	// bearer_token := config.ReturnEnv()
	req, err := CreateGetRequest(url)
	if err != nil {
		log.Fatal(err)
	}

	res, err := GetResponse(req)
	if err != nil {
		log.Fatal(err)
	}

	return res

}
