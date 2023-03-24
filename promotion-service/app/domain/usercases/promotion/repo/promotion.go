package repo

import (
	"context"
	"promotion-service/app/transport/proto-gen/message"
	"promotion-service/app/transport/proto-gen/rpc"
)

type promotionRepo struct {
	rpc.UnimplementedPromotionServiceServer
}

var PromotionRepo rpc.PromotionServiceServer

func init() {
	PromotionRepo = &promotionRepo{
		UnimplementedPromotionServiceServer: rpc.UnimplementedPromotionServiceServer{},
	}
}

func (p promotionRepo) GetPromotionById(ctx context.Context, request *message.PromotionRequest) (*message.PromotionResponse, error) {
	return nil, nil
}
func (p promotionRepo) GetPromotions(ctx context.Context, request *message.PromotionRequest) (*message.PromotionResponses, error) {
	return nil, nil
}

func (p promotionRepo) UpdatePromotion(crx context.Context, request *message.PromotionRequest) (*message.PromotionResponse, error) {
	return nil, nil
}

func (p promotionRepo) CreatePromotion(ctx context.Context, newProduct *message.PromotionRequest) (*message.PromotionResponse, error) {
	return nil, nil
}
