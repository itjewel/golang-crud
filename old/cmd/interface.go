type IUserRepository interface {
	Insert(u models.Users) (int64, error)
	GetAll() ([]models.Users, error)
	GetOneUser(id int) (*models.Users, error)
	Update(ctx context.Context, u models.Users) (int64, error)
	DeleteUser(u models.Users) (int64, error)
	TextSearch(u models.Users) ([]models.Users, error)
}

type IServiceInterface interface {
	Insert(u models.User) (int64, error)
	GetAll() []models.Users
	Update()
	Delete()
}