package handler

type OKResponse struct {
	Code int `json:"code"`
}

type RegResponse struct {
	Code   int `json:"code"`
	Status int `json:"status"`
}

type RegGetItem struct {
	AppKeyId     string `json:"app_key_id" binding:"required"`
	AppKeySecret string `json:"app_key_secret" binding:"required"`
	TemplateId   string `json:"template_id" binding:"required"`
	DomainList   string `json:"domain_list" binding:"required"`
}

type RegGetResponse struct {
	Code int           `json:"code"`
	List []*RegGetItem `json:"list"`
}
