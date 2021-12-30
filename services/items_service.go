package services

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type itemsService struct{}

type ItemsServiceInterface interface {
	GetItem()
	SaveItem()
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
