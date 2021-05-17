package comments

import (
	"context"
	"database/sql"

	"github.com/Chasec98/ERP-HelpDesk-Backend/pkg/logger"
)

type Repository interface {
	CreateComment(ctx context.Context) (Comment, error)
	UpdateComment(ctx context.Context) (Comment, error)
	GetCommentByID(ctx context.Context) (Comment, error)
	GetCommentsByTicketID(ctx context.Context) ([]Comment, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateComment(ctx context.Context) (Comment, error) {
	commentContext := ctx.Value(CommentCtxKey).(Comment)

	query := `INSERT INTO Comments (Text, TicketID, CreatedByID) VALUES (?,?,?)`
	var commentSQL = commentContext.ToSQL()
	result, err := r.db.ExecContext(ctx, query, commentSQL.Text, commentSQL.TicketID, commentSQL.CreatedByID)
	if err != nil {
		logger.Error.Println(err.Error())
		return Comment{}, err
	}

	commentID, err := result.LastInsertId()
	if err != nil {
		logger.Error.Println(err.Error())
		return Comment{}, err
	}

	var id = int(commentID)
	commentContext.ID = &id
	ctx = context.WithValue(ctx, CommentCtxKey, commentContext)
	return r.GetCommentByID(ctx)
}

func (r repository) UpdateComment(ctx context.Context) (Comment, error) {
	commentContext := ctx.Value(CommentCtxKey).(Comment)

	query := `UPDATE Comments SET Text = ? WHERE ID = ?`
	commentSQL := commentContext.ToSQL()
	_, err := r.db.ExecContext(ctx, query, commentSQL.Text, commentSQL.ID)
	if err != nil {
		logger.Error.Println(err.Error())
		return Comment{}, err
	}

	return r.GetCommentByID(ctx)
}

func (r repository) GetCommentByID(ctx context.Context) (Comment, error) {
	commentContext := ctx.Value(CommentCtxKey).(Comment)

	query := `SELECT ID, Text, CreatedByID, TicketID FROM Comments WHERE ID = ?`
	row := r.db.QueryRowContext(ctx, query, commentContext.ID)
	var commentSQL = CommentSQL{}
	err := row.Scan(&commentSQL.ID, &commentSQL.Text, &commentSQL.CreatedByID, &commentSQL.TicketID)
	if err != nil {
		logger.Error.Println(err.Error())
		return Comment{}, err
	}

	return commentSQL.ToComment(), nil
}

func (r repository) GetCommentsByTicketID(ctx context.Context) ([]Comment, error) {
	commentContext := ctx.Value(CommentCtxKey).(Comment)

	query := `SELECT ID, Text, CreatedByID, TicketID FROM Comments WHERE TicketID = ?`
	var commentSQL = commentContext.ToSQL()
	rows, err := r.db.QueryContext(ctx, query, commentSQL.TicketID)
	if err != nil {
		logger.Error.Println(err.Error())
		return []Comment{}, err
	}

	var comments []Comment
	for rows.Next() {
		var csql CommentSQL
		err := rows.Scan(&csql.ID, &csql.Text, &csql.CreatedByID, &csql.TicketID)
		if err != nil {
			logger.Error.Println(err.Error())
			return []Comment{}, err
		}
		comments = append(comments, csql.ToComment())
	}

	return comments, nil
}
