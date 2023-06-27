package repository

import (
	"errors"
	"fmt"
	"jerseyhub/pkg/domain"
	interfaces "jerseyhub/pkg/repository/interface"
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
	fmt.Println("i:", i)
	if err != nil {
		return false, err
	}

	if i == 0 {
		return false, err
	}

	return true, err
}

func (p *categoryRepository) UpdateCategory(current, new string) (domain.Category, error) {
	fmt.Println("values:", new, current)

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

	fmt.Println(newcat)

	return newcat, nil
}

// func (c *categoryRepository) DeleteCategory(categoryID string) error {

// 	id, _ := strconv.Atoi(categoryID)
// 	fmt.Println("heyy this is the id:", id)
// 	result := c.DB.Exec("delete from categories where id = ?", id)

// 	if result.RowsAffected < 1 {
// 		return errors.New("no records were of that id exists")
// 	}

// 	fmt.Println(result.Error)
// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil

// }

func (c *categoryRepository) DeleteCategory(categoryID string) error {
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}
	fmt.Println("This is the ID:", id)

	result := c.DB.Exec("DELETE FROM categories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}
