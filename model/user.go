package model

type (
	User struct {
		ID       int    `json:"id" gorm:"primary_key"`
		Name     string `json:"name"`
		Email    string `json:"email" gorm:"uniqueIndex,length:30"`
		Password string `json:"password" column:"password"`
		Age      int    `json:"age"`
	}

	LoginUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password" column:"password"`
		Age      int    `json:"age"`
	}

	UpdateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password" column:"password"`
		Age      int    `json:"age"`
	}

	UserResponse struct {
		ID       int    `json:"id" gorm:"primary_key"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"-" column:"password"`
		Age      int    `json:"age"`
	}

	ErrorResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
)
