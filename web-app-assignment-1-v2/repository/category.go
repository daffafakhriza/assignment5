package repository

import (
	"fmt"
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	// Validasi input
	if Category == nil {
		return fmt.Errorf("kategori tidak boleh nil")
	}

	if Category.ID == 0 {
		return fmt.Errorf("ID kategori harus ditentukan")
	}

	// Simpan kategori menggunakan metode dari filebased
	err := c.filebasedDb.StoreCategory(*Category)
	if err != nil {
		return fmt.Errorf("gagal menyimpan kategori: %v", err)
	}

	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	// Validasi input
	if id == 0 {
		return fmt.Errorf("ID kategori tidak valid")
	}

	// Pastikan ID pada parameter sama dengan ID kategori
	category.ID = id

	// Gunakan metode UpdateCategory dari filebased
	err := c.filebasedDb.UpdateCategory(id, category)
	if err != nil {
		return fmt.Errorf("gagal memperbarui kategori: %v", err)
	}

	return nil
}

func (c *categoryRepository) Delete(id int) error {
	// Validasi input
	if id == 0 {
		return fmt.Errorf("ID kategori tidak valid")
	}

	// Cek apakah kategori ada sebelum menghapus
	_, err := c.GetByID(id)
	if err != nil {
		return fmt.Errorf("kategori tidak ditemukan: %v", err)
	}

	// Hapus kategori menggunakan metode dari filebased
	err = c.filebasedDb.DeleteCategory(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus kategori: %v", err)
	}

	return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	// Validasi input
	if id == 0 {
		return nil, fmt.Errorf("ID kategori tidak valid")
	}

	// Ambil kategori menggunakan metode dari filebased
	category, err := c.filebasedDb.GetCategoryByID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil kategori: %v", err)
	}

	return category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	// Ambil daftar kategori menggunakan metode dari filebased
	categories, err := c.filebasedDb.GetCategories()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar kategori: %v", err)
	}

	// Jika tidak ada kategori, kembalikan slice kosong
	if len(categories) == 0 {
		return []model.Category{}, nil
	}

	return categories, nil
}