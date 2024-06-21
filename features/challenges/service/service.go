package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
	"log"
	"math/rand"
)

type ChallengeService struct {
	c challenges.ChallengeDataInterface
}

func New(c challenges.ChallengeDataInterface) challenges.ChallengeServiceInterface {
	return &ChallengeService{
		c: c,
	}
}

func (cs *ChallengeService) GetAllForUser(userId string) ([]challenges.Challenge, error) {
	challenges, err := cs.c.GetAllForUser(userId)
	if err != nil {
		return nil, err
	}
	rand.Shuffle(len(challenges), func(i, j int) {
		challenges[i], challenges[j] = challenges[j], challenges[i]
	})
	return challenges, nil
}

func (cs *ChallengeService) GetChallengeParticipant(challengeId string) (int, error) {
	challenge, err := cs.c.GetChallengeParticipant(challengeId)
	if err != nil {
		return 0, err
	}
	return challenge, nil
}
func (cs *ChallengeService) GetByID(challengeId string) (challenges.Challenge, error) {
	challenge, err := cs.c.GetByID(challengeId)
	log.Println(challenge)
	if err != nil {
		return challenges.Challenge{}, err
	}
	return challenge, nil
}

func (cs *ChallengeService) Swipe(userId string, challengeId string, challengeType string) error {
	// Kode Anda di sini
	if userId == "" || challengeId == "" || challengeType == "" {
		return constant.ErrChallengeFieldSwipe
	}
	if challengeType != "Diterima" && challengeType != "Ditolak" {
		return constant.ErrChallengeType
	}
	err := cs.c.Swipe(userId, challengeId, challengeType)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChallengeService) GetAllForAdmin(page int) ([]challenges.Challenge, int, error) {
	if page <= 0 {
		return nil, 0, constant.ErrPageInvalid
	}
	challenges, total, err := cs.c.GetAllForAdmin(page)
	if err != nil {
		return nil, 0, err
	}
	if page > total {
		return nil, 0, constant.ErrPageInvalid
	}
	return challenges, total, nil
}

func (cs *ChallengeService) Create(challenge challenges.Challenge) error {
	if challenge.Title == "" || challenge.Description == "" || challenge.Exp == 0 || challenge.Coin == 0 || challenge.DateStart.IsZero() || challenge.DateEnd.IsZero() || challenge.Difficulty == "" || challenge.ImpactCategories == nil {
		return constant.ErrChallengeFieldCreate
	}

	err := cs.c.Create(challenge)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChallengeService) GetUserParticipate(userId string, status string) ([]challenges.ChallengeConfirmation, error) {
	challenges, err := cs.c.GetUserParticipate(userId, status)
	if err != nil {
		return nil, err
	}
	return challenges, nil
}
func (cs *ChallengeService) Update(challenge challenges.Challenge) error {
	// Kode Anda di sini
	if challenge.Title == "" || challenge.Description == "" || challenge.Exp == 0 || challenge.Coin == 0 || challenge.DateStart.IsZero() || challenge.DateEnd.IsZero() || challenge.Difficulty == "" || challenge.ImpactCategories == nil {
		return constant.ErrChallengeFieldCreate
	}
	isChallengeExist, err := cs.c.GetByID(challenge.ID)
	if err != nil {
		return err
	}
	if isChallengeExist.ID == "" {
		return constant.ErrChallengeNotFound
	}

	err = cs.c.Update(challenge)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChallengeService) Delete(challengeId string) error {
	isChallengeExist, err := cs.c.GetByID(challengeId)
	if err != nil {
		return err
	}
	if isChallengeExist.ID == "" {
		return constant.ErrChallengeNotFound
	}

	err = cs.c.Delete(challengeId)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChallengeService) GetChallengeForUserByID(challengeConfirmationId string) (challenges.ChallengeConfirmation, error) {
	if challengeConfirmationId == "" {
		return challenges.ChallengeConfirmation{}, constant.ErrGetChallengeByID
	}
	challenge, err := cs.c.GetChallengeForUserByID(challengeConfirmationId)
	if err != nil {
		return challenges.ChallengeConfirmation{}, err
	}
	return challenge, nil
}
func (cs *ChallengeService) EditChallengeForUserByID(challengeId string, images []string) error {
	if challengeId == "" {
		return constant.ErrGetChallengeByID
	}
	if images == nil {
		return constant.ErrChallengeFieldUpdate
	}
	return cs.c.EditChallengeForUserByID(challengeId, images)
}

// Buat mock
func (cs *ChallengeService) InsertCoinAndExpUser(challengeConfirmationId string) error {
	if challengeConfirmationId == "" {
		return constant.ErrGetChallengeByID
	}
	err := cs.c.InsertCoinAndExpUser(challengeConfirmationId)
	if err != nil {
		return err
	}
	return nil
}
