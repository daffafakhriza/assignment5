package service

import (
	"fmt"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryService struct {
	categoryRepository repo.CategoryRepository
}

func NewCategoryService(categoryRepository repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) Store(category *model.Category) error {
	// Validasi input
	if category == nil {
		return fmt.Errorf("kategori tidak boleh nil")
	}

	// Validasi nama kategori
	if category.Name == "" {
		return fmt.Errorf("nama kategori harus diisi")
	}

	// Simpan kategori melalui repository
	err := c.categoryRepository.Store(category)
	if err != nil {
		return fmt.Errorf("gagal menyimpan kategori: %v", err)
	}

	return nil
}

func (c *categoryService) Update(id int, category model.Category) error {
	// Validasi ID
	if id <= 0 {
		return fmt.Errorf("ID kategori tidak valid")
	}

	// Validasi nama kategori
	if category.Name == "" {
		return fmt.Errorf("nama kategori harus diisi")
	}

	// Pastikan ID pada kategori sesuai dengan parameter
	category.ID = id

	// Perbarui kategori melalui repository
	err := c.categoryRepository.Update(id, category)
	if err != nil {
		return fmt.Errorf("gagal memperbarui kategori: %v", err)
	}

	return nil
}

func (c *categoryService) Delete(id int) error {
	// Validasi ID
	if id <= 0 {
		return fmt.Errorf("ID kategori tidak valid")
	}

	// Hapus kategori melalui repository
	err := c.categoryRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("gagal menghapus kategori: %v", err)
	}

	return nil
}

func (c *categoryService) GetByID(id int) (*model.Category, error) {
	// Validasi ID
	if id <= 0 {
		return nil, fmt.Errorf("ID kategori tidak valid")
	}

	// Ambil kategori melalui repository
	category, err := c.categoryRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil kategori: %v", err)
	}

	return category, nil
}

func (c *categoryService) GetList() ([]model.Category, error) {
	// Ambil daftar kategori melalui repository
	categories, err := c.categoryRepository.GetList()
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil daftar kategori: %v", err)
	}

	return categories, nil
}