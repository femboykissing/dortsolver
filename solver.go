package dortsolver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func (solver *Solver) Solve() (*Response, error) {
	var response *Response

	// Check if there is an api key specified
	if len(solver.ApiKey) <= 0 {
		return nil, ErrNoApiKey
	}

	// Check if there is an api url specified if not just default it to client-api.arkoselabs.com
	if len(solver.ApiUrl) <= 0 {
		solver.ApiUrl = "https://client-api.arkoselabs.com"
	}

	// Check if there is an site url specified
	if len(solver.SiteUrl) <= 0 {
		return nil, ErrNoSiteUrl
	}

	// Check if there is an public key specified
	if len(solver.PublicKey) <= 0 {
		return nil, ErrNoPublicKey
	}

	// Marshal the json
	jsonData, err := json.Marshal(solver)
	if err != nil {
		return nil, err
	}

	// Send a request to the solver
	resp, err := http.Post(
		"https://api.dort.shop/captcha/solve/fc",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Read body
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check if captcha is unsolvable
	if strings.Contains(string(respBytes), "error") {
		return nil, ErrUnsolvableCaptcha
	}

	// Parse json body with the final response
	err = json.NewDecoder(bytes.NewReader(respBytes)).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
