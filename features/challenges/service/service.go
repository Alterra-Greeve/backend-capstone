package service

import "backendgreeve/features/challenges"

type ChallengeService struct {
	c challenges.ChallengeDataInterface
}

func New(c challenges.ChallengeDataInterface) challenges.ChallengeServiceInterface {
	return &ChallengeService{
		c: c,
	}
}

func (cs *ChallengeService) GetAllForUser() ([]challenges.Challenge, int, error) {
	// Kode Anda di sini
	return nil, 0, nil
}

func (cs *ChallengeService) GetByID(challengeId string) (challenges.Challenge, error) {
	// Kode Anda di sini
	return challenges.Challenge{}, nil
}

func (cs *ChallengeService) SwipeChallenge(userId string, challengeId string) (challenges.Challenge, error) {
	// Kode Anda di sini
	return challenges.Challenge{}, nil
}

func (cs *ChallengeService) GetAllForAdmin(page int) ([]challenges.Challenge, int, error) {
	// Kode Anda di sini
	return cs.c.GetAllForAdmin(page)
}

func (cs *ChallengeService) Create(challenge challenges.Challenge) error {
	// Kode Anda di sini
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
