package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("uuid").Comment(" User's UUID | 用户的UUID"),
		field.String("token").SchemaType(map[string]string{
			dialect.MySQL: "varchar(512)",
		}).Comment("Token string | Token 字符串"),
		field.String("source").Comment("Log in source such as GitHub | Token 来源 （本地为core, 第三方如github等）"),
		field.Time("expired_at").Comment(" Expire time | 过期时间"),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.StatusMixin{},
	}
}

func (Token) Edges() []ent.Edge {
	return nil
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("uuid"),
		index.Fields("expired_at"),
	}
}

func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_tokens"},
	}
}
