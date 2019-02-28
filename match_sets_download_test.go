package exactonline_xml_test

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	exactonline "github.com/omniboost/go-exactonline-xml"
	"golang.org/x/oauth2"
)

func getClient() (*exactonline.Client, error) {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	refreshToken := os.Getenv("OAUTH_REFRESH_TOKEN")
	tokenURL := os.Getenv("EO_TOKEN_URL")
	divisionID, err := strconv.Atoi(os.Getenv("EO_DIVISION_ID"))
	if err != nil {
		return nil, err
	}

	oauthConfig := exactonline.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set oauth base url
	// if p.config.BaseURL.String() != "" {
	// 	oauthConfig.SetBaseURL(&p.config.BaseURL.URL)
	// }

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client := exactonline.NewClient(httpClient, divisionID)
	client.SetDebug(false)
	return client, nil
}

func TestCustomersPostTest(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Error(err)
	}

	req := client.NewMatchSetsDownloadRequest()
	// req.QueryParams().TSPaging = ""
	// req.SetRequestBody(yoobi.CustomersPostRequestBody{
	// 	Name: "sagrosync test",
	// 	Code: "12",
	// })

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
