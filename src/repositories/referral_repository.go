package repositories

import (
	"database/sql"
	"time"

	"github.com/dribeiroferr/referral-system-go/src/models"
)

type ReferralRepository interface {
	CreateReferral(referral *models.Referral) error
	FindReferralByHash(hash string) (*models.Referral, error)
	IncrementReferralCount(user string) error
}

type referralRepository struct {
	db *sql.DB
}

func NewReferralRepository(db *sql.DB) ReferralRepository {
	return &referralRepository{db: db}
}

func (r *referralRepository) CreateReferral(referral *models.Referral) error {
	query := "INSERT INTO referrals (user, hash, promo_code, date_generated) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, referral.User, referral.Hash, referral.PromoCode, time.Now())
	return err
}

func (r *referralRepository) FindReferralByHash(hash string) (*models.Referral, error) {
	query := "SELECT id, user, promo_code, count, date_generated, date_registered FROM referrals WHERE hash = ?"
	row := r.db.QueryRow(query, hash)

	var dateGenerated, dateRegistered sql.NullString
	referral := &models.Referral{}
	err := row.Scan(&referral.ID, &referral.User, &referral.PromoCode, &referral.Count, &dateGenerated, &dateRegistered)
	if err != nil {
		return nil, err
	}

	if dateGenerated.Valid && dateGenerated.String != "" {
		referral.DateGenerated, err = time.Parse(time.RFC3339, dateGenerated.String)
		if err != nil {
			return nil, err
		}
	}

	if dateRegistered.Valid && dateRegistered.String != "" {
		referral.DateRegistered, err = time.Parse(time.RFC3339, dateRegistered.String)
		if err != nil {
			return nil, err
		}
	}

	return referral, nil
}

func (r *referralRepository) IncrementReferralCount(user string) error {
	query := "UPDATE referrals SET count = count + 1, date_registered = ? WHERE user = ?"
	_, err := r.db.Exec(query, time.Now(), user)
	return err
}
