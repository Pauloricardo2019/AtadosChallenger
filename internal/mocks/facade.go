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
	ActionFacadeMock struct {
		mock.Mock
	}
)

func (p *VoluntaryFacadeMock) CreateVoluntary(ctx context.Context, voluntaryRequest *dto.CreateVoluntaryRequest) (*dto.CreateVoluntaryResponse, error) {
	args := p.Called(ctx, voluntaryRequest)

	var voluntaryReq *dto.CreateVoluntaryResponse
	var err error

	if args.Get(0) != nil {
		voluntaryReq = args.Get(0).(*dto.CreateVoluntaryResponse)
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

func (p *ActionFacadeMock) CreateAction(ctx context.Context, actionRequest *dto.CreateActionRequest) (*dto.CreateActionResponse, error) {
	args := p.Called(ctx, actionRequest)

	var actionReq *dto.CreateActionResponse
	var err error

	if args.Get(0) != nil {
		actionReq = args.Get(0).(*dto.CreateActionResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return actionReq, err
}

func (p *ActionFacadeMock) GetActionByID(ctx context.Context, id uint64) (*dto.GetActionByIDResponse, error) {
	args := p.Called(ctx, id)

	var actionReq *dto.GetActionByIDResponse
	var err error

	if args.Get(0) != nil {
		actionReq = args.Get(0).(*dto.GetActionByIDResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return actionReq, err
}

func (p *ActionFacadeMock) GetAllActions(ctx context.Context, limit, offset int) (*dto.GetAllActionsResponse, error) {
	args := p.Called(ctx, limit, offset)

	var actionReq *dto.GetAllActionsResponse
	var err error

	if args.Get(0) != nil {
		actionReq = args.Get(0).(*dto.GetAllActionsResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return actionReq, err
}

func (p *ActionFacadeMock) UpdateAction(ctx context.Context, actionID uint64, actionRequest *dto.UpdateActionRequest) error {
	args := p.Called(ctx, actionID, actionRequest)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err

}

func (p *ActionFacadeMock) DeleteAction(ctx context.Context, id uint64) error {
	args := p.Called(ctx, id)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}

	return err
}
