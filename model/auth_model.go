package model

import "gorm.io/gorm"

type AuthUser struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthRepository interface {
	Create(user *AuthUser) error
	Update(user *AuthUser) error
	Delete(user *AuthUser) error
	FindByID(user *AuthUser) (*AuthUser, error)
	FindByEmail(user *AuthUser) (*AuthUser, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Create(u *AuthUser) error {
	return r.db.Create(&u).Error
}

func (r *authRepository) Update(u *AuthUser) error {
	return r.db.Update("email, password", &u).Error
}

func (r *authRepository) Delete(u *AuthUser) error {
	return r.db.Where(&AuthUser{ID: u.ID}).Delete(&u).Error
}

func (r *authRepository) FindByID(u *AuthUser) (*AuthUser, error) {
	var user AuthUser
	if err := r.db.Where(&AuthUser{ID: u.ID}).First(&user); err != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (r *authRepository) FindByEmail(u *AuthUser) (*AuthUser, error) {
	var user AuthUser
	if err := r.db.Where(&AuthUser{Email: u.Email}).First(&user); err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}
