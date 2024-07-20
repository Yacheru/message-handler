package repository

import (
	"Messaggio/internal/entity"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db}
}

func (p *Postgres) InsertNew(c *gin.Context, message entity.Message) (*entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		INSERT INTO messages (id, message)
		VALUES ($1, $2)
		RETURNING *;
	`
	err := p.db.GetContext(c, &dbMessage, query, uuid.New(), message.Message)

	return &dbMessage, err
}

func (p *Postgres) GetAll(c *gin.Context) ([]entity.DBMessage, error) {
	var dbMessages []entity.DBMessage

	query := `
		SELECT * 
		FROM messages 
		ORDER BY created_at DESC;
	`
	err := p.db.SelectContext(c, &dbMessages, query)

	return dbMessages, err
}

func (p *Postgres) GetById(c *gin.Context, id string) (entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		SELECT * 
		FROM messages 
		WHERE id = $1;
	`
	err := p.db.GetContext(c, &dbMessage, query, id)

	return dbMessage, err
}

func (p *Postgres) DeleteMessage(c *gin.Context, id string) (entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		DELETE FROM messages 
		WHERE id = $1 
		RETURNING *;
	`
	err := p.db.GetContext(c, &dbMessage, query, id)

	return dbMessage, err
}

func (p *Postgres) EditMessage(c *gin.Context, id string, message entity.Message) (*entity.DBMessage, error) {
	var dbMessage entity.DBMessage

	query := `
		UPDATE messages 
		SET message = $1 
		WHERE id = $2
		RETURNING *;
	`
	err := p.db.GetContext(c, &dbMessage, query, message.Message, id)
	if err != nil {
		return nil, err
	}

	return &dbMessage, nil
}

func (p *Postgres) GetStats(c *gin.Context) (*entity.Statistic, error) {
	var stats entity.Statistic

	query := `
		SELECT 
		    (SELECT COUNT(*) FROM messages WHERE marked = false) AS unmarked,
			(SELECT COUNT(*) FROM messages WHERE marked = true) AS marked;
	`
	err := p.db.GetContext(c, &stats, query)

	return &stats, err
}

func (p *Postgres) Mark(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE messages 
		SET marked = true 
		WHERE id = $1
	`
	_, err := p.db.ExecContext(ctx, query, id)

	return err
}
