package dataplane

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/repo"
)

// ShowFact godoc
// @tags Fact
// @summary Show a fact
// @description Show a fact by ID
// @id show-fact-by-id
// @accept json
// @produce json
// @security ApiKeyAuth
// @param id path string true "Fact ID"
// @success 200 {object} apimodel.Fact
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /facts/{id} [get]
func (dp *DataPlane) ShowFact(c echo.Context) error {
	ctx := c.Request().Context()
	f, err := dp.Repo.GetFact(ctx, &repo.GetFactOption{FactID: c.Param("id"), Domain: currentDomain(c)})
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	value, err := dp.Encryptor.Decrypt(f.Edges.Scope.Nonce, f.EncryptedValue)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	return c.JSON(http.StatusOK, apimodel.Fact{
		ID:            f.ID,
		ScopeCustomID: f.Edges.Scope.CustomID,
		FactTypeSlug:  f.Edges.FactType.Slug,
		Value:         value,
		Domain:        f.Domain,
	})
}

// CreateFact godoc
// @tags Fact
// @summary Create a fact
// @description create a fact
// @id create-fact
// @accept json
// @produce json
// @security ApiKeyAuth
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

	domain := currentDomain(c)

	s, err := dp.Repo.CreateScope(ctx, &repo.CreateScopeOption{
		ScopeCustomID: cf.ScopeCustomID,
		Domain:        domain,
	})
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	ft, err := dp.Repo.CreateFactType(ctx, &repo.CreateFactTypeOption{
		FactTypeSlug:       cf.FactTypeSlug,
		FactTypeValidation: "",
		BuiltIn:            true,
	})
	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	encryptedValue, err := dp.Encryptor.Encrypt(s.Nonce, cf.Value)
	if err != nil {
		return apimodel.NewEntError(c, err)
	}
	hashedValue := dp.Hasher.Hash(cf.Value, domain)

	f, err := dp.Repo.CreateFact(ctx, &repo.CreateFactOption{
		Domain:         domain,
		Scope:          s,
		FactType:       ft,
		HashedValue:    hashedValue,
		EncryptedValue: encryptedValue,
	})

	if err != nil {
		return apimodel.NewEntError(c, err)
	}

	return c.JSON(http.StatusOK, apimodel.Fact{
		ID:            f.ID,
		ScopeCustomID: s.CustomID,
		FactTypeSlug:  ft.Slug,
		Domain:        f.Domain,
	})
}
