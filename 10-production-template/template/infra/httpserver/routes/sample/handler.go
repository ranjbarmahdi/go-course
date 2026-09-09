package sample

import (
	"net/http"

	createsample "template/application/usecase/sample/create-sample"
	getsample "template/application/usecase/sample/get-sample"
	valueobjects "template/domain/value-objects"
	"template/infra/httpserver/middlewares"
	"template/infra/httpserver/utils"
)

type Handler struct {
	createSample createsample.UseCase
	getSample    getsample.UseCase
}

func NewHandler(
	createSample createsample.UseCase,
	getSample getsample.UseCase,
) *Handler {
	return &Handler{
		createSample: createSample,
		getSample:    getSample,
	}
}

func ToAppUser(user middlewares.OwnUser) *valueobjects.OwnUser {
	u := valueobjects.NewOwnUser(
		user.Username,
		user.FirstName,
		user.LastName,
		user.ID,
		user.Admin,
		valueobjects.UserLvl(user.UserLvl),
		user.OrganizationID,
	)
	return &u
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	user, ok := middlewares.UserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "invalid user")
		return
	}

	var req CreateSampleRequest
	if err := utils.DecodeAndValidate(r, &req); err != nil {
		utils.WriteValidationError(w, err)
		return
	}

	result, err := h.createSample.Execute(r.Context(), createsample.Request{
		Name:   req.Name,
		Number: req.Number,
	}, ToAppUser(user))
	if err != nil {
		utils.WriteAppError(w, err)
		return
	}

	utils.WriteSuccess(w, http.StatusCreated, "sample created", ToApiResponse(result))
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	user, ok := middlewares.UserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, "invalid user")
		return
	}

	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, "id is required")
		return
	}

	result, err := h.getSample.Execute(r.Context(), id, ToAppUser(user))
	if err != nil {
		utils.WriteAppError(w, err)
		return
	}

	utils.WriteSuccess(w, http.StatusOK, "sample found", ToApiResponse(result))
}
