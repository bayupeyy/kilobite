package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)

	//Digunakan untuk cek apakah email sudah terdaftar atau belum
	IsEmailAvailable(input CheckEmailInput) (bool, error)

	//Digunakan untuk Save Avatar
	SaveAvatar(ID int, fileLocation string) (User, error)

	//Digunakan untuk mendapatkan user id
	GetUserByID(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash) //Untuk Hashing password
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// Membuat Fungsi untuk kontrak LoginInput
func (s *service) Login(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	//Menggunakan fungsi FindByEmail untuk mencari user berdasarkan email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User tidak ditemukan")

	}

	//Untuk menemukan Password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	//Melakukan pengecekan jika ditemukan error
	if err != nil {
		return user, err
	}
	return user, nil
}

// Membuat Fungsi dari kontrak IsEmailAvailable
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	//Melakukan pengecekan apakah ada error atau tidak
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

// Membuat Fungsi untuk Save Avatar
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//Mendapatkan user berdasarkan ID
	//Update attribute avatar name
	//Simpan perubahan avatar file name

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	//Update atribut avatar
	user.AvatarFileName = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	//Pengecekan error
	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return user, errors.New("No User found on with that ID")
	}

	return user, nil
}
