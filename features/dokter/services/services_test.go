package services

import (
	"errors"
	"jadwaldokter/config"
	"jadwaldokter/features/dokter/domain"
	"jadwaldokter/features/dokter/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddDokter(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Success Add ", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.DokterCore{ID: uint(1), Nama_dokter: " dr. Lukman, Sp.OG ", Poli_ID: uint(1)}, nil).Once()
		srv := New(repo)
		input := domain.DokterCore{Nama_dokter: " dr. Lukman, Sp.OG ", Poli_ID: 1}
		res, err := srv.AddDokter(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.ID, "harusnya ada id yang terbuat")
		assert.Equal(t, input.Nama_dokter, res.Nama_dokter, "seharusnya nama sama")
		repo.AssertExpectations(t)
	})

	t.Run("Problem", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.DokterCore{}, errors.New("some problem on database")).Once()
		srv := New(repo)
		input := domain.DokterCore{Nama_dokter: " dr. Lukman, Sp.OG ", Poli_ID: 1}
		res, err := srv.AddDokter(input)
		assert.NotNil(t, err)
		assert.Empty(t, res, "karena object gagal dibuat")
		assert.Equal(t, uint(0), res.ID, "id harusnya 0 karena tidak ada data")
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}

func TestDeleteDokter(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses delete vehicle", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.DokterCore{}, nil).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteDokter(id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.DokterCore{}, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteDokter(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("error data on database", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.DokterCore{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteDokter(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

}

func TestGetDokter(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Succses Get ", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.DokterCore{{ID: uint(1), Nama_dokter: " dr. Lukman, Sp.OG ", Poli_ID: uint(1), Nama_poli: "Spesialis Anak / Pediatrician"}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetDokter()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Greater(t, res[0].ID, uint(0)) //lebih besar
		assert.GreaterOrEqual(t, len(res), 1) //lebih besar atau sama
		repo.AssertExpectations(t)
	})

	t.Run("Cant Retrive on database", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		res, err := srv.GetDokter()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		res, err := srv.GetDokter()
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

}
