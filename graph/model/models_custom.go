package model

type EditUser struct {
	ID          string  `json:"id" bson:"-"`
	Name        *string `json:"name,omitempty" bson:"name,omitempty"`
	Email       *string `json:"email,omitempty" bson:"email,omitempty"`
	Phone       *string `json:"phone,omitempty" bson:"phone,omitempty"`
	AccessToken *string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
}
