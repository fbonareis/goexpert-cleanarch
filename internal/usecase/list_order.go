package usecase

import "github.com/fbonareis/goexpert-cleanarch/internal/entity"

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderOutputDTO struct {
	Total int        `json:"total"`
	Data  []OrderDTO `json:"data"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}
	var dto ListOrderOutputDTO
	for _, order := range orders {
		dto.Data = append(dto.Data, OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	dto.Total = len(orders)
	return dto, nil
}
