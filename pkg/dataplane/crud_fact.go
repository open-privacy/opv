package dataplane

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
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

	if err := dp.validateFactType(ctx, ft.Slug, cf.Value); err != nil {
		return apimodel.NewHTTPError(c, err, http.StatusBadRequest)
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

var builtInFactTypeValuations map[string]interface{}

func init() {
	customValidationTags := map[string]func(str string) bool{
		"phonenumber": func(str string) bool {
			return govalidator.StringMatches(str, `^((\+?[0-9]{1,3})|(\+?\([0-9]{1,3}\)))[\s-]?(?:\(0?[0-9]{1,5}\)|[0-9]{1,5})[-\s]?[0-9][\d\s-]{5,7}\s?(?:x[\d-]{0,4})?$`)
		},
		"ssn": func(str string) bool {
			return govalidator.StringMatches(str, `^\d{3}[- ]?\d{2}[- ]?\d{4}$`)
		},
		"ssnstrict": func(str string) bool {
			return govalidator.StringMatches(str, `^(?!666|000|9\d{2})\d{3}-(?!00)\d{2}-(?!0{4})\d{4}$`)
		},
	}
	customValidationMap := map[string]interface{}{
		"address": map[string]interface{}{
			"name":            "type(string)",
			"phone":           "alphanum",
			"company":         "type(string)",
			"email":           "type(string)",
			"address_line1":   "type(string)",
			"address_line2":   "type(string)",
			"address_city":    "type(string)",
			"address_state":   "type(string)",
			"address_zip":     "type(string)",
			"address_country": "type(string)",
		},
	}

	for k, v := range customValidationTags {
		govalidator.TagMap[k] = govalidator.Validator(v)
	}

	builtInFactTypeValuations = make(map[string]interface{})
	for k := range govalidator.TagMap {
		builtInFactTypeValuations[k] = k
	}
	for k, v := range customValidationMap {
		builtInFactTypeValuations[k] = v
	}
}

func (dp *DataPlane) validateFactType(ctx context.Context, factTypeSlug string, factValue string) error {
	rule, ok := builtInFactTypeValuations[factTypeSlug]
	if !ok {
		return fmt.Errorf("not supported fact type slug: %s", factTypeSlug)
	}

	if _, ok := rule.(string); ok {
		input := map[string]interface{}{"value": factValue}
		ruleMap := map[string]interface{}{"value": rule}
		_, err := govalidator.ValidateMap(input, ruleMap)
		return err
	}

	var factValueMap map[string]interface{}
	err := json.Unmarshal([]byte(factValue), &factValueMap)
	if err != nil {
		return err
	}
	input := map[string]interface{}{"value": factValueMap}
	ruleMap := map[string]interface{}{"value": rule}
	_, err = govalidator.ValidateMap(input, ruleMap)
	return err
}
