package services

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/dribeiroferr/referral-system-go/src/models"
	"github.com/dribeiroferr/referral-system-go/src/repositories"
)

type ReferralService interface {
	GenerateReferralLink(user string) (string, string, string, error)
	HandleReferral(hash string)(*models.Referral, error)
}

type referralService struct { 
	repo repositories.ReferralRepository
}

func NewReferralService(repo repositories.ReferralRepository) ReferralService {
	return &referralService{repo}
}

func (s *referralService) GenerateReferralLink(user string)(string, string, string,error) {
	hash := generateHash(user)
	promoCode := generatePromoCode()

	referral := &models.Referral{
		User: user,
		Hash: hash,
		PromoCode: promoCode,
	}

	err := s.repo.CreateReferral(referral)
	
	if err != nil {
		return "", "", "", err
	}

	return fmt.Sprintf("http://localhost:9001/referral/%s", hash), hash, promoCode, nil
}


func (s *referralService) HandleReferral(hash string) (*models.Referral, error){
	referral, err := s.repo.FindReferralByHash(hash)
	
	if err !=nil { 
		return nil, err
	}

	err = s.repo.IncrementReferralCount(referral.User)
	
	if err != nil {
		return nil, err
	}

	return referral, nil
}

func generateHash(user string) string { 
	hash := sha256.Sum256([]byte(user + time.Now().String()))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func generatePromoCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource((time.Now().UnixNano())))
	promo := make([] byte, 10)
	
	for i := range promo {
		promo[i] = charset[seededRand.Intn(len(charset))]
	}
	
	return string(promo)
}