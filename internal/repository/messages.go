package repository

import (
	"Messaggio/internal/entity"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MessagesPostgres struct {
	db *sqlx.DB
}

func NewMessagesPostgres(db *sqlx.DB) *MessagesPostgres {
	return &MessagesPostgres{db}
}

func (p *MessagesPostgres) InsertNew(ctx *gin.Context, message entity.Message) (*entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		INSERT INTO messages (id, message)
		VALUES ($1, $2)
		RETURNING *;
	`
	err := p.db.GetContext(ctx.Request.Context(), &dbMessage, query, uuid.New(), message.Message)

	return &dbMessage, err
}

func (p *MessagesPostgres) GetAll(ctx *gin.Context) ([]entity.DBMessage, error) {
	var dbMessages []entity.DBMessage

	query := `
		SELECT * 
		FROM messages 
		ORDER BY created_at DESC;
	`
	err := p.db.SelectContext(ctx.Request.Context(), &dbMessages, query)

	return dbMessages, err
}

func (p *MessagesPostgres) GetById(ctx *gin.Context, id string) (entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		SELECT * 
		FROM messages 
		WHERE id = $1;
	`
	err := p.db.GetContext(ctx.Request.Context(), &dbMessage, query, id)

	return dbMessage, err
}

func (p *MessagesPostgres) DeleteMessage(ctx *gin.Context, id string) (entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		DELETE FROM messages 
		WHERE id = $1 
		RETURNING *;
	`
	err := p.db.GetContext(ctx.Request.Context(), &dbMessage, query, id)

	return dbMessage, err
}

func (p *MessagesPostgres) EditMessage(ctx *gin.Context, id string, message entity.Message) (*entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		UPDATE messages 
		SET message = $1 
		WHERE id = $2
		RETURNING *;
	`
	err := p.db.GetContext(ctx.Request.Context(), &dbMessage, query, message.Message, id)

	return &dbMessage, err
}

func (p *MessagesPostgres) GetStats(ctx *gin.Context) (*entity.Statistic, error) {
	var stats entity.Statistic

	query := `
		SELECT 
		    (SELECT COUNT(*) FROM messages WHERE marked = false) AS unmarked,
			(SELECT COUNT(*) FROM messages WHERE marked = true) AS marked;
	`
	err := p.db.GetContext(ctx.Request.Context(), &stats, query)

	return &stats, err
}

func (p *MessagesPostgres) Mark(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE messages 
		SET marked = true 
		WHERE id = $1
	`
	_, err := p.db.ExecContext(ctx, query, id)

	return err
}
