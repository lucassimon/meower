package db

type Repository interface {
	Close() error
	InsertMeow(meow Meow) error
}

var impl Repository = &MemoryRepository{}

func SetRepository(repository Repository) {
	impl = repository
}

func InsertMeow(meow Meow) error {
	return impl.InsertMeow(meow)
}