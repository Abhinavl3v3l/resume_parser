package ops

import (
	"context"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/db"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/db/gen"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/types"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/utils/logger"
)

// AddCandidates adds a new candidate to the database.
func AddCandidates(candidateInfo types.CandidateInfo) error {
	logger.Info("Initiating AddCandidates operation")

	dbInstance, err := db.GetDBInstance()
	if err != nil {
		logger.Error("Failed to get database instance", err)
		return err
	}

	candidate, err := dbInstance.Queries.CreateCandidate(context.Background(), gen.CreateCandidateParams{
		Email:           candidateInfo.Email,
		ExperienceLevel: int32(candidateInfo.ExperienceLevel),
	})

	if err != nil {
		logger.Error("Error adding candidate", err)
		return err
	}

	logger.Info("Successfully added candidate")

	if err := addSkills(candidateInfo, candidate, dbInstance); err != nil {
		logger.Error("Failed to add skills for candidate", err)
		return err
	}

	return nil
}

// addSkills is an internal function to add skills for a given candidate.
func addSkills(candidateInfo types.CandidateInfo, candidate gen.Candidate, dbInstance *db.Database) error {
	logger.Info("Initiating AddSkills operation")

	for _, skill := range candidateInfo.Skills {
		err := dbInstance.Queries.CreateSkill(context.Background(), gen.CreateSkillParams{
			CandidateID: candidate.ID,
			SkillName:   skill,
		})

		if err != nil {
			logger.Error("Error adding skill", skill, err)
			return err
		}
	}

	logger.Info("Successfully added skills for candidate", candidate.ID)
	return nil
}

// GetCandidate fetches a candidate by ID.
func GetCandidate(candidateID int32) (*types.CandidateInfo, error) {
	dbInstance, err := db.GetDBInstance()
	if err != nil {
		logger.Error("Failed to get database instance", err)
		return nil, err
	}

	candidate, err := dbInstance.Queries.GetCandidateByID(context.Background(), candidateID)
	if err != nil {
		logger.Error("Failed to fetch candidate", candidateID, err)
		return nil, err
	}

	return &types.CandidateInfo{
		Email:           candidate.Email,
		ExperienceLevel: int(candidate.ExperienceLevel),
		// Add other fields as necessary
	}, nil
}

// GetSkillsForCandidate fetches all skills for a given candidate ID.
func GetSkillsForCandidate(candidateID int32) ([]string, error) {
	dbInstance, err := db.GetDBInstance()
	if err != nil {
		logger.Error("Failed to get database instance", err)
		return nil, err
	}

	skills, err := dbInstance.Queries.ListSkillsByCandidateID(context.Background(), candidateID)
	if err != nil {
		logger.Error("Failed to fetch skills for candidate", candidateID, err)
		return nil, err
	}

	skillNames := make([]string, len(skills))
	for i, skill := range skills {
		skillNames[i] = skill.SkillName
	}

	return skillNames, nil
}

func GetCandidateByEmail(email string) (*gen.Candidate, error) {
	dbInstance, err := db.GetDBInstance()
	if err != nil {
		logger.Error("Failed to get database instance", err)
		return nil, err
	}

	candidate, err := dbInstance.Queries.GetCandidateByEmail(context.Background(), email)
	if err != nil {
		logger.Error("Failed to fetch candidate by email", email, err)
		return nil, err
	}

	return &candidate, nil
}

// DeleteCandidate deletes a candidate and their associated skills.
func DeleteCandidate(candidateID int32) error {
	dbInstance, err := db.GetDBInstance()
	if err != nil {
		logger.Error("Failed to get database instance", err)
		return err
	}

	// Delete candidate's skills first to maintain FK constraints
	err = dbInstance.Queries.DeleteSkillsByCandidateID(context.Background(), candidateID)
	if err != nil {
		logger.Error("Failed to delete skills for candidate", candidateID, err)
		return err
	}

	// Delete the candidate
	err = dbInstance.Queries.DeleteCandidate(context.Background(), candidateID)
	if err != nil {
		logger.Error("Failed to delete candidate", candidateID, err)
		return err
	}

	logger.Info("Successfully deleted candidate and their skills", candidateID)
	return nil
}
