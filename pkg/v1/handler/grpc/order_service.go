package grpc

import (
	"context"
	"errors"
	"github.com/NickNaskida/Gorder/internal/models"
	"github.com/google/uuid"
	"strconv"

	interfaces "github.com/NickNaskida/Gorder/pkg/v1"
	pb "github.com/NickNaskida/Gorder/proto"
	"google.golang.org/grpc"
)

type OrderServStruct struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedOrderServiceServer
}

func (srv *OrderServStruct) transformCreateOrderRPC(req *pb.CreateOrderRequest) models.Order {
	orderId := uuid.New()
	return models.Order{OrderId: orderId.String(), Name: req.GetName(), Description: req.GetDescription(), Price: req.GetPrice()}
}

func (srv *OrderServStruct) transformUpdateOrderRPC(req *pb.UpdateOrderRequest) models.Order {
	return models.Order{OrderId: req.GetOrderId(), Name: req.GetName(), Description: req.GetDescription(), Price: req.GetPrice()}
}

func (srv *OrderServStruct) transformOrderModel(order models.Order) *pb.OrderResponse {
	return &pb.OrderResponse{Id: strconv.Itoa(int(order.ID)), OrderId: order.OrderId, Name: order.Name, Description: order.Description, Price: order.Price}
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &OrderServStruct{useCase: usecase}
	pb.RegisterOrderServiceServer(grpcServer, userGrpc)
}

// CreateOrder function creates an order with the supplied data from CreateOrderRequest message of proto
func (srv *OrderServStruct) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	data := srv.transformCreateOrderRPC(req)
	if data.Name == "" || data.Price == 0 {
		return &pb.OrderResponse{}, errors.New("name and price are required")
	} else if data.Price < 0 {
		return &pb.OrderResponse{}, errors.New("price cannot be negative")
	}

	order, err := srv.useCase.Create(data)
	if err != nil {
		return &pb.OrderResponse{}, err
	}

	return srv.transformOrderModel(order), nil
}

// GetOrder function retrieves an order from the database using the supplied id from GetOrderRequest message of proto
func (srv *OrderServStruct) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	orderId := req.GetOrderId()
	if orderId == "" {
		return &pb.OrderResponse{}, errors.New("order_id is required")
	}

	order, err := srv.useCase.Get(orderId)
	if err != nil {
		return &pb.OrderResponse{}, err
	}

	return srv.transformOrderModel(order), nil
}

// UpdateOrder function updates an order in the database using the supplied data from UpdateOrderRequest message of proto
func (srv *OrderServStruct) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	data := srv.transformUpdateOrderRPC(req)
	if data.OrderId == "" {
		return &pb.OrderResponse{}, errors.New("order_id is required")
	}

	order, err := srv.useCase.Update(data)
	if err != nil {
		return &pb.OrderResponse{}, err
	}

	return srv.transformOrderModel(order), nil
}

// DeleteOrder function deletes an order from the database using the supplied id from DeleteOrderRequest message of proto
func (srv *OrderServStruct) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.SuccessResponse, error) {
	orderId := req.GetOrderId()
	if orderId == "" {
		return &pb.SuccessResponse{}, errors.New("order_id is required")
	}

	err := srv.useCase.Delete(orderId)
	if err != nil {
		return &pb.SuccessResponse{}, err
	}

	return &pb.SuccessResponse{Message: "Deleted Successfully"}, nil
}
