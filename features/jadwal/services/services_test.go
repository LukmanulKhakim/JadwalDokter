package services

import (
	"errors"
	"jadwaldokter/config"
	"jadwaldokter/features/jadwal/domain"
	"jadwaldokter/features/jadwal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddJadwal(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Success Add ", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.JadwalCore{ID: uint(1), Hari: "senin", Jam: "09.00-10.00", Dokter_ID: uint(1), Poli_ID: uint(2)}, nil).Once()
		srv := New(repo)
		input := domain.JadwalCore{Hari: "senin", Jam: "09.00-10.00", Dokter_ID: uint(1), Poli_ID: uint(2)}
		res, err := srv.AddJadwal(input)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.ID, "harusnya ada id yang terbuat")
		assert.Equal(t, input.Hari, res.Hari, "seharusnya nama sama")
		repo.AssertExpectations(t)
	})

	t.Run("Problem", func(t *testing.T) {
		repo.On("Add", mock.Anything).Return(domain.JadwalCore{}, errors.New("some problem on database")).Once()
		srv := New(repo)
		input := domain.JadwalCore{Hari: "senin", Jam: "09.00-10.00", Dokter_ID: uint(1), Poli_ID: uint(2)}
		res, err := srv.AddJadwal(input)
		assert.NotNil(t, err)
		assert.Empty(t, res, "karena object gagal dibuat")
		assert.Equal(t, uint(0), res.ID, "id harusnya 0 karena tidak ada data")
		assert.EqualError(t, err, "some problem on database")
		repo.AssertExpectations(t)
	})
}

func TestGetJadwal(t *testing.T) {
	repo := mocks.NewRepository(t)
	var key string
	t.Run("Succses Get ", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return([]domain.JadwalCore{{ID: uint(1), Hari: "senin", Jam: "09.00-10.00", Dokter_ID: uint(1), Poli_ID: uint(2)}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetJadwal(key)
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Greater(t, res[0].ID, uint(0)) //lebih besar
		assert.GreaterOrEqual(t, len(res), 1) //lebih besar atau sama
		repo.AssertExpectations(t)
	})

	t.Run("Cant Retrive on database", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		res, err := srv.GetJadwal(key)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		res, err := srv.GetJadwal(key)
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		assert.Equal(t, len(res), 0, "len harusnya 0 karena tidak ada data")
		repo.AssertExpectations(t)
	})

}

func TestDeleteDokter(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sucses delete vehicle", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.JadwalCore{}, nil).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteJadwal(id)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.JadwalCore{}, gorm.ErrRecordNotFound).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteJadwal(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, gorm.ErrRecordNotFound.Error(), "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

	t.Run("error data on database", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.JadwalCore{}, errors.New(config.DATABASE_ERROR)).Once()
		srv := New(repo)
		var id = uint(1)
		_, err := srv.DeleteJadwal(id)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DATABASE_ERROR, "pesan error tidak sesuai")
		repo.AssertExpectations(t)
	})

}
