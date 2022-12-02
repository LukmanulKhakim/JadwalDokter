package services

import (
	"errors"
	"jadwaldokter/config"
	"jadwaldokter/features/poli/domain"
	"jadwaldokter/features/poli/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddPoli(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Success Add ", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.PoliCore{ID: uint(1), Nama_poli: "Spesialis Anak / Pediatrician "}, nil).Once()
		srv := New(repo)
		input := domain.PoliCore{Nama_poli: "Spesialis Anak / Pediatrician "}
		res, err := srv.AddPoli(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.ID, "harusnya ada id yang terbuat")
		assert.Equal(t, input.Nama_poli, res.Nama_poli, "seharusnya nama sama")
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate data", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.PoliCore{}, errors.New("there's duplicate data")).Once()
		srv := New(repo)
		input := domain.PoliCore{Nama_poli: "Spesialis Anak / Pediatrician "}
		res, err := srv.AddPoli(input)
		assert.NotNil(t, err)
		assert.Empty(t, res, "karena object gagal dibuat")
		assert.Equal(t, uint(0), res.ID, "id harusnya 0 karena tidak ada data")
		assert.EqualError(t, err, "rejected from database")
		repo.AssertExpectations(t)
	})

	t.Run("Problem", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.PoliCore{}, errors.New("cannot connect")).Once()
		srv := New(repo)
		input := domain.PoliCore{Nama_poli: "Spesialis Anak / Pediatrician "}
		res, err := srv.AddPoli(input)
		assert.NotNil(t, err)
		assert.Empty(t, res, "karena object gagal dibuat")
		assert.Equal(t, uint(0), res.ID, "id harusnya 0 karena tidak ada data")
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}

func TestDeletePoli(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses delete poli", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.PoliCore{}, nil).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeletePoli(id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.PoliCore{}, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeletePoli(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("error data on database", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.PoliCore{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeletePoli(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

}

func TestGetPoli(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Succses Get ", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.PoliCore{{ID: uint(1), Nama_poli: "Spesialis Anak / Pediatrician "}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetAllPoli()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Greater(t, res[0].ID, uint(0)) //lebih besar
		assert.GreaterOrEqual(t, len(res), 1) //lebih besar atau sama
		repo.AssertExpectations(t)
	})

	t.Run("Cant Retrive on database", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		res, err := srv.GetAllPoli()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		res, err := srv.GetAllPoli()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

}

func TestUpdatePoli(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update User", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.PoliCore{ID: uint(1), Nama_poli: "Spesialis Anak / Pediatrician "}, nil).Once()
		srv := New(repo)
		input := domain.PoliCore{Nama_poli: "Spesialis Anak / Pediatrician "}
		res, err := srv.UpdatePoli(input, uint(1))
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Update Profile", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(domain.PoliCore{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		input := domain.PoliCore{Nama_poli: "Spesialis Anak / Pediatrician "}
		res, err := srv.UpdatePoli(input, uint(1))
		assert.Empty(t, res)
		assert.Error(t, err)
		assert.EqualError(t, err, config.DATABASE_ERROR)
		repo.AssertExpectations(t)
	})

}
