// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// APIAuditsColumns holds the columns for the "api_audits" table.
	APIAuditsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "plane", Type: field.TypeString},
		{Name: "hashed_grant_token", Type: field.TypeString, Nullable: true},
		{Name: "domain", Type: field.TypeString, Nullable: true},
		{Name: "http_path", Type: field.TypeString, Nullable: true},
		{Name: "http_method", Type: field.TypeString, Nullable: true},
		{Name: "sent_http_status", Type: field.TypeInt, Nullable: true},
	}
	// APIAuditsTable holds the schema information for the "api_audits" table.
	APIAuditsTable = &schema.Table{
		Name:        "api_audits",
		Columns:     APIAuditsColumns,
		PrimaryKey:  []*schema.Column{APIAuditsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "apiaudit_id",
				Unique:  true,
				Columns: []*schema.Column{APIAuditsColumns[0]},
			},
			{
				Name:    "apiaudit_created_at",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[1]},
			},
			{
				Name:    "apiaudit_updated_at",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[2]},
			},
			{
				Name:    "apiaudit_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[3]},
			},
			{
				Name:    "apiaudit_plane",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[4]},
			},
			{
				Name:    "apiaudit_hashed_grant_token",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[5]},
			},
			{
				Name:    "apiaudit_domain",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[6]},
			},
			{
				Name:    "apiaudit_http_path",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[7]},
			},
			{
				Name:    "apiaudit_http_method",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[8]},
			},
			{
				Name:    "apiaudit_sent_http_status",
				Unique:  false,
				Columns: []*schema.Column{APIAuditsColumns[9]},
			},
		},
	}
	// FactsColumns holds the columns for the "facts" table.
	FactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "hashed_value", Type: field.TypeString},
		{Name: "encrypted_value", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "fact_type_facts", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "scope_facts", Type: field.TypeString, Nullable: true, Size: 255},
	}
	// FactsTable holds the schema information for the "facts" table.
	FactsTable = &schema.Table{
		Name:       "facts",
		Columns:    FactsColumns,
		PrimaryKey: []*schema.Column{FactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "facts_fact_types_facts",
				Columns:    []*schema.Column{FactsColumns[7]},
				RefColumns: []*schema.Column{FactTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "facts_scopes_facts",
				Columns:    []*schema.Column{FactsColumns[8]},
				RefColumns: []*schema.Column{ScopesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "fact_id",
				Unique:  true,
				Columns: []*schema.Column{FactsColumns[0]},
			},
			{
				Name:    "fact_created_at",
				Unique:  false,
				Columns: []*schema.Column{FactsColumns[1]},
			},
			{
				Name:    "fact_updated_at",
				Unique:  false,
				Columns: []*schema.Column{FactsColumns[2]},
			},
			{
				Name:    "fact_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{FactsColumns[3]},
			},
			{
				Name:    "fact_hashed_value",
				Unique:  false,
				Columns: []*schema.Column{FactsColumns[4]},
			},
			{
				Name:    "fact_domain",
				Unique:  false,
				Columns: []*schema.Column{FactsColumns[6]},
			},
		},
	}
	// FactTypesColumns holds the columns for the "fact_types" table.
	FactTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "slug", Type: field.TypeString},
		{Name: "built_in", Type: field.TypeBool, Default: false},
		{Name: "validation", Type: field.TypeString, Nullable: true},
	}
	// FactTypesTable holds the schema information for the "fact_types" table.
	FactTypesTable = &schema.Table{
		Name:        "fact_types",
		Columns:     FactTypesColumns,
		PrimaryKey:  []*schema.Column{FactTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "facttype_id",
				Unique:  true,
				Columns: []*schema.Column{FactTypesColumns[0]},
			},
			{
				Name:    "facttype_created_at",
				Unique:  false,
				Columns: []*schema.Column{FactTypesColumns[1]},
			},
			{
				Name:    "facttype_updated_at",
				Unique:  false,
				Columns: []*schema.Column{FactTypesColumns[2]},
			},
			{
				Name:    "facttype_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{FactTypesColumns[3]},
			},
			{
				Name:    "facttype_slug",
				Unique:  true,
				Columns: []*schema.Column{FactTypesColumns[4]},
			},
		},
	}
	// GrantsColumns holds the columns for the "grants" table.
	GrantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "hashed_grant_token", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
		{Name: "version", Type: field.TypeString},
		{Name: "allowed_http_methods", Type: field.TypeString},
		{Name: "paths", Type: field.TypeJSON},
	}
	// GrantsTable holds the schema information for the "grants" table.
	GrantsTable = &schema.Table{
		Name:        "grants",
		Columns:     GrantsColumns,
		PrimaryKey:  []*schema.Column{GrantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "grant_id",
				Unique:  true,
				Columns: []*schema.Column{GrantsColumns[0]},
			},
			{
				Name:    "grant_created_at",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[1]},
			},
			{
				Name:    "grant_updated_at",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[2]},
			},
			{
				Name:    "grant_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[3]},
			},
			{
				Name:    "grant_hashed_grant_token",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[4]},
			},
			{
				Name:    "grant_domain",
				Unique:  false,
				Columns: []*schema.Column{GrantsColumns[5]},
			},
		},
	}
	// ScopesColumns holds the columns for the "scopes" table.
	ScopesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Size: 255},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "custom_id", Type: field.TypeString},
		{Name: "nonce", Type: field.TypeString},
		{Name: "domain", Type: field.TypeString},
	}
	// ScopesTable holds the schema information for the "scopes" table.
	ScopesTable = &schema.Table{
		Name:        "scopes",
		Columns:     ScopesColumns,
		PrimaryKey:  []*schema.Column{ScopesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "scope_id",
				Unique:  true,
				Columns: []*schema.Column{ScopesColumns[0]},
			},
			{
				Name:    "scope_created_at",
				Unique:  false,
				Columns: []*schema.Column{ScopesColumns[1]},
			},
			{
				Name:    "scope_updated_at",
				Unique:  false,
				Columns: []*schema.Column{ScopesColumns[2]},
			},
			{
				Name:    "scope_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{ScopesColumns[3]},
			},
			{
				Name:    "scope_custom_id",
				Unique:  true,
				Columns: []*schema.Column{ScopesColumns[4]},
			},
			{
				Name:    "scope_domain",
				Unique:  false,
				Columns: []*schema.Column{ScopesColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		APIAuditsTable,
		FactsTable,
		FactTypesTable,
		GrantsTable,
		ScopesTable,
	}
)

func init() {
	FactsTable.ForeignKeys[0].RefTable = FactTypesTable
	FactsTable.ForeignKeys[1].RefTable = ScopesTable
}
