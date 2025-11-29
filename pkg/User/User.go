package User

import (
	"fmt"
	"net/http"
	"xo/internal"
)

func GetUser(userName string) (*http.Response, error) {
	
	url := fmt.Sprintf(internal.BaseURL+"/users/by/%v", userName)
	req, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		return nil, err
	}
	internal.BearerOAUTH(req)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UserMentions(tweetID any) (*http.Response, error) {
	//url := fmt.Sprintf("https://api.x.com/2/users/%v/mentions", tweetID)
	url := fmt.Sprintf(internal.BaseURL+"/users/%v/mentions", tweetID)
	req, err := http.NewRequest("GET", url, nil)
	
	if err != nil {
		return nil, err

	}
	internal.BearerOAUTH(req)
	
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil

}
