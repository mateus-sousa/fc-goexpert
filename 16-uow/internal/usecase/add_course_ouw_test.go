package usecase

import (
	"context"
	"database/sql"
	"github.com/mateus-sousa/fc-goexpert/15-uow/internal/db"
	"github.com/mateus-sousa/fc-goexpert/15-uow/internal/repository"
	"github.com/mateus-sousa/fc-goexpert/15-uow/pkg/uow"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCourseUow(t *testing.T) {
	dbt, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	dbt.Exec("DROP TABLE if exists `courses`;")
	dbt.Exec("DROP TABLE if exists `categories`;")

	dbt.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT,name varchar(255) NOT NULL);")
	dbt.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, category_id INTEGER NOT NULL, name varchar(255) NOT NULL, FOREIGN KEY (category_id) REFERENCES  categories(id));")

	ctx := context.Background()
	uowTest := uow.NewUow(ctx, dbt)
	uowTest.Register("CategoryRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCategoryRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})
	uowTest.Register("CourseRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewCourseRepository(dbt)
		repo.Queries = db.New(tx)
		return repo
	})

	input := InputUseCase{CategoryName: "Category 1", CourseName: "Course 1", CourseCategoryID: 2}

	useCase := NewAddCourseUseCaseUow(uowTest)
	err = useCase.Execute(ctx, input)
	assert.NoError(t, err)
}
