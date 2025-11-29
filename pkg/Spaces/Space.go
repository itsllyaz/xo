package space

import (
	"fmt"
	"net/http"
	"xo/internal"
)
var query_parmas = []string{
		"ids",
		"SPACE_ID",
		"space.fields",
		"title",
		"created_at",
		"expansions",
		"creator_id",
	}
func SearchSpace(search_term string) (*http.Response, error) {
	
	// query_parmas := map[string]any{
	// 	"query":        search_term,
	// 	"space.fields": []string{"title, created_at"},
	// 	"expansions":   "creator_id",
	// }
	// optional add params...
	// params := url.Values{}
	// params.Add("user.Fields", query_params[])
	// url := search_url + params.Encode()

	// sending the request to get spaces
	url := fmt.Sprintf("%v/space/search", internal.BaseURL)
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



func SpaceLookup(user_fields []string) (*http.Response, error) {
	// optional add params...
	// params := url.Values{}
	// params.Add("user.Fields", user_fields[0])
	
	url := fmt.Sprintf("%s/spaces", internal.BaseURL)
	Req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	internal.BearerOAUTH(Req)
	client := http.Client{}
	res, err := client.Do(Req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
