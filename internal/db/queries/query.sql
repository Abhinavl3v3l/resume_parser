-- name: CreateCandidate :one
INSERT INTO candidates (email, experience_level)
VALUES ($1, $2)
RETURNING *;

-- name: CreateSkill :exec
INSERT INTO skills (candidate_id, skill_name)
VALUES ($1, $2)
RETURNING id;

-- name: GetCandidateByEmail :one
SELECT * FROM candidates WHERE email = $1;

-- name: GetCandidateByID :one
SELECT * FROM candidates WHERE id = $1;

-- name: GetSkillByID :many
SELECT * FROM skills WHERE id = $1;

-- name: ListSkillsByCandidateID :many
SELECT * FROM skills WHERE candidate_id = $1;

-- name: DeleteSkillsByCandidateID :exec
DELETE FROM skills WHERE candidate_id = $1;

-- name: DeleteCandidate :exec
DELETE FROM candidates WHERE id = $1;

-- name: ListAllCandidates :many
SELECT * FROM candidates;

-- name: ListAllSkills :many
SELECT * FROM skills;
