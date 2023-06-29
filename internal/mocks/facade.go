package mocks

import (
	"atados/challenger/internal/dto"
	"context"
	"github.com/stretchr/testify/mock"
)

type (
	VoluntaryFacadeMock struct {
		mock.Mock
	}
)

func (p *VoluntaryFacadeMock) CreateVoluntary(ctx context.Context, voluntaryRequest *dto.CreateVoluntaryRequest) (*dto.CreateVoluntaryVO, error) {
	args := p.Called(ctx, voluntaryRequest)

	var voluntaryReq *dto.CreateVoluntaryVO
	var err error

	if args.Get(0) != nil {
		voluntaryReq = args.Get(0).(*dto.CreateVoluntaryVO)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaryReq, err
}

func (p *VoluntaryFacadeMock) GetVoluntaryByID(ctx context.Context, id uint64) (*dto.GetVoluntaryByIDResponse, error) {
	args := p.Called(ctx, id)

	var voluntaryReq *dto.GetVoluntaryByIDResponse
	var err error

	if args.Get(0) != nil {
		voluntaryReq = args.Get(0).(*dto.GetVoluntaryByIDResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaryReq, err
}

func (p *VoluntaryFacadeMock) GetAllVoluntaries(ctx context.Context, limit, offset int) (*dto.GetAllVoluntariesResponse, error) {
	args := p.Called(ctx, limit, offset)

	var voluntaryReq *dto.GetAllVoluntariesResponse
	var err error

	if args.Get(0) != nil {
		voluntaryReq = args.Get(0).(*dto.GetAllVoluntariesResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaryReq, err
}

func (p *VoluntaryFacadeMock) UpdateVoluntary(ctx context.Context, voluntaryID uint64, voluntaryRequest *dto.UpdateVoluntaryRequest) error {
	args := p.Called(ctx, voluntaryID, voluntaryRequest)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err

}

func (p *VoluntaryFacadeMock) DeleteVoluntary(ctx context.Context, id uint64) error {
	args := p.Called(ctx, id)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}
