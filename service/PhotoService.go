package service

import (
	"errors"
	"fmt"

	"github.com/srohatgi01/tpp-helper/data"
)

type PhotoService struct {
	Width      int
	Height     int
	Paper      string
	Lamination bool
}

func (p *PhotoService) Validate() error {
	// Check for invalid dimensions
	if p.Width <= 0 {
		return errors.New("width must be greater than zero")
	}
	if p.Height <= 0 {
		return errors.New("height must be greater than zero")
	}
	if p.Paper == "" {
		return errors.New("paper type is required")
	}
	return nil
}

func (p *PhotoService) GetArea() float64 {
	return float64(p.Width) * float64(p.Height)
}

func (p *PhotoService) PhotoPrice() (float64, error) {
	paperRates := data.GetPaperPriceMap()

	rate, ok := paperRates[p.Paper]
	if !ok {
		return 0, errors.New("invalid paper type")
	}

	total := p.GetArea() * rate
	return total, nil
}

func (p *PhotoService) GetLaminationPrice() float64 {
	if p.Lamination {
		lamination := data.GetLamination("normal_lamination")
		return p.GetArea() * lamination.Price
	}
	return 0.0
}

func (p *PhotoService) CalculateTotal() (float64, error) {
	if err := p.Validate(); err != nil {
		return 0, err
	}

	photo_price, err := p.PhotoPrice()
	if err != nil {
		return 0, err
	}
	total := photo_price + p.GetLaminationPrice()
	return total, nil
}

func (p *PhotoService) GetBreakdown() [][]interface{} {
	photo_price, err := p.PhotoPrice()
	if err != nil {
		return [][]interface{}{}
	}
	return [][]interface{}{
		{"Size: ", fmt.Sprintf("%d x %d inches", p.Width, p.Height)},
		{"Area: ", fmt.Sprintf("%.2f sq. inches", p.GetArea())},
		{"Paper Type: ", p.Paper},
		{"Photo Price: ", fmt.Sprintf("%.2f", photo_price)},
		{"Lamination Price: ", fmt.Sprintf("%.2f", p.GetLaminationPrice())},
	}
}
