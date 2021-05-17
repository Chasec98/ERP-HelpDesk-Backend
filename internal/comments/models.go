package comments

import "database/sql"

const CommentCtxKey = "comments"

type Comment struct {
	ID          *int
	Text        *string
	CreatedByID *int
	TicketID    *int
}

func (t Comment) ToSQL() CommentSQL {
	var ret CommentSQL
	if t.ID != nil {
		ret.ID = sql.NullInt64{Valid: true, Int64: int64(*t.ID)}
	}
	if t.Text != nil {
		ret.Text = sql.NullString{Valid: true, String: *t.Text}
	}
	if t.CreatedByID != nil {
		ret.CreatedByID = sql.NullInt64{Valid: true, Int64: int64(*t.CreatedByID)}
	}
	if t.TicketID != nil {
		ret.TicketID = sql.NullInt64{Valid: true, Int64: int64(*t.TicketID)}
	}
	return ret
}

type CommentSQL struct {
	ID          sql.NullInt64
	Text        sql.NullString
	CreatedByID sql.NullInt64
	TicketID    sql.NullInt64
}

func (t CommentSQL) ToComment() Comment {
	var ret Comment
	if t.ID.Valid {
		var id = int(t.ID.Int64)
		ret.ID = &id
	}
	if t.Text.Valid {
		ret.Text = &t.Text.String
	}
	if t.CreatedByID.Valid {
		var id = int(t.CreatedByID.Int64)
		ret.CreatedByID = &id
	}
	if t.TicketID.Valid {
		var id = int(t.TicketID.Int64)
		ret.TicketID = &id
	}
	return ret
}
