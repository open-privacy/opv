package dataplane

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFactType(t *testing.T) {
	dp := &DataPlane{}
	ctx := context.Background()

	tests := []struct {
		factTypeSlug string
		factValue    string
		valid        bool
	}{
		// phonenumber
		{factTypeSlug: "phonenumber", factValue: "3979435680", valid: true},
		{factTypeSlug: "phonenumber", factValue: "397-943-5680", valid: true},

		// photourl
		{factTypeSlug: "photourl", factValue: "https://example.com/photo.png", valid: true},

		// ssn
		{factTypeSlug: "ssn", factValue: "123-45-6789", valid: true},
		{factTypeSlug: "ssn", factValue: "123 45 6789", valid: true},
		{factTypeSlug: "ssn", factValue: "123456789", valid: true},
		{factTypeSlug: "ssn", factValue: "1234", valid: false},
		{factTypeSlug: "ssn", factValue: "666-00-0000", valid: true},
		{factTypeSlug: "ssnstrict", factValue: "1234", valid: false},
		{factTypeSlug: "ssnstrict", factValue: "666-00-0000", valid: false},

		// ipv4
		{factTypeSlug: "ipv4", factValue: "10.0.0.1", valid: true},
		{factTypeSlug: "ipv4", factValue: "10.1", valid: false},
		{factTypeSlug: "ipv4", factValue: "2001:db8:0000:1:1:1:1:1", valid: false},

		// ipv6
		{factTypeSlug: "ipv6", factValue: "2001:db8:0000:1:1:1:1:1", valid: true},
		{factTypeSlug: "ipv6", factValue: "10.1", valid: false},
		{factTypeSlug: "ipv6", factValue: "10.0.0.1", valid: false},

		// email
		{factTypeSlug: "email", factValue: "me@example.com", valid: true},
		{factTypeSlug: "email", factValue: "@example.com", valid: false},

		// address
		{
			factTypeSlug: "address",
			factValue: `{
				"name": "JOHN DOE",
				"phone": "5555555555",
				"company": "Example",
				"email": "john@example.com",
				"address_line1": "1 EXAMPLE ST STE 2000",
				"address_line2": null,
				"address_city": "SAN FRANCISCO",
				"address_state": "CA",
				"address_zip": "94107-1741",
				"address_country": "UNITED STATES"
			}`,
			valid: true,
		},
		{
			factTypeSlug: "address",
			// invalid address_zip not set as string
			factValue: `{
				"name": "JOHN DOE",
				"phone": "5555555555",
				"company": "Example",
				"email": "john@example.com",
				"address_line1": "1 EXAMPLE ST STE 2000",
				"address_line2": null,
				"address_city": "SAN FRANCISCO",
				"address_state": "CA",
				"address_zip": 123,
				"address_country": "UNITED STATES"
			}`,
			valid: false,
		},
	}

	for _, test := range tests {
		err := dp.validateFactType(ctx, test.factTypeSlug, test.factValue)
		assert.Equal(t, test.valid, err == nil, fmt.Sprintf("err:%s, test:%+v", err, test))
	}
}
