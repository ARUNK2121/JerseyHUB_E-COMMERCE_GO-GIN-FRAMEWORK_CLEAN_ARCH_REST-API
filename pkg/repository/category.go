package repository

import (
	"errors"
	"jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
	"jerseyhub/pkg/utils/models"
	"strconv"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{DB}
}

func (p *categoryRepository) AddCategory(c domain.Category) (domain.Category, error) {

	var b string
	err := p.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", c.Category).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}

	var categoryResponse domain.Category
	err = p.DB.Raw(`
	SELECT
		p.id,
		p.category
		FROM
			categories p
		WHERE
			p.category = ?
			`, b).Scan(&categoryResponse).Error

	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil

}

func (p *categoryRepository) CheckCategory(current string) (bool, error) {
	var i int
	err := p.DB.Raw("SELECT COUNT(*) FROM categories WHERE category=?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}

	if i == 0 {
		return false, err
	}

	return true, err
}

func (p *categoryRepository) UpdateCategory(current, new string) (domain.Category, error) {

	// Check the database connection
	if p.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")
	}

	// Update the category
	if err := p.DB.Exec("UPDATE categories SET category = $1 WHERE category = $2", new, current).Error; err != nil {
		return domain.Category{}, err
	}

	// Retrieve the updated category
	var newcat domain.Category
	if err := p.DB.First(&newcat, "category = ?", new).Error; err != nil {
		return domain.Category{}, err
	}

	return newcat, nil
}

func (c *categoryRepository) DeleteCategory(categoryID string) error {
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}

	result := c.DB.Exec("DELETE FROM categories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (c *categoryRepository) GetCategories() ([]domain.Category, error) {
	var model []domain.Category
	err := c.DB.Raw("SELECT * FROM categories").Scan(&model).Error
	if err != nil {
		return []domain.Category{}, err
	}

	return model, nil
}

func (c *categoryRepository) GetBannersForUsers() ([]models.Banner, error) {
	var banners []models.Banner
	err := c.DB.Raw(`select offers.category_id,categories.category as category_name,offers.discount_rate as discount_percentage
	 from offers
	 join categories on categories.id = offers.category_id
	 where offers.discount_rate > 10 
	 Order by offers.discount_rate desc
	 limit 3`).Scan(&banners).Error
	if err != nil {
		return []models.Banner{}, err
	}
	return banners, nil
}

func (c *categoryRepository) GetImagesOfProductsFromACategory(CategoryID int) ([]string, error) {
	var images []string
	err := c.DB.Raw("select image from inventories where category_id = $1 limit 2", CategoryID).Scan(&images).Error
	if err != nil {
		return []string{}, err
	}

	return images, nil
}
