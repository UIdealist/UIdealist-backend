package models

type Rank struct {
	ID    int    `db:"rankid" json:"id"`
	Name  string `db:"rankname" json:"name" validate:"required"`
	Color int    `db:"rankcolor" json:"color" validate:"required"`
}

type RankAndMembers struct {
	RankID     int    `db:"rankid" json:"rank_id"`
	RankName   string `db:"rankname" json:"rank_name"`
	RankColor  int    `db:"rankcolor" json:"rank_color"`
	MemberID   int    `db:"memid" json:"member_id"`
	MemberName string `db:"memname" json:"member_name"`
}

type RankMembersAndRoles struct {
	RankAndMembers
	MemberRoleID int    `db:"memroleid" json:"member_role_id"`
	RoleID       int    `db:"roleid" json:"role_id"`
	RoleName     string `db:"rolename" json:"role_name"`
}
