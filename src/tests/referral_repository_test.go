package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dribeiroferr/referral-system-go/src/models"
	"github.com/dribeiroferr/referral-system-go/src/repositories"
)

func TestReferralRepository_CreateReferral(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	repo := repositories.NewReferralRepository(db)

	referral := &models.Referral{
		User:      "test_user",
		Hash:      "test_hash",
		PromoCode: "test_code",
	}

	err := repo.CreateReferral(referral)
	require.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM referrals WHERE user = ?", referral.User).Scan(&count)
	require.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestReferralRepository_FindReferralByHash(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	repo := repositories.NewReferralRepository(db)

	referral := &models.Referral{
		User:      "test_user",
		Hash:      "test_hash",
		PromoCode: "test_code",
	}

	err := repo.CreateReferral(referral)
	require.NoError(t, err)

	found, err := repo.FindReferralByHash("test_hash")
	require.NoError(t, err)
	assert.Equal(t, "test_user", found.User)
	assert.Equal(t, "test_code", found.PromoCode)
}

func TestReferralRepository_IncrementReferralCount(t *testing.T) {
	db := SetupTestDB()
	defer db.Close()

	repo := repositories.NewReferralRepository(db)

	referral := &models.Referral{
		User:      "test_user",
		Hash:      "test_hash",
		PromoCode: "test_code",
	}

	err := repo.CreateReferral(referral)
	require.NoError(t, err)

	err = repo.IncrementReferralCount("test_user")
	require.NoError(t, err)

	found, err := repo.FindReferralByHash("test_hash")
	require.NoError(t, err)
	assert.Equal(t, 1, found.Count)
}
