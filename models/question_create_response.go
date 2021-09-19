package models

type CreateQuestionResponse struct {
	NumCreated int         `json:"num_created"`
	NumOmitted int         `json:"num_omitted"`
	Created    interface{} `json:"created"`
}
