package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	domainsample "template/domain/sample"
	valueobjects "template/domain/value-objects"
	"template/infra/database/postgres"
)

type SampleRepository struct {
	db *sql.DB
}

func NewSampleRepository(db *sql.DB) *SampleRepository {
	return &SampleRepository{db: db}
}

var _ domainsample.SampleRepository = (*SampleRepository)(nil)

func (r *SampleRepository) CreateSample(ctx context.Context, s domainsample.Sample) error {
	exec := postgres.Executor(ctx, r.db)

	_, err := exec.ExecContext(ctx, `
		INSERT INTO sample_table (uuid, name, number, created_at)
		VALUES ($1, $2, $3, $4)
	`,
		s.ID().String(),
		nullStringFromPtr(s.Name()),
		nullInt32FromPtr(s.Number()),
		s.CreatedAt(),
	)
	if err != nil {
		return fmt.Errorf("insert sample: %w", err)
	}

	return nil
}

func (r *SampleRepository) FindByID(
	ctx context.Context,
	id valueobjects.SampleID,
) (*domainsample.Sample, error) {
	exec := postgres.Executor(ctx, r.db)

	var (
		uuid      string
		name      sql.NullString
		number    sql.NullInt32
		createdAt time.Time
	)

	err := exec.QueryRowContext(ctx, `
		SELECT uuid, name, number, created_at
		FROM sample_table
		WHERE uuid = $1
	`, id.String()).Scan(&uuid, &name, &number, &createdAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domainsample.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("select sample: %w", err)
	}

	entity, err := domainsample.NewSample(
		valueobjects.SampleIDFromTrusted(uuid),
		stringPtrFromNull(name),
		intPtrFromNull(number),
		createdAt,
	)
	if err != nil {
		return nil, fmt.Errorf("map sample row: %w", err)
	}

	return &entity, nil
}

func nullStringFromPtr(v *string) sql.NullString {
	if v == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *v, Valid: true}
}

func nullInt32FromPtr(v *int) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: int32(*v), Valid: true}
}

func stringPtrFromNull(v sql.NullString) *string {
	if !v.Valid {
		return nil
	}
	s := v.String
	return &s
}

func intPtrFromNull(v sql.NullInt32) *int {
	if !v.Valid {
		return nil
	}
	n := int(v.Int32)
	return &n
}
