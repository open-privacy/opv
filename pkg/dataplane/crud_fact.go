package dataplane

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/ent"
	"github.com/open-privacy/opv/pkg/ent/fact"
	"github.com/open-privacy/opv/pkg/ent/facttype"
	"github.com/open-privacy/opv/pkg/ent/scope"
)

// ShowFact godoc
// @tags Fact
// @summary Show a fact
// @description Show a fact by ID
// @id show-fact-by-id
// @accept  json
// @produce  json
// @param id path string true "Fact ID"
// @success 200 {object} apimodel.Fact
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /facts/{id} [get]
func (dp *DataPlane) ShowFact(c echo.Context) error {
	ctx := c.Request().Context()
	f, err := dp.EntClient.Fact.Query().WithScope().WithFactType().Where(fact.ID(c.Param("id"))).Only(ctx)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	s, err := f.Edges.ScopeOrErr()
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	ft, err := f.Edges.FactTypeOrErr()
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	value, err := dp.Encryptor.Decrypt(s.Nonce, f.EncryptedValue)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	return c.JSON(http.StatusOK, apimodel.Fact{
		ID:            f.ID,
		ScopeCustomID: s.CustomID,
		FactTypeSlug:  ft.Slug,
		Value:         value,
	})
}

// CreateFact godoc
// @tags Fact
// @summary Create a fact
// @description create a fact
// @id create-fact
// @accept  json
// @produce  json
// @param createFact body apimodel.CreateFact true "Create Fact Parameters"
// @success 200 {object} apimodel.Fact
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /facts [post]
func (dp *DataPlane) CreateFact(c echo.Context) error {
	ctx := c.Request().Context()
	cf := &apimodel.CreateFact{}
	err := c.Bind(cf)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	s, err := dp.createScopeIfNotExists(ctx, cf.ScopeCustomID)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	ft, err := dp.createFactTypeIfNotExists(ctx, cf.FactTypeSlug)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	encryptedValue, err := dp.Encryptor.Encrypt(s.Nonce, cf.Value)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}
	hashedValue := dp.Hasher.Hash(cf.Value)

	f, err := dp.EntClient.Fact.
		Create().
		SetScope(s).
		SetFactType(ft).
		SetEncryptedValue(encryptedValue).
		SetHashedValue(hashedValue).
		Save(ctx)

	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	return c.JSON(http.StatusOK, apimodel.Fact{
		ID:            f.ID,
		ScopeCustomID: s.CustomID,
		FactTypeSlug:  ft.Slug,
		Value:         cf.Value,
	})
}

func (dp *DataPlane) createScopeIfNotExists(ctx context.Context, scopeCustomID string) (*ent.Scope, error) {
	s, err := dp.EntClient.Scope.Query().Where(scope.CustomID(scopeCustomID)).Only(ctx)
	if ent.IsNotFound(err) {
		s, err = dp.EntClient.Scope.Create().SetCustomID(scopeCustomID).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return s, err
}

func (dp *DataPlane) createFactTypeIfNotExists(ctx context.Context, factTypeSlug string) (*ent.FactType, error) {
	ft, err := dp.EntClient.FactType.Query().Where(facttype.Slug(factTypeSlug)).Only(ctx)
	if ent.IsNotFound(err) {
		ft, err = dp.EntClient.FactType.Create().SetSlug(factTypeSlug).Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	return ft, err
}
