package dortsolver

type Solver struct {
	ApiKey    string `json:"api_key"`  // Your API Key
	PublicKey string `json:"site_key"` // The FunCraptcha public key
	ApiUrl    string `json:"api_url"`
	SiteUrl   string `json:"site_url"`
	Proxy     string `json:"proxy,omitempty"`
}

type Response struct {
	GameToken   string `json:"game[token]"`
	GameWaves   int64  `json:"game[waves]"`
	GameVariant string `json:"game[variant]"`
}
