package models

import "time"

type History struct {
	ID          int       `db:"historyid" json:"id"`
	Date        time.Time `db:"historydate" json:"date" validate:"required"`
	Description string    `db:"historydescription" json:"description" validate:"required"`
	Member      int       `db:"memid" json:"member"`
	Rank        int       `db:"rankid" json:"rank"`
	Role        int       `db:"roleid" json:"role"`
}
