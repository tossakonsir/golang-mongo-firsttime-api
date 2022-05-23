package services

import (
	"context"
	"errors"
	"example/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductServiceImpl struct {
	productcollection *mongo.Collection
	ctx               context.Context
}

func NewProductService(productcollection *mongo.Collection, ctx context.Context) ProductService {
	return &ProductServiceImpl{
		productcollection: productcollection,
		ctx:               ctx,
	}
}

func (p *ProductServiceImpl) CreateProduct(product *models.Product) error {
	_, err := p.productcollection.InsertOne(p.ctx, product)
	return err

}

func (p *ProductServiceImpl) GetAll() ([]*models.Product, error) {
	var products []*models.Product
	cursor, err := p.productcollection.Find(p.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(p.ctx)

	if len(products) == 0 {
		return nil, errors.New("documents not found")
	}
	return products, nil
}

func (p *ProductServiceImpl) UpdateProduct(product *models.Product) error {
	filter := bson.D{primitive.E{Key: "itemname", Value: product.Itemname}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "description", Value: product.Description}}}}
	result, _ := p.productcollection.UpdateOne(p.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (p *ProductServiceImpl) DeleteProduct(itemname *string) error {
	filter := bson.D{primitive.E{Key: "itemname", Value: itemname}}
	result, _ := p.productcollection.DeleteOne(p.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}
