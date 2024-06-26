package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dribeiroferr/referral-system-go/src/models"
	"github.com/dribeiroferr/referral-system-go/src/repositories"
	"github.com/dribeiroferr/referral-system-go/src/services"
)

func TestReferralService_GenerateReferralLink(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	repo := repositories.NewReferralRepository(db)
	service := services.NewReferralService(repo)

	link, hash, promoCode, err := service.GenerateReferralLink("test_user")
	require.NoError(t, err)

	assert.NotEmpty(t, link)
	assert.NotEmpty(t, hash)
	assert.NotEmpty(t, promoCode)

	found, err := repo.FindReferralByHash(hash)
	require.NoError(t, err)
	assert.Equal(t, "test_user", found.User)
	assert.Equal(t, promoCode, found.PromoCode)
}

func TestReferralService_HandleReferral(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	repo := repositories.NewReferralRepository(db)
	service := services.NewReferralService(repo)

	referral := &models.Referral{
		User:      "test_user",
		Hash:      "test_hash",
		PromoCode: "test_code",
	}

	err := repo.CreateReferral(referral)
	require.NoError(t, err)

	found, err := service.HandleReferral("test_hash")
	require.NoError(t, err)
	assert.Equal(t, "test_user", found.User)
	assert.Equal(t, "test_code", found.PromoCode)
	assert.Equal(t, 1, found.Count) // A contagem deve ser incrementada
}
