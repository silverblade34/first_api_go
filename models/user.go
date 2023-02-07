package models
import "time"

type User struct{
	ID 			string		 ` json:"id" `
	Name		string		 ` json:"name" `
	Email		string		 ` json:"email" `
	CreatedAt	time.Time	 ` json:"created_at" `
	UpdateAt	time.Time	 ` json:"update_at,omitempty"`
}

type Users []User

