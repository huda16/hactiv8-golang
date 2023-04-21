package helpers

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type LoginInput struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type CommentInput struct {
	Message string `json:"message" valid:"required"`
	PhotoID uint   `json:"photo_id" valid:"required"`
}

type PhotoInput struct {
	Title    string `json:"title" valid:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"required"`
}

type SocialMediaInput struct {
	Name           string `json:"name" valid:"required"`
	SocialMediaUrl string `json:"social_media_url" valid:"required"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type APIError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
