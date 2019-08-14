package squad

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"go.zenithar.org/pkg/db"
	"go.zenithar.org/pkg/log"

	"go.zenithar.org/spotigraph/internal/models"
	"go.zenithar.org/spotigraph/internal/repositories"
	"go.zenithar.org/spotigraph/internal/services"
	"go.zenithar.org/spotigraph/internal/services/internal/constraints"
	"go.zenithar.org/spotigraph/pkg/protocol/v1/spotigraph"
)

type service struct {
	squads      repositories.Squad
	users       repositories.User
	memberships repositories.Membership
}

// New returns a service instance
func New(squads repositories.Squad, users repositories.User, memberships repositories.Membership) services.Squad {
	return &service{
		squads:      squads,
		users:       users,
		memberships: memberships,
	}
}

// -----------------------------------------------------------------------------

func (s *service) Create(ctx context.Context, req *spotigraph.SquadCreateReq) (*spotigraph.SingleSquadRes, error) {
	res := &spotigraph.SingleSquadRes{}

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Name must be unique
		constraints.SquadNameMustBeUnique(s.squads, req.Name),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		}
		return res, err
	}

	// Prepare squad creation
	entity := models.NewSquad(req.Name)

	// Create use in database
	if err := s.squads.Create(ctx, entity); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to create squad",
		}
		return res, err
	}

	// Prepare response
	res.Entity = spotigraph.FromSquad(entity)

	// Return result
	return res, nil
}

func (s *service) Get(ctx context.Context, req *spotigraph.SquadGetReq) (*spotigraph.SingleSquadRes, error) {
	res := &spotigraph.SingleSquadRes{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Retrieve squad from database
	entity, err := s.squads.Get(ctx, req.Id)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to retrieve squad",
		}
		return res, err
	}
	if entity == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusNotFound,
			Message: "Squad not found",
		}
		return res, db.ErrNoResult
	}

	// Prepare response
	res.Entity = spotigraph.FromSquad(entity)

	// Return result
	return res, nil
}

func (s *service) Update(ctx context.Context, req *spotigraph.SquadUpdateReq) (*spotigraph.SingleSquadRes, error) {
	res := &spotigraph.SingleSquadRes{}

	// Prepare expected results

	var entity models.Squad

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Squad must exists
		constraints.SquadMustExists(s.squads, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	updated := false

	if req.Name != nil {
		if err := constraints.Validate(ctx,
			// Check acceptable name value
			constraints.MustBeAName(req.Name.Value),
			// Is already used ?
			constraints.SquadNameMustBeUnique(s.squads, req.Name.Value),
		); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusConflict,
				Message: err.Error(),
			}
			return res, err
		}
		entity.Name = req.Name.Value
		updated = true
	}

	// Skip operation when no updates
	if updated {
		// Create account in database
		if err := s.squads.Update(ctx, &entity); err != nil {
			res.Error = &spotigraph.Error{
				Code:    http.StatusInternalServerError,
				Message: "Unable to update squad object",
			}
			return res, err
		}
	}

	// Prepare response
	res.Entity = spotigraph.FromSquad(&entity)

	// Return expected result
	return res, nil
}

func (s *service) Delete(ctx context.Context, req *spotigraph.SquadGetReq) (*spotigraph.EmptyRes, error) {
	res := &spotigraph.EmptyRes{}

	// Prepare expected results

	var entity models.Squad

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Squad must exists
		constraints.SquadMustExists(s.squads, req.Id, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	if err := s.squads.Delete(ctx, req.Id); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to delete squad object",
		}
		return res, err
	}

	// Return expected result
	return res, nil
}

func (s *service) Search(ctx context.Context, req *spotigraph.SquadSearchReq) (*spotigraph.PaginatedSquadRes, error) {
	res := &spotigraph.PaginatedSquadRes{}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must not be nil
		constraints.MustNotBeNil(req, "Request must not be nil"),
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Prepare request parameters
	sortParams := db.SortConverter(req.Sorts)
	pagination := db.NewPaginator(uint(req.Page), uint(req.PerPage))

	// Build search filter
	filter := &repositories.SquadSearchFilter{}
	if req.SquadId != nil {
		filter.SquadID = req.SquadId.Value
	}
	if req.Name != nil {
		filter.Name = req.Name.Value
	}

	// Do the search
	entities, total, err := s.squads.Search(ctx, filter, pagination, sortParams)
	if err != nil && err != db.ErrNoResult {
		res.Error = &spotigraph.Error{
			Code:    http.StatusInternalServerError,
			Message: "Unable to process request",
		}
		return res, err
	}

	// Set pagination total for paging calcul
	pagination.SetTotal(uint(total))
	res.Total = uint32(pagination.Total())
	res.Count = uint32(pagination.CurrentPageCount())
	res.PerPage = uint32(pagination.PerPage)
	res.CurrentPage = uint32(pagination.Page)

	// If no result back to first page
	if err != db.ErrNoResult {
		res.Members = spotigraph.FromSquads(entities)
	}

	// Return results
	return res, nil
}

func (s *service) AddMembers(ctx context.Context, req *spotigraph.SquadMemberReq) (*spotigraph.EmptyRes, error) {
	res := &spotigraph.EmptyRes{}

	// Prepare expected results

	var entity models.Squad

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Squad must exists
		constraints.SquadMustExists(s.squads, req.SquadId, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Check user existence
	valid := []*models.User{}
	for _, uid := range req.UserIds {
		user, err := s.users.Get(ctx, uid)
		if err != nil {
			log.For(ctx).Error("Invalid user, ignored", zap.Error(err))
			continue
		}

		// Add to valid users
		valid = append(valid, user)
	}

	if len(valid) > 0 {
		for _, u := range valid {
			if err := s.memberships.Join(ctx, u, &entity); err != nil {
				log.For(ctx).Error("Invalid user, ignored", zap.Error(err))
				continue
			}
		}
	}

	// Return expected result
	return res, nil
}

func (s *service) RemoveMembers(ctx context.Context, req *spotigraph.SquadMemberReq) (*spotigraph.EmptyRes, error) {
	res := &spotigraph.EmptyRes{}

	// Prepare expected results

	var entity models.Squad

	// Check request
	if req == nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusBadRequest,
			Message: "request must not be nil",
		}
		return res, fmt.Errorf("request must not be nil")
	}

	// Validate service constraints
	if err := constraints.Validate(ctx,
		// Request must be syntaxically valid
		constraints.MustBeValid(req),
		// Squad must exists
		constraints.SquadMustExists(s.squads, req.SquadId, &entity),
	); err != nil {
		res.Error = &spotigraph.Error{
			Code:    http.StatusPreconditionFailed,
			Message: "Unable to validate request",
		}
		return res, err
	}

	// Check user existence
	valid := []*models.User{}
	for _, uid := range req.UserIds {
		user, err := s.users.Get(ctx, uid)
		if err != nil {
			log.For(ctx).Error("Invalid user, ignored", zap.Error(err))
			continue
		}

		// Add to valid users
		valid = append(valid, user)
	}

	if len(valid) > 0 {
		for _, u := range valid {
			if err := s.memberships.Leave(ctx, u, &entity); err != nil {
				log.For(ctx).Error("Invalid user, ignored", zap.Error(err))
				continue
			}
		}
	}

	// Return expected result
	return res, nil
}
