package gen

import (
	"context"
	"database/sql"
	"testing"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/types"
	"github.com/stretchr/testify/require"
)

func TestCreateCandidateAndSkills(t *testing.T) {
	candidateInfos := []types.CandidateInfo{
		{
			Email:           "candidate1@example.com",
			Skills:          []string{"Programming", "Database"},
			ExperienceLevel: 5,
		},
		{
			Email:           "candidate2@example.com",
			Skills:          []string{"Testing", "Debugging"},
			ExperienceLevel: 4,
		},
	}

	// Create candidates and their associated skills
	for _, info := range candidateInfos {
		// Create a candidate
		candidate := CreateCandidateParams{
			Email:           info.Email,
			ExperienceLevel: int32(info.ExperienceLevel),
		}

		candidateResult, err := testQueries.CreateCandidate(context.Background(), candidate)
		require.NoError(t, err)
		require.NotEmpty(t, candidateResult)

		// Create skills for the candidate
		for _, skillName := range info.Skills {
			skill := CreateSkillParams{
				CandidateID: candidateResult.ID,
				SkillName:   skillName,
			}

			err = testQueries.CreateSkill(context.Background(), skill)
			require.NoError(t, err)
		}
	}

	// Retrieve and verify the candidates and their skills
	for _, info := range candidateInfos {
		// Retrieve the candidate by email
		retrievedCandidate, err := testQueries.GetCandidateByEmail(context.Background(), info.Email)
		require.NoError(t, err)
		require.NotEmpty(t, retrievedCandidate)

		// Assertions for retrieved candidate
		require.Equal(t, info.Email, retrievedCandidate.Email)
		require.Equal(t, int32(info.ExperienceLevel), retrievedCandidate.ExperienceLevel)

		// Retrieve the skills for the candidate
		candidateSkills, err := testQueries.ListSkillsByCandidateID(context.Background(), retrievedCandidate.ID)
		require.NoError(t, err)
		require.NotEmpty(t, candidateSkills)

		// Assertions for retrieved skills
		for _, skill := range candidateSkills {
			require.Contains(t, info.Skills, skill.SkillName)
		}
	}
}

func TestDeleteCandidateAndSkills(t *testing.T) {
	candidateInfos := []types.CandidateInfo{
		{
			Email:           "candidate1@example.com",
			Skills:          []string{"Programming", "Database"},
			ExperienceLevel: 5,
		},
		{
			Email:           "candidate2@example.com",
			Skills:          []string{"Testing", "Debugging"},
			ExperienceLevel: 4,
		},
	}

	// Create candidates and their associated skills
	for _, info := range candidateInfos {
		// Create a candidate
		candidate := CreateCandidateParams{
			Email:           info.Email,
			ExperienceLevel: int32(info.ExperienceLevel),
		}

		candidateResult, err := testQueries.CreateCandidate(context.Background(), candidate)
		require.NoError(t, err)
		require.NotEmpty(t, candidateResult)

		// Create skills for the candidate
		for _, skillName := range info.Skills {
			skill := CreateSkillParams{
				CandidateID: candidateResult.ID,
				SkillName:   skillName,
			}

			err = testQueries.CreateSkill(context.Background(), skill)
			require.NoError(t, err)
		}
	}

	// Delete candidates and their associated skills
	for _, info := range candidateInfos {
		// Retrieve the candidate by email
		retrievedCandidate, err := testQueries.GetCandidateByEmail(context.Background(), info.Email)
		require.NoError(t, err)
		require.NotEmpty(t, retrievedCandidate)

		// Delete skills for the candidate
		err = testQueries.DeleteSkillsByCandidateID(context.Background(), retrievedCandidate.ID)
		require.NoError(t, err)

		// Verify that skills have been deleted
		candidateSkills, err := testQueries.ListSkillsByCandidateID(context.Background(), retrievedCandidate.ID)
		require.Error(t, err)
		require.Empty(t, candidateSkills) // The skills slice should be empty

		// Delete the candidate
		err = testQueries.DeleteCandidate(context.Background(), retrievedCandidate.ID)
		require.NoError(t, err)

		// Attempt to retrieve the deleted candidate (should result in an error)
		_, err = testQueries.GetCandidateByEmail(context.Background(), info.Email)
		require.Error(t, err)
		require.EqualError(t, err, sql.ErrNoRows.Error())
	}
}
