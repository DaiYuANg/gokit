package eds

type Storage interface {
	Get(name string) ([]Record, error)
	Set(record Record) error
	Delete(name string) error
	List() ([]Record, error)
}
