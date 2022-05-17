package handler

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	validate "github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"net/http"
)

type OrganizationHandler struct {
	service  OrganizationService
	validate *validate.DtoValidator
}

func NewOrganizationHandler(service OrganizationService, validate *validate.DtoValidator) *OrganizationHandler {
	return &OrganizationHandler{
		service:  service,
		validate: validate,
	}
}

type OrganizationContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID() (int32, error)
	PaginationQueryParam(*dto.PaginationQueryParams) error
}

type OrganizationService interface {
	FindAll(*dto.PaginationQueryParams) (*proto.OrganizationPagination, *dto.ResponseErr)
	FindOne(int32) (*proto.Organization, *dto.ResponseErr)
	Create(*dto.OrganizationDto) (*proto.Organization, *dto.ResponseErr)
	Update(int32, *dto.OrganizationDto) (*proto.Organization, *dto.ResponseErr)
	Delete(int32) (*proto.Organization, *dto.ResponseErr)
}

// FindAll is a function that get all organizations in database
// @Summary Get all organizations
// @Description Return the arrays of organization dto if successfully
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} dto.ResponseErr "Invalid query param"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /organization [get]
func (h *OrganizationHandler) FindAll(c OrganizationContext) {
	query := dto.PaginationQueryParams{}

	err := c.PaginationQueryParam(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Cannot parse query param",
		})
		return
	}

	organizations, errRes := h.service.FindAll(&query)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, organizations)
	return
}

// FindOne is a function that get the specific organizations with id
// @Summary Get specific organization with id
// @Description Return the organization dto if successfully
// @Param id path int true "id"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found organization"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /organization/{id} [get]
func (h *OrganizationHandler) FindOne(c OrganizationContext) {

	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	organization, errRes := h.service.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, organization)
	return
}

// Create is a function that create the organization
// @Summary Create the organization
// @Description Return the organization dto if successfully
// @Param organization body dto.OrganizationDto true "organization dto"
// @Tags organization
// @Accept json
// @Produce json
// @Success 201 {object} proto.Organization
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found organization"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /organization [post]
func (h *OrganizationHandler) Create(c OrganizationContext) {
	organizationDto := dto.OrganizationDto{}
	err := c.Bind(&organizationDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse organization dto",
		})
		return
	}

	if errors := h.validate.Validate(organizationDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	organization, errRes := h.service.Create(&organizationDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, organization)
	return
}

// Update is a function that update the organization
// @Summary Update the existing organization
// @Description Return the organization dto if successfully
// @Param id path int true "id"
// @Param organization body dto.OrganizationDto true "organization dto"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found organization"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /organization/{id} [patch]
func (h *OrganizationHandler) Update(c OrganizationContext) {
	organizationDto := dto.OrganizationDto{}
	err := c.Bind(&organizationDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse organization dto",
		})
		return
	}

	if errors := h.validate.Validate(organizationDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	organization, errRes := h.service.Update(id, &organizationDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, organization)
	return
}

// Delete is a function that delete the organization
// @Summary Delete the organization
// @Description Return the organization dto if successfully
// @Param id path int true "id"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found organization"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /organization/{id} [delete]
func (h *OrganizationHandler) Delete(c OrganizationContext) {

	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	organization, errRes := h.service.Delete(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, organization)
	return
}
