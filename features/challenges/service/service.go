package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
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
	// Kode Anda di sini
	challenges, err := cs.c.GetAllForUser(userId)
	if err != nil {
		return nil, err
	}
	rand.Shuffle(len(challenges), func(i, j int) {
		challenges[i], challenges[j] = challenges[j], challenges[i]
	})
	return challenges, nil
}

func (cs *ChallengeService) GetByID(challengeId string) (challenges.Challenge, error) {
	// Kode Anda di sini
	challenge, err := cs.c.GetByID(challengeId)
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
	if challengeType != "accept" && challengeType != "decline" {
		return constant.ErrChallengeType
	}
	err := cs.c.Swipe(userId, challengeId, challengeType)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChallengeService) GetAllForAdmin(page int) ([]challenges.Challenge, int, error) {
	// Kode Anda di sini
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
	// Kode Anda di sini
	if challenge.Title == "" || challenge.Description == "" || challenge.Exp == 0 || challenge.Coin == 0 || challenge.DateStart.IsZero() || challenge.DateEnd.IsZero() || challenge.Difficulty == "" || challenge.ImpactCategories == nil {
		return constant.ErrChallengeFieldCreate
	}

	err := cs.c.Create(challenge)
	if err != nil {
		return err
	}
	return nil
}

func (cs *ChallengeService) Update(challenge challenges.Challenge) error {
	// Kode Anda di sini
	return nil
}

func (cs *ChallengeService) Delete(challengeId string) error {
	// Kode Anda di sini
	return nil
}
