package repo

type DTO interface {
}

type Repository interface {
	// niech create zwraca id stworzonego elementu, a nie caly element, od tego masz read: atomowosc operacji ;)
	Create(model *DTO) (int64, error)
	Read(id int64) (*DTO, error)
	Update(id int64, model *DTO) error
	Delete(id int64) error
	List() ([]*DTO, error)
}
