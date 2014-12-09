package board

import ()

type UserType struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	IsAdmin      bool   `json:"is_admin"`
	ProfileImage string `json:"profile_image"`
}
