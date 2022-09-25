// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoryColumns holds the columns for the "category" table.
	CategoryColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "slug", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true, Size: 1023},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "full_path", Type: field.TypeString, Nullable: true},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true},
		{Name: "priority", Type: field.TypeInt32, Nullable: true},
		{Name: "post_count", Type: field.TypeUint32, Nullable: true},
	}
	// CategoryTable holds the schema information for the "category" table.
	CategoryTable = &schema.Table{
		Name:       "category",
		Columns:    CategoryColumns,
		PrimaryKey: []*schema.Column{CategoryColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "category_id",
				Unique:  false,
				Columns: []*schema.Column{CategoryColumns[0]},
			},
		},
	}
	// LinkColumns holds the columns for the "link" table.
	LinkColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true, Size: 1023},
		{Name: "logo", Type: field.TypeString, Nullable: true, Size: 1023},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "team", Type: field.TypeString, Nullable: true},
		{Name: "priority", Type: field.TypeInt32, Nullable: true},
	}
	// LinkTable holds the schema information for the "link" table.
	LinkTable = &schema.Table{
		Name:       "link",
		Columns:    LinkColumns,
		PrimaryKey: []*schema.Column{LinkColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "link_id",
				Unique:  false,
				Columns: []*schema.Column{LinkColumns[0]},
			},
		},
	}
	// MenuColumns holds the columns for the "menu" table.
	MenuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "priority", Type: field.TypeInt32, Nullable: true},
		{Name: "target", Type: field.TypeString, Nullable: true},
		{Name: "icon", Type: field.TypeString, Nullable: true},
		{Name: "parent_id", Type: field.TypeUint32, Nullable: true},
		{Name: "team", Type: field.TypeString, Nullable: true},
	}
	// MenuTable holds the schema information for the "menu" table.
	MenuTable = &schema.Table{
		Name:       "menu",
		Columns:    MenuColumns,
		PrimaryKey: []*schema.Column{MenuColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "menu_id",
				Unique:  false,
				Columns: []*schema.Column{MenuColumns[0]},
			},
		},
	}
	// PhotoColumns holds the columns for the "photo" table.
	PhotoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true},
		{Name: "take_time", Type: field.TypeInt64, Nullable: true},
		{Name: "url", Type: field.TypeString, Nullable: true},
		{Name: "team", Type: field.TypeString, Nullable: true},
		{Name: "location", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "likes", Type: field.TypeInt32, Nullable: true},
	}
	// PhotoTable holds the schema information for the "photo" table.
	PhotoTable = &schema.Table{
		Name:       "photo",
		Columns:    PhotoColumns,
		PrimaryKey: []*schema.Column{PhotoColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "photo_id",
				Unique:  false,
				Columns: []*schema.Column{PhotoColumns[0]},
			},
		},
	}
	// PostColumns holds the columns for the "post" table.
	PostColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "slug", Type: field.TypeString, Nullable: true},
		{Name: "meta_keywords", Type: field.TypeString, Nullable: true},
		{Name: "meta_description", Type: field.TypeString, Nullable: true},
		{Name: "full_path", Type: field.TypeString, Nullable: true},
		{Name: "original_content", Type: field.TypeString, Nullable: true},
		{Name: "content", Type: field.TypeString, Nullable: true},
		{Name: "summary", Type: field.TypeString, Nullable: true},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "template", Type: field.TypeString, Nullable: true},
		{Name: "comment_count", Type: field.TypeInt32, Nullable: true},
		{Name: "visits", Type: field.TypeInt32, Nullable: true},
		{Name: "likes", Type: field.TypeInt32, Nullable: true},
		{Name: "word_count", Type: field.TypeInt32, Nullable: true},
		{Name: "top_priority", Type: field.TypeInt32, Nullable: true},
		{Name: "status", Type: field.TypeInt32, Nullable: true},
		{Name: "editor_type", Type: field.TypeInt32, Nullable: true},
		{Name: "edit_time", Type: field.TypeInt64, Nullable: true},
		{Name: "disallow_comment", Type: field.TypeBool, Nullable: true},
		{Name: "in_progress", Type: field.TypeBool, Nullable: true},
	}
	// PostTable holds the schema information for the "post" table.
	PostTable = &schema.Table{
		Name:       "post",
		Columns:    PostColumns,
		PrimaryKey: []*schema.Column{PostColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "post_id",
				Unique:  false,
				Columns: []*schema.Column{PostColumns[0]},
			},
		},
	}
	// TagColumns holds the columns for the "tag" table.
	TagColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true, SchemaType: map[string]string{"mysql": "int", "postgres": "serial"}},
		{Name: "create_time", Type: field.TypeInt64, Nullable: true},
		{Name: "update_time", Type: field.TypeInt64, Nullable: true},
		{Name: "delete_time", Type: field.TypeInt64, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "slug", Type: field.TypeString, Nullable: true},
		{Name: "color", Type: field.TypeString, Nullable: true},
		{Name: "thumbnail", Type: field.TypeString, Nullable: true},
		{Name: "slug_name", Type: field.TypeString, Nullable: true},
		{Name: "post_count", Type: field.TypeUint32, Nullable: true},
	}
	// TagTable holds the schema information for the "tag" table.
	TagTable = &schema.Table{
		Name:       "tag",
		Columns:    TagColumns,
		PrimaryKey: []*schema.Column{TagColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "tag_id",
				Unique:  false,
				Columns: []*schema.Column{TagColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoryTable,
		LinkTable,
		MenuTable,
		PhotoTable,
		PostTable,
		TagTable,
	}
)

func init() {
	CategoryTable.Annotation = &entsql.Annotation{
		Table:     "category",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	LinkTable.Annotation = &entsql.Annotation{
		Table:     "link",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	MenuTable.Annotation = &entsql.Annotation{
		Table:     "menu",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	PhotoTable.Annotation = &entsql.Annotation{
		Table:     "photo",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	PostTable.Annotation = &entsql.Annotation{
		Table:     "post",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
	TagTable.Annotation = &entsql.Annotation{
		Table:     "tag",
		Charset:   "utf8mb4",
		Collation: "utf8mb4_bin",
	}
}