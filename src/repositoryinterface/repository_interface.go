package repositoryinterface

/***This interface is common for all the repository files****/
type WeatherRepository interface {
	Insert(interface{}) error
	Get(interface{}) error
	Delete(id uint) error
}
