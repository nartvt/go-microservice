package entity

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"promotion-service/app/domain/model"
	"promotion-service/app/domain/usercases/promotion/repository"
	"promotion-service/app/proto-gen/message"
	"promotion-service/app/proto-gen/rpc"
	"promotion-service/app/uerror"
)

type prodmotionEntity struct {
	rpc.UnimplementedPromotionServiceServer
}

var PromotionEntity rpc.PromotionServiceServer

func init() {
	PromotionEntity = &prodmotionEntity{
		UnimplementedPromotionServiceServer: rpc.UnimplementedPromotionServiceServer{},
	}
}

func (p prodmotionEntity) GetPromotionById(ctx context.Context, request *message.PromotionRequest) (*message.PromotionResponse, error) {
	if request == nil || request.Id <= 0 {
		return &message.PromotionResponse{}, nil
	}
	promotion, err := repository.PromotionRepository.GetPromotionById(int(request.Id), true)
	if err != nil && err == gorm.ErrRecordNotFound {
		return &message.PromotionResponse{}, nil
	}
	if err != nil {
		return &message.PromotionResponse{}, uerror.InternalError(err, err.Error())
	}
	if promotion == nil {
		return &message.PromotionResponse{}, uerror.BadRequestError(fmt.Errorf("promotion not found %d", request.Id), "promotion not found")
	}
	return p.bind(promotion), nil
}

func (prodmotionEntity) bind(product *model.Promotion) *message.PromotionResponse {
	if product == nil {
		return &message.PromotionResponse{}
	}
	return &message.PromotionResponse{
		Id:    int64(product.Id),
		Code
		Active: &message.Active{
			Active: product.Active,
		},
		Price:     product.Price,
		CreatedAt: product.CreatedAt.UnixMilli(),
		UpdatedAt: product.UpdatedAt.UnixMilli(),
	}
}
