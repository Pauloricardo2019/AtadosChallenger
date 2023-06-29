package mocks

import (
	"atados/challenger/internal/model"
	"context"
	"github.com/stretchr/testify/mock"
)

type (
	VoluntaryRepositoryMock struct {
		mock.Mock
	}
	ActionRepositoryMock struct {
		mock.Mock
	}
)

func (p *VoluntaryRepositoryMock) Create(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	args := p.Called(ctx, voluntary)

	var voluntaryReq *model.Voluntary
	var err error

	if args.Get(0) != nil {
		voluntaryReq = args.Get(0).(*model.Voluntary)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaryReq, err
}

func (p *VoluntaryRepositoryMock) GetCount(ctx context.Context) (int64, error) {
	args := p.Called(ctx)

	var err error
	var count int64

	if args.Get(0) != nil {
		count = args.Get(0).(int64)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return count, err
}

func (p *VoluntaryRepositoryMock) GetByID(ctx context.Context, id uint64) (bool, *model.Voluntary, error) {
	args := p.Called(ctx, id)

	found := args.Get(0).(bool)

	var voluntary *model.Voluntary
	var err error

	if args.Get(1) != nil {
		voluntary = args.Get(1).(*model.Voluntary)
	}

	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}

	return found, voluntary, err
}

func (p *VoluntaryRepositoryMock) GetAll(ctx context.Context, limit, offset int) ([]model.Voluntary, error) {
	args := p.Called(ctx, limit, offset)

	var voluntaries []model.Voluntary
	var err error

	if args.Get(0) != nil {
		voluntaries = args.Get(0).([]model.Voluntary)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaries, err
}

func (p *VoluntaryRepositoryMock) Update(ctx context.Context, voluntary *model.Voluntary) (*model.Voluntary, error) {
	args := p.Called(ctx, voluntary)

	var voluntaryReq *model.Voluntary
	var err error

	if args.Get(0) != nil {
		voluntaryReq = args.Get(0).(*model.Voluntary)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaryReq, err
}

func (p *VoluntaryRepositoryMock) Delete(ctx context.Context, id uint64) error {
	args := p.Called(ctx, id)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}

func (p *ActionRepositoryMock) Create(ctx context.Context, action *model.Action) (*model.Action, error) {
	args := p.Called(ctx, action)

	var actionReq *model.Action
	var err error

	if args.Get(0) != nil {
		actionReq = args.Get(0).(*model.Action)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return actionReq, err
}

func (p *ActionRepositoryMock) GetCount(ctx context.Context) (int64, error) {
	args := p.Called(ctx)

	var err error
	var count int64

	if args.Get(0) != nil {
		count = args.Get(0).(int64)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return count, err
}

func (p *ActionRepositoryMock) GetByID(ctx context.Context, id uint64) (bool, *model.Action, error) {
	args := p.Called(ctx, id)

	found := args.Get(0).(bool)

	var action *model.Action
	var err error

	if args.Get(1) != nil {
		action = args.Get(1).(*model.Action)
	}

	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}

	return found, action, err
}

func (p *ActionRepositoryMock) GetAll(ctx context.Context, limit, offset int) ([]model.Action, error) {
	args := p.Called(ctx, limit, offset)

	var voluntaries []model.Action
	var err error

	if args.Get(0) != nil {
		voluntaries = args.Get(0).([]model.Action)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return voluntaries, err
}

func (p *ActionRepositoryMock) Update(ctx context.Context, action *model.Action) (*model.Action, error) {
	args := p.Called(ctx, action)

	var actionReq *model.Action
	var err error

	if args.Get(0) != nil {
		actionReq = args.Get(0).(*model.Action)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return actionReq, err
}

func (p *ActionRepositoryMock) Delete(ctx context.Context, id uint64) error {
	args := p.Called(ctx, id)

	var err error

	if args.Get(0) != nil {
		err = args.Get(0).(error)
	}
	return err
}
