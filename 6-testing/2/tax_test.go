package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, tax, 10.0)

	tax, err = CalculateTax(-1.5)
	assert.Error(t, err, "value should be greater than 0")
	assert.Equal(t, tax, 0.0)
}

func TestTaxCalculateAndSave(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	// define: na chamada do metodo tal, passando tal parametro, qual retorno acontecera.
	repositoryMock.On("SaveTax", 10.0).Return(nil)
	repositoryMock.On("SaveTax", 0.0).Return(errors.New("error on save tax"))
	err := TaxCalculateAndSave(1000, repositoryMock)
	assert.Nil(t, err)
	err = TaxCalculateAndSave(0.0, repositoryMock)
	assert.Error(t, err, "error on save tax")

	// Verifica se o metodo do mock realmente foi chamado em todas as chamadas da unidade em teste
	repositoryMock.AssertExpectations(t)
	// Verifica se o metodo do mock foi chamado uma quantidade X de vezes
	repositoryMock.AssertNumberOfCalls(t, "SaveTax", 2)
}
