package models

import "e-commerce/initializers"

type Product struct {
	ID           uint   `gorm:"primaryKey" json:"productId"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	CountInStock int    `json:"countInStck"`
	Description  string `json:"description"`
}

func (p *Product) Create(prod *Product) error {
    if result := initializers.DB.Model(p).Create(prod); result.Error != nil{
        return result.Error
    }
    return nil 	
}

func (p *Product) Delete(id int) error{
    if result := initializers.DB.Delete(p, id); result.Error != nil{
        return result.Error
    }
    return nil 
}

func (p *Product) Update(id int, newProd *Product) error {
    oldProd := &Product{}
    if result := initializers.DB.Model(p).First(oldProd); result.Error != nil{
        return result.Error
    }
    if result := initializers.DB.Model(oldProd).Updates(newProd); result.Error != nil{
        return result.Error
    }
    return nil 
}



func (p *Product) Get(id int)(*Product, error){
    pr := &Product{}
    if result := initializers.DB.Model(p).First(pr); result.Error != nil{
        return nil, result.Error
    }
    return pr, nil 
}


func (p *Product) GetAll()([]Product, error){
    var products []Product
    if result := initializers.DB.Model(p).Find(&products); result.Error != nil{
        return nil, result.Error
    }
    return products, nil
}









