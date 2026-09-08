package postgres

import (
	"context"
	"database/sql"
	"time"

	"template/infra/runtime"
)

type Indicator struct {
	db *sql.DB
}

func NewIndicator(db *sql.DB) *Indicator {
	return &Indicator{db: db}
}

func (p *Indicator) Name() string {
	return "postgres"
}

func (p *Indicator) Ready(ctx context.Context) bool {
	if p.db == nil {
		return false
	}

	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return p.db.PingContext(pingCtx) == nil
}

var _ runtime.Indicator = (*Indicator)(nil)
