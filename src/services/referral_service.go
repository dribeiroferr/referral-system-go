package services

import (
	"errors"
	"time"

	"github.com/dribeiroferr/referral-system-go/src/models"
	"github.com/dribeiroferr/referral-system-go/src/repositories"
)

type ReferralService interface {
	GenerateReferralLink(user string) (string, string, string, error)
	HandleReferral(hash string) (*models.Referral, error)
}

type referralService struct {
	repo repositories.ReferralRepository
}

func NewReferralService(repo repositories.ReferralRepository) ReferralService {
	return &referralService{repo}
}

func (s *referralService) GenerateReferralLink(user string) (string, string, string, error) {
	hash := generateHash(user)
	promoCode := generatePromoCode()

	referral := &models.Referral{
		User:      user,
		Hash:      hash,
		PromoCode: promoCode,
	}

	err := s.repo.CreateReferral(referral)
	if err != nil {
		return "", "", "", err
	}

	return "http://localhost:8080/referral/" + hash, hash, promoCode, nil
}

func (s *referralService) HandleReferral(hash string) (*models.Referral, error) {
	referral, err := s.repo.FindReferralByHash(hash)
	if err != nil {
		return nil, err
	}

	if referral.User == "" {
		return nil, errors.New("hash inválido")
	}

	err = s.repo.IncrementReferralCount(referral.User)
	if err != nil {
		return nil, err
	}

	// Recarregar a referência para obter a contagem atualizada
	referral, err = s.repo.FindReferralByHash(hash)
	if err != nil {
		return nil, err
	}

	return referral, nil
}

func generateHash(user string) string {
	// Geração de hash simplificada para exemplo
	return user + time.Now().String()
}

func generatePromoCode() string {
	// Geração de código promocional simplificada para exemplo
	return "PROMO" + time.Now().Format("20060102150405")
}
