package domain

import "context"

type Profile struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Votes       uint64 `json:"votes"`
	Students    uint64 `json:"students"`
	Website     string `json:"website"`
	Youtube     string `json:"youtube"`
	Linkedin    string `json:"linkedin"`
	Twitter     string `json:"twitter"`
	Facebook    string `json:"facebook"`
	Created     int64  `json:"created"`
	Modified    int64  `json:"modified"`
}
type Service interface {
	SaveProfile(ctx context.Context, profile *Profile) error
	UpdateProfile(ctx context.Context, profile *Profile, id uint64) error
	GetProfileById(ctx context.Context, id int64) (*Profile, error)
	DeleteProfileById(ctx context.Context, id uint64) error
}
