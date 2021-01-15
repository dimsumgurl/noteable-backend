package domain

type UserRepo interface {
	CreateCollection() error
	Create() error
	RetrieveByID() error
	RetrieveByEmail() error
	Update() error
	Delete() error
}
