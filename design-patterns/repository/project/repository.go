package project

// Repository is a Repository pattern interface
type Repository interface {
	FindAll() ([]Project, error)
	Store(Project) (Project, error)
	Delete(Project) error
}
