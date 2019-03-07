package main

type Material struct {
	ID           string `json:"id"`
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	Count        int64  `json:"count"`
	provider     string `json:"provider"`
	providerLink string `json:"provider_link"`
	images       string `json:"images"`
}

type MaterialResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	Count        int64  `json:"count"`
	provider     string `json:"provider"`
	providerLink string `json:"provider_link"`
	images       string `json:"images"`
	User         struct {
		UserId   string `json:"user_id"`
		UserName string `json:"user_name"`
	}
}
