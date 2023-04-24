// Code generated by entc, DO NOT EDIT.

// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/roney492/opv/pkg/ent/schema","Package":"github.com/roney492/opv/pkg/ent","Schemas":[{"name":"APIAudit","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":255,"default":true,"default_kind":19,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"plane","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"hashed_grant_token","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"domain","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"http_path","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"http_method","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"sent_http_status","type":{"Type":12,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["id"]},{"fields":["created_at"]},{"fields":["updated_at"]},{"fields":["deleted_at"]},{"fields":["plane"]},{"fields":["hashed_grant_token"]},{"fields":["domain"]},{"fields":["http_path"]},{"fields":["http_method"]},{"fields":["sent_http_status"]}]},{"name":"Fact","config":{"Table":""},"edges":[{"name":"scope","type":"Scope","ref_name":"facts","unique":true,"inverse":true},{"name":"fact_type","type":"FactType","ref_name":"facts","unique":true,"inverse":true}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":255,"default":true,"default_kind":19,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"hashed_value","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"encrypted_value","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"domain","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["id"]},{"fields":["created_at"]},{"fields":["updated_at"]},{"fields":["deleted_at"]},{"fields":["hashed_value"]},{"fields":["domain"]}]},{"name":"FactType","config":{"Table":""},"edges":[{"name":"facts","type":"Fact"}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":255,"default":true,"default_kind":19,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"slug","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"built_in","type":{"Type":1,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"validation","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["id"]},{"fields":["created_at"]},{"fields":["updated_at"]},{"fields":["deleted_at"]},{"unique":true,"fields":["slug"]}]},{"name":"Grant","config":{"Table":""},"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":255,"default":true,"default_kind":19,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"hashed_grant_token","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"domain","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"version","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"allowed_http_methods","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0}},{"name":"paths","type":{"Type":3,"Ident":"[]string","PkgPath":"","Nillable":true,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["id"]},{"fields":["created_at"]},{"fields":["updated_at"]},{"fields":["deleted_at"]},{"fields":["hashed_grant_token"]},{"fields":["domain"]}]},{"name":"Scope","config":{"Table":""},"edges":[{"name":"facts","type":"Fact"}],"fields":[{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0}},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","Nillable":false,"RType":null},"nillable":true,"optional":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0}},{"name":"id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"size":255,"default":true,"default_kind":19,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"custom_id","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"nonce","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"domain","type":{"Type":7,"Ident":"","PkgPath":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["id"]},{"fields":["created_at"]},{"fields":["updated_at"]},{"fields":["deleted_at"]},{"unique":true,"fields":["custom_id"]},{"fields":["domain"]}]}],"Features":["privacy","entql","schema/snapshot"]}`
