package models

import "time"

type History struct {
	ID          int       `db:"historyId" json:"id"`
	Date        time.Time `db:"historyDate" json:"date" validate:"required"`
	Description string    `db:"historyDescription" json:"description" validate:"required"`
	Member      int       `db:"memId" json:"member"`
	Rank        int       `db:"rankId" json:"rank"`
	Role        int       `db:"roleId" json:"role"`
}
