package models

import (
	"encoding/json"
	"time"
)

type (
	User struct {
		ID          int64           `json:"id"`
		First_name  string          `json:"first_name"`
		Middle_name string          `json:"middle_name"`
		Last_name   string          `json:"last_name"`
		Birthdate   time.Time       `json:"birthdate"`
		Address     json.RawMessage `json:"address"`
		Username    string          `json:"username"`
		Password    string          `json:"password"`
		Mpin        string          `json:"mpin"`

		// Address     map[string]interface{} `json:"address"`
	}
	Cred struct {
		Userid    int64     `json:"user_id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		Mpin      string    `json:"mpin"`
		Createdat time.Time `json:"created_at"`
		Updatedat time.Time `json:"updated_at"`
	}
	Credentials struct {
		User_id     int64           `json:"id"`
		Username    string          `json:"username"`
		First_name  string          `json:"first_name"`
		Middle_name string          `json:"middle_name"`
		Last_name   string          `json:"last_name"`
		Birthdate   time.Time       `json:"birthdate"`
		Address     json.RawMessage `json:"address"`
		Password    string          `json:"password"`
		Mpin        string          `json:"mpin"`
		Createdat   time.Time       `json:"created_at"`
		Updatedat   time.Time       `json:"updated_at"`
	}

	Login struct {
		User_id  int64  `json:"id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Mpin     string `json:"mpin"`
	}
	Register struct {
		ID         int64           `json:"id"`
		User_id    int64           `json:"user_id"`
		Username   string          `json:"username"`
		Password   string          `json:"password"`
		Createdat  time.Time       `json:"created_at"`
		Updatedat  time.Time       `json:"updated_at"`
		Firstname  string          `json:"first_name"`
		Middlename string          `json:"middle_name"`
		Lastname   string          `json:"last_name"`
		Birthdate  string          `json:"birthdate"`
		Address    json.RawMessage `json:"address"`
	}

	Updaterequest struct {
		Userid   int64  `json:"user_id"`
		Password string `json:"password"`
	}

	Updatempin struct {
		Userid string `json:"user_id"`
		Mpin   string `json:"mpin"`
	}
)
