// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: query.sql

package gen

import (
	"context"
)

const createCandidate = `-- name: CreateCandidate :one
INSERT INTO candidates (email, experience_level)
VALUES ($1, $2)
RETURNING id, experience_level, created_at, updated_at, email
`

type CreateCandidateParams struct {
	Email           string `json:"email"`
	ExperienceLevel int32  `json:"experience_level"`
}

func (q *Queries) CreateCandidate(ctx context.Context, arg CreateCandidateParams) (Candidate, error) {
	row := q.db.QueryRowContext(ctx, createCandidate, arg.Email, arg.ExperienceLevel)
	var i Candidate
	err := row.Scan(
		&i.ID,
		&i.ExperienceLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}

const createSkill = `-- name: CreateSkill :exec
INSERT INTO skills (candidate_id, skill_name)
VALUES ($1, $2)
RETURNING id
`

type CreateSkillParams struct {
	CandidateID int32  `json:"candidate_id"`
	SkillName   string `json:"skill_name"`
}

func (q *Queries) CreateSkill(ctx context.Context, arg CreateSkillParams) error {
	_, err := q.db.ExecContext(ctx, createSkill, arg.CandidateID, arg.SkillName)
	return err
}

const deleteCandidate = `-- name: DeleteCandidate :exec
DELETE FROM candidates WHERE id = $1
`

func (q *Queries) DeleteCandidate(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCandidate, id)
	return err
}

const deleteSkillsByCandidateID = `-- name: DeleteSkillsByCandidateID :exec
DELETE FROM skills WHERE candidate_id = $1
`

func (q *Queries) DeleteSkillsByCandidateID(ctx context.Context, candidateID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSkillsByCandidateID, candidateID)
	return err
}

const getCandidateByEmail = `-- name: GetCandidateByEmail :one
SELECT id, experience_level, created_at, updated_at, email FROM candidates WHERE email = $1
`

func (q *Queries) GetCandidateByEmail(ctx context.Context, email string) (Candidate, error) {
	row := q.db.QueryRowContext(ctx, getCandidateByEmail, email)
	var i Candidate
	err := row.Scan(
		&i.ID,
		&i.ExperienceLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}

const getCandidateByID = `-- name: GetCandidateByID :one
SELECT id, experience_level, created_at, updated_at, email FROM candidates WHERE id = $1
`

func (q *Queries) GetCandidateByID(ctx context.Context, id int32) (Candidate, error) {
	row := q.db.QueryRowContext(ctx, getCandidateByID, id)
	var i Candidate
	err := row.Scan(
		&i.ID,
		&i.ExperienceLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
	)
	return i, err
}

const getSkillByID = `-- name: GetSkillByID :many
SELECT id, candidate_id, skill_name FROM skills WHERE id = $1
`

func (q *Queries) GetSkillByID(ctx context.Context, id int32) ([]Skill, error) {
	rows, err := q.db.QueryContext(ctx, getSkillByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Skill
	for rows.Next() {
		var i Skill
		if err := rows.Scan(&i.ID, &i.CandidateID, &i.SkillName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllCandidates = `-- name: ListAllCandidates :many
SELECT id, experience_level, created_at, updated_at, email FROM candidates
`

func (q *Queries) ListAllCandidates(ctx context.Context) ([]Candidate, error) {
	rows, err := q.db.QueryContext(ctx, listAllCandidates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Candidate
	for rows.Next() {
		var i Candidate
		if err := rows.Scan(
			&i.ID,
			&i.ExperienceLevel,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Email,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllSkills = `-- name: ListAllSkills :many
SELECT id, candidate_id, skill_name FROM skills
`

func (q *Queries) ListAllSkills(ctx context.Context) ([]Skill, error) {
	rows, err := q.db.QueryContext(ctx, listAllSkills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Skill
	for rows.Next() {
		var i Skill
		if err := rows.Scan(&i.ID, &i.CandidateID, &i.SkillName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSkillsByCandidateID = `-- name: ListSkillsByCandidateID :many
SELECT id, candidate_id, skill_name FROM skills WHERE candidate_id = $1
`

func (q *Queries) ListSkillsByCandidateID(ctx context.Context, candidateID int32) ([]Skill, error) {
	rows, err := q.db.QueryContext(ctx, listSkillsByCandidateID, candidateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Skill
	for rows.Next() {
		var i Skill
		if err := rows.Scan(&i.ID, &i.CandidateID, &i.SkillName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}