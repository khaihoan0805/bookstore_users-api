package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct {
}

type itemsServiceInterface interface {
	GetItem()
	CreateItem()
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) CreateItem() {

}
