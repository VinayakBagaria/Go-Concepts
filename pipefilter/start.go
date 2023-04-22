package pipefilter

import (
	"context"
	"errors"
	"fmt"
)

type Product struct {
	Name         string
	Price        float64
	Rating       int
	Discount     float64 // percentage
	NetPrice     float64 // price - discount
	Availability bool
}

const (
	RatingFilterKey       = "filter.rating"
	AvailabilityFilterKey = "filter.availability"
)

type Filter func(context.Context, []Product) ([]Product, error)

type Pipeline struct {
	Filters []Filter
}

func (p *Pipeline) Use(filter Filter) {
	p.Filters = append(p.Filters, filter)
}

func (p *Pipeline) Execute(ctx context.Context, input []Product) ([]Product, error) {
	output := input
	var err error

	for _, filter := range p.Filters {
		output, err = filter(ctx, output)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

var RatingFilter Filter = func(ctx context.Context, products []Product) ([]Product, error) {
	rating := ctx.Value(RatingFilterKey).(int)
	if rating == 0 {
		return []Product{}, errors.New("invalid rating")
	}

	var filteredProducts []Product
	for _, product := range products {
		if product.Rating >= rating {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

var AvailabilityFilter Filter = func(ctx context.Context, products []Product) ([]Product, error) {
	availability := ctx.Value(AvailabilityFilterKey).(bool)
	var filteredProducts []Product
	for _, product := range products {
		if product.Availability == availability {
			filteredProducts = append(filteredProducts, product)
		}
	}
	return filteredProducts, nil
}

func DoWork() {
	productList := []Product{
		{Name: "cup", Price: 10.24, Rating: 3, Discount: 0, Availability: true},
		{Name: "glass", Price: 12.4, Rating: 5, Discount: 10, Availability: false},
		{Name: "pot", Price: 45.98, Rating: 4, Discount: 20, Availability: true},
		{Name: "dish", Price: 5.5, Rating: 4, Discount: 5, Availability: true},
		{Name: "plate", Price: 25.65, Rating: 2, Discount: 25, Availability: false},
	}

	pipeline := Pipeline{}
	pipeline.Use(RatingFilter)
	pipeline.Use(AvailabilityFilter)

	ctx := context.Background()
	ctx = context.WithValue(ctx, RatingFilterKey, 4)
	ctx = context.WithValue(ctx, AvailabilityFilterKey, false)

	products, err := pipeline.Execute(ctx, productList)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("\nproducts: %#v\n", products)
}
