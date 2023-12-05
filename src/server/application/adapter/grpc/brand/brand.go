package brand

import context "context"

type Server struct {
	findBrandsUseCase usecase.FindBrandsUseCase
}

func NewServer(findBrandsUseCase usecase.FindBrandsUseCase) *Server {
	return &Server{
		findBrandsUseCase: findBrandsUseCase,
	}
}

func (s *Server) GetOrders(context context.Context, getOrdersRequest *GetOrdersRequest) (*GetOrdersResponse, error) {

	orders, err := s.findBrandsUseCase.Execute(context)
	if err != nil {
		return nil, err
	}

	orderResponse := make([]*Order, 0)

	for _, o := range orders {
		orderResponse = append(orderResponse, &Order{
			Id:           o.ID,
			OrderType:    o.OrderType,
			Store:        o.Store,
			Address:      o.Address,
			CreationDate: o.CreationDate,
			Status:       o.Status,
		})
	}

	return &GetOrdersResponse{
		Orders: orderResponse,
	}, nil
}
