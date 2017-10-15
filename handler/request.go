package handler

type RegStartRequest struct {
	AppKeyId     string `json:"app_key_id" binding:"required"`
	AppKeySecret string `json:"app_key_secret" binding:"required"`
	TemplateId   string `json:"template_id" binding:"required"`
	DomainList   string `json:"domain_list" binding:"required"`
	ThreadCount  int    `json:"thread_count" binding:"required"`
}
