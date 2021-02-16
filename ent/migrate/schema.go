// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FactsColumns holds the columns for the "facts" table.
	FactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "encrypted_value", Type: field.TypeBytes},
		{Name: "fact_type_facts", Type: field.TypeUUID, Nullable: true},
		{Name: "scope_facts", Type: field.TypeUUID, Nullable: true},
	}
	// FactsTable holds the schema information for the "facts" table.
	FactsTable = &schema.Table{
		Name:       "facts",
		Columns:    FactsColumns,
		PrimaryKey: []*schema.Column{FactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "facts_fact_types_facts",
				Columns: []*schema.Column{FactsColumns[4]},

				RefColumns: []*schema.Column{FactTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "facts_scopes_facts",
				Columns: []*schema.Column{FactsColumns[5]},

				RefColumns: []*schema.Column{ScopesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FactTypesColumns holds the columns for the "fact_types" table.
	FactTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// FactTypesTable holds the schema information for the "fact_types" table.
	FactTypesTable = &schema.Table{
		Name:        "fact_types",
		Columns:     FactTypesColumns,
		PrimaryKey:  []*schema.Column{FactTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ScopesColumns holds the columns for the "scopes" table.
	ScopesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "nonce", Type: field.TypeUUID},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
	}
	// ScopesTable holds the schema information for the "scopes" table.
	ScopesTable = &schema.Table{
		Name:        "scopes",
		Columns:     ScopesColumns,
		PrimaryKey:  []*schema.Column{ScopesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FactsTable,
		FactTypesTable,
		ScopesTable,
	}
)

func init() {
	FactsTable.ForeignKeys[0].RefTable = FactTypesTable
	FactsTable.ForeignKeys[1].RefTable = ScopesTable
}
