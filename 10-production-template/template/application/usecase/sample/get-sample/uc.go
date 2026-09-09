package getsample

import (
	"context"
	"template/application/contracts"
	apperrors "template/application/errors"
	"template/application/usecase/sample"
	domainsample "template/domain/sample"
	valueobjects "template/domain/value-objects"
)

type UseCase interface {
	Execute(ctx context.Context, id string, user *valueobjects.OwnUser) (sample.SampleResponse, error)
}

type Implementation struct {
	sampleRepository domainsample.SampleRepository
	idGenerator      contracts.IDGenerator
}

func (i *Implementation) Execute(ctx context.Context, id string, user *valueobjects.OwnUser) (sample.SampleResponse, error) {
	if user.UserLvl() != valueobjects.SuperAdmin {
		return sample.SampleResponse{}, apperrors.New(apperrors.Unauthorized, "user is not authorized to get sample")
	}

	validatedID, err := i.idGenerator.ParseUUID(id)

	if err != nil {
		return sample.SampleResponse{}, apperrors.New(apperrors.InvalidInput, "invalid sample id")
	}

	sampleID := valueobjects.SampleIDFromTrusted(validatedID.String())
	res, err := i.sampleRepository.FindByID(ctx, sampleID)

	if err != nil {
		return sample.SampleResponse{}, apperrors.New(apperrors.NotFound, "sample not found")
	}

	return sample.ToAppResponse(*res), nil
}

func New(
	sampleRepository domainsample.SampleRepository,
	idGenerator contracts.IDGenerator,
) *Implementation {
	return &Implementation{
		sampleRepository: sampleRepository,
		idGenerator:      idGenerator,
	}
}

var _ UseCase = &Implementation{}
