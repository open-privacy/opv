package dataplane

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-privacy/opv/pkg/apimodel"
	"github.com/open-privacy/opv/pkg/repo"
)

// QueryFactTypes godoc
// @tags Fact
// @summary Query fact types
// @description Query fact types
// @id show-fact-types
// @produce json
// @security ApiKeyAuth
// @param slug query string false "Fact Type Slug"
// @param builtin query boolean false "Builtin Fact Type Slug"
// @success 200 {object} []apimodel.FactType
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /fact_types [get]
func (dp *DataPlane) QueryFactTypes(c echo.Context) error {
	if c.QueryParam("slug") != "" {
		ctx := c.Request().Context()
		ft, err := dp.Repo.GetFactTypeBySlug(c.Request().Context(), c.QueryParam("slug"))
		if err != nil {
			return dp.Repo.HandleError(ctx, err)
		}
		return c.JSON(http.StatusOK, []apimodel.FactType{
			{
				ID:         ft.ID,
				Slug:       ft.Slug,
				Validation: ft.Validation,
				BuiltIn:    ft.BuiltIn,
			},
		})
	}

	if c.QueryParam("builtin") == "false" {
		return c.JSON(http.StatusOK, []apimodel.FactType{})
	}
	return c.JSON(http.StatusOK, dp.queryBuiltInFactTypes())
}

func (dp *DataPlane) queryBuiltInFactTypes() []apimodel.FactType {
	fts := []apimodel.FactType{}
	for slug, rule := range builtInFactTypeValuations {
		validation, _ := json.Marshal(rule)
		fts = append(fts, apimodel.FactType{
			Slug:       slug,
			BuiltIn:    true,
			Validation: string(validation),
		})
	}
	return fts
}

// CreateFactType godoc
// @tags Fact
// @summary Create a fact type
// @description create a fact type
// @id create-fact-type
// @accept json
// @produce json
// @security ApiKeyAuth
// @param createFact body apimodel.CreateFactType true "Create Fact Type Parameters"
// @success 200 {object} apimodel.CreateFactType
// @failure 400 {object} apimodel.HTTPError
// @failure 500 {object} apimodel.HTTPError
// @router /fact_types [post]
func (dp *DataPlane) CreateFactType(c echo.Context) error {
	ctx := c.Request().Context()
	var cft apimodel.CreateFactType
	if err := c.Bind(&cft); err != nil {
		return apimodel.FormatHTTPError(c, apimodel.ErrJSONMalformatted)
	}

	ft, err := dp.Repo.CreateFactType(c.Request().Context(), &repo.CreateFactTypeOption{
		FactTypeSlug:       cft.Slug,
		FactTypeValidation: cft.Validation,
		BuiltIn:            true,
	})
	if err != nil {
		return dp.Repo.HandleError(ctx, err)
	}

	return c.JSON(http.StatusOK, apimodel.FactType{
		ID:         ft.ID,
		Slug:       ft.Slug,
		Validation: ft.Validation,
		BuiltIn:    ft.BuiltIn,
	})
}
