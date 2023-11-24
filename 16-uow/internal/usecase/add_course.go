package usecase

import (
	"context"
	"github.com/mateus-sousa/fc-goexpert/15-uow/internal/entity"
	"github.com/mateus-sousa/fc-goexpert/15-uow/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewAddCourseUseCase(courseRepository repository.CourseRepositoryInterface, categoryRepository repository.CategoryRepositoryInterface) *AddCourseUseCase {
	return &AddCourseUseCase{CourseRepository: courseRepository, CategoryRepository: categoryRepository}
}

func (a *AddCourseUseCase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{Name: input.CategoryName}
	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}
	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}
	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}
	return nil
}
