package exactonline_xml

import (
	"net/http"
	"net/url"
)

func (c *Client) NewMatchSetsDownloadRequest() MatchSetsDownloadRequest {
	return MatchSetsDownloadRequest{
		client:      c,
		queryParams: c.NewMatchSetsDownloadQueryParams(),
		pathParams:  c.NewMatchSetsDownloadPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewMatchSetsDownloadRequestBody(),
	}
}

type MatchSetsDownloadRequest struct {
	client      *Client
	queryParams *MatchSetsDownloadQueryParams
	pathParams  *MatchSetsDownloadPathParams
	method      string
	headers     http.Header
	requestBody MatchSetsDownloadRequestBody
}

func (c *Client) NewMatchSetsDownloadQueryParams() *MatchSetsDownloadQueryParams {
	return &MatchSetsDownloadQueryParams{}
}

type MatchSetsDownloadQueryParams struct {
	DefaultQueryParams

	// Grootboekrekening
	GLAccount string `schema:"Params_GLAccount,omitempty"`
	// Relatiecode
	AccountCode string `schema:"Params_AccountCode,omitempty"`
	// Data subscription token
	DownloadID string `schema:"Params_DownloadID,omitempty"`
}

func (p MatchSetsDownloadQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *MatchSetsDownloadRequest) QueryParams() *MatchSetsDownloadQueryParams {
	return r.queryParams
}

func (c *Client) NewMatchSetsDownloadPathParams() *MatchSetsDownloadPathParams {
	return &MatchSetsDownloadPathParams{}
}

type MatchSetsDownloadPathParams struct {
}

func (p *MatchSetsDownloadPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MatchSetsDownloadRequest) PathParams() *MatchSetsDownloadPathParams {
	return r.pathParams
}

func (r *MatchSetsDownloadRequest) SetMethod(method string) {
	r.method = method
}

func (r *MatchSetsDownloadRequest) Method() string {
	return r.method
}

func (s *Client) NewMatchSetsDownloadRequestBody() MatchSetsDownloadRequestBody {
	return MatchSetsDownloadRequestBody{}
}

type MatchSetsDownloadRequestBody struct {
}

func (r *MatchSetsDownloadRequest) RequestBody() *MatchSetsDownloadRequestBody {
	return &r.requestBody
}

func (r *MatchSetsDownloadRequest) SetRequestBody(body MatchSetsDownloadRequestBody) {
	r.requestBody = body
}

func (r *MatchSetsDownloadRequest) NewResponseBody() *MatchSetsDownloadResponseBody {
	return &MatchSetsDownloadResponseBody{}
}

type MatchSetsDownloadResponseBody struct {
	MatchSets MatchSets
	Topics    Topics
	Messages  Messages
}

func (r *MatchSetsDownloadRequest) URL() url.URL {
	return r.client.GetEndpointURL("XMLDownload.aspx?Topic=MatchSets", r.PathParams())
}

func (r *MatchSetsDownloadRequest) Do() (MatchSetsDownloadResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *MatchSetsDownloadRequest) All() ([]MatchSetsDownloadResponseBody, error) {
	rr := []MatchSetsDownloadResponseBody{}

	page := ""
	for {
		r.QueryParams().TSPaging = page
		resp, err := r.Do()
		if err != nil {
			return rr, err
		}

		rr = append(rr, resp)

		if len(resp.Topics) == 0 {
			break
		}

		if resp.Topics[0].Count >= resp.Topics[0].PageSize {
			break
		}

		page = resp.Topics[0].TSD
	}

	return rr, nil
}
