package models

import "e-commerce/initializers"

type Address struct {
	ID          int
	City        string `gorm:"type:varchar(50)" json:"city"`
	Country     string `gorm:"type:varchar(50)" json:"country"`
	AddressLine string `gorm:"type:varchar(50)" json:"addressLine"`
}

func (a *Address) Create(addr *Address) error {
	result := initializers.DB.Create(addr)
    if result.Error != nil {
        return result.Error
    }
    return nil 
}

func (a *Address) Delete(id int) error {
    result := initializers.DB.Delete(a, id)
    if result.Error != nil {
        return result.Error
    }

    return nil 
}

func (a *Address) UpdateAddr(newAddr *Address, id int) error{
    var oldAddr *Address
    result := initializers.DB.First(oldAddr, id)
    if result.Error != nil {
        return result.Error
    }
    updates := initializers.DB.Model(oldAddr).Updates(newAddr)
    if updates.Error != nil{
        return updates.Error
    }

    return nil 

}
