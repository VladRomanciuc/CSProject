package structs

// GoogleResponse is the response sent by google
type CallbackResponse struct {
	ID       	string `json:"id"`
	Name 		string `json:"name"`
	GivenName	string `json:"given_name" ,json:"first_name"`
	FamilyName	string `json:"family_name" ,json:"last_name"`
	Gender		string `json:"gender"`
	Email    	string `json:"email"`
	Picture  	string `json:"picture"`
}