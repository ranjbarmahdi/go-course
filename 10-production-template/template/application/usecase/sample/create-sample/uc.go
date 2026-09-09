package createsample

import (
	"context"
	"time"

	"template/application/contracts"
	apperrors "template/application/errors"
	"template/application/usecase/sample"
	"template/application/utils"
	domainsample "template/domain/sample"
	valueobjects "template/domain/value-objects"
)

type UseCase interface {
	Execute(ctx context.Context, req Request, user *valueobjects.OwnUser) (sample.SampleResponse, error)
}

type Implementation struct {
	sampleRepository domainsample.SampleRepository
	idGenerator      contracts.IDGenerator
	runInTx          contracts.RunInTx[error]
}

func New(
	sampleRepository domainsample.SampleRepository,
	idGenerator contracts.IDGenerator,
	runInTx contracts.RunInTx[error],
) *Implementation {
	return &Implementation{
		sampleRepository: sampleRepository,
		idGenerator:      idGenerator,
		runInTx:          runInTx,
	}
}

func (i *Implementation) Execute(ctx context.Context, req Request, user *valueobjects.OwnUser) (sample.SampleResponse, error) {
	if user == nil {
		return sample.SampleResponse{}, apperrors.New(apperrors.Unauthorized, "user is required")
	}
	if user.UserLvl() != valueobjects.SuperAdmin {
		return sample.SampleResponse{}, apperrors.New(apperrors.Unauthorized, "user is not authorized to create sample")
	}

	sampleID := valueobjects.SampleIDFromTrusted(i.idGenerator.NewUUID().String())

	entity, err := domainsample.NewSample(
		sampleID,
		&req.Name,
		&req.Number,
		time.Now(),
	)
	if err != nil {
		return sample.SampleResponse{}, utils.TranslateDomainError(err)
	}

	err = i.runInTx(ctx, func(txCtx context.Context) error {
		return i.sampleRepository.CreateSample(txCtx, entity)
	})
	if err != nil {
		return sample.SampleResponse{}, utils.TranslateDomainError(err)
	}

	return sample.ToAppResponse(entity), nil
}

var _ UseCase = (*Implementation)(nil)
