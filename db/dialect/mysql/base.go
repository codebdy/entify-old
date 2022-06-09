package mysql

import (
	"fmt"

	"rxdrag.com/entify/consts"
)

func (*MySQLBuilder) BuildMeSQL() string {
	return "select id, name, loginName from user where loginName = ?"
}

func (*MySQLBuilder) BuildRolesSQL() string {
	povit := fmt.Sprintf(
		"%s_%d_%d_%d",
		consts.PIVOT,
		consts.ROLE_INNER_ID,
		consts.ROLE_USER_RELATION_INNER_ID,
		consts.USER_INNER_ID,
	)
	return fmt.Sprintf("select a.id, a.name  from role a left join %s b on a.id = b.role where b.user = ?", povit)
}

func (*MySQLBuilder) BuildLoginSQL() string {
	return "select password from user where loginName = ?"
}

func (*MySQLBuilder) BuildCreateMetaSQL() string {
	return `CREATE TABLE meta (
		id bigint NOT NULL AUTO_INCREMENT,
		content json DEFAULT NULL,
		publishedAt datetime DEFAULT NULL,
		createdAt datetime DEFAULT NULL,
		updatedAt datetime DEFAULT NULL,
		status varchar(45) DEFAULT NULL,
		PRIMARY KEY (id)
	) ENGINE=InnoDB AUTO_INCREMENT=1507236403010867251 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
	`
}
func (*MySQLBuilder) BuildCreateAbilitySQL() string {
	return `CREATE TABLE ability (
	id bigint NOT NULL AUTO_INCREMENT,
	entityUuid varchar(255) NOT NULL,
	columnUuid varchar(255) DEFAULT NULL,
	can tinyint(1) NOT NULL,
	expression text NOT NULL,
	abilityType varchar(128) NOT NULL,
	roleId bigint NOT NULL,
	PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=4503621102206976 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
`
}
func (*MySQLBuilder) BuildCreateEntityAuthSettingsSQL() string {
	return `CREATE TABLE entity_auth_settings (
	id bigint NOT NULL AUTO_INCREMENT,
	entityUuid varchar(255) NOT NULL,
	expand tinyint(1) NOT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY entityUuid_UNIQUE (entityUuid)
) ENGINE=InnoDB AUTO_INCREMENT=4503616807239680 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
`
}

func (b *MySQLBuilder) BuildTableCheckSQL(name string, database string) string {
	return fmt.Sprintf(
		"SELECT COUNT(*) FROM information_schema.TABLES WHERE table_name ='%s' AND table_schema ='%s'",
		name,
		database,
	)
}

func nullableString(nullable bool) string {
	if nullable {
		return " NULL "
	}
	return " NOT NULL "
}
