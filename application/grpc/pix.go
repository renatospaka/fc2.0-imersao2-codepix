package grpc

import (
	"context"

	"github.com/renatospaka/imersao/codepix-go/application/grpc/pb"
	"github.com/renatospaka/imersao/codepix-go/application/usecase"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)
	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error: err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult {
		Status: "created",
		Id: key.ID,
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := p.PixUseCase.FindKey(in.Key, in.Kind)
	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		id: pixKey.ID,
		kind: pixKey.Kind,
		key: pixKey.Key
		account: &pb.Account{
			accountId: pixKey.AccountID,
			accountNumber: pixKey.Account.Number,
			bankId: pixKey.Account.BankID,
			bankName: pixKey.Account.BankName, 
			ownerName: pixKey.Account.OwnerName,
			createdAt: pixKey.Account.CreatedAt.String(),
		}
		createdAt: pixKey.CreatedAt.String(),
	}, nil
}