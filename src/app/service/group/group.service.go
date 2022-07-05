package group

import (
	"context"
	"github.com/google/uuid"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model"
	"github.com/isd-sgcu/rnkm65-backend/src/app/model/group"
	"github.com/isd-sgcu/rnkm65-backend/src/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	repo IRepository
}

type IRepository interface {
	FindOne(string, *group.Group) error
	FindByToken(string, *group.Group) error
	Create(*group.Group) error
	Update(string, *group.Group) error
	Delete(string) error
}

func NewService(repo IRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindOne(_ context.Context, req *proto.FindOneGroupRequest) (res *proto.FindOneGroupResponse, err error) {
	raw := group.Group{}

	err = s.repo.FindOne(req.Id, &raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}
	return &proto.FindOneGroupResponse{Group: RawToDto(&raw)}, nil
}

func (s *Service) FindByToken(_ context.Context, req *proto.FindByTokenGroupRequest) (res *proto.FindByTokenGroupResponse, err error) {
	raw := group.Group{}

	err = s.repo.FindByToken(req.Token, &raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}
	return &proto.FindByTokenGroupResponse{Group: RawToDto(&raw)}, nil
}

func (s *Service) Create(_ context.Context, req *proto.CreateGroupRequest) (res *proto.CreateGroupResponse, err error) {
	raw, _ := DtoToRaw(req.Group)

	err = s.repo.Create(raw)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create group")
	}

	return &proto.CreateGroupResponse{Group: RawToDto(raw)}, nil
}

func (s *Service) Update(_ context.Context, req *proto.UpdateGroupRequest) (res *proto.UpdateGroupResponse, err error) {
	raw, err := DtoToRaw(req.Group)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid group id")
	}

	err = s.repo.Update(req.Group.Id, raw)
	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}

	return &proto.UpdateGroupResponse{Group: RawToDto(raw)}, nil
}

func (s *Service) Delete(_ context.Context, req *proto.DeleteGroupRequest) (res *proto.DeleteGroupResponse, err error) {
	err = s.repo.Delete(req.Id)

	if err != nil {
		return nil, status.Error(codes.NotFound, "group not found")
	}
	return &proto.DeleteGroupResponse{Success: true}, nil
}

func DtoToRaw(in *proto.Group) (result *group.Group, err error) {
	var id uuid.UUID

	if in.Id != "" {
		id, err = uuid.Parse(in.Id)
		if err != nil {
			return nil, err
		}
	}

	return &group.Group{
		Base: model.Base{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		LeaderID: in.LeaderID,
		Token:    in.Token,
	}, nil
}

func RawToDto(in *group.Group) *proto.Group {
	return &proto.Group{
		Id:       in.ID.String(),
		LeaderID: in.LeaderID,
		Token:    in.Token,
	}
}
