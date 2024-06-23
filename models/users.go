package models

import (
	"e-commerce/initializers"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
}


func (u *User) Create(user *User) error {
    result := initializers.DB.Model(user).Create(user)
    if result.Error != nil {
        return result.Error
    }

    return nil 
}

func (u *User) Delete(id int) error {
    result := initializers.DB.Delete(u, id)
    if result.Error != nil {
        return result.Error 
    }

    return nil 
}


func (u *User) Update(newUser *User, id int) error {
    var oldUser *User 
    result := initializers.DB.First(oldUser, id)    
    if result.Error != nil{
        return result.Error 
    }
    updates := initializers.DB.Model(oldUser).Updates(newUser)
    if updates.Error != nil{
        return updates.Error 
    }
    return nil 
}


func (u *User) GetUsers(id int) (*User, error){
    var user *User 
    if result := initializers.DB.First(user, id); result.Error != nil{
        return nil, result.Error
    }

    return user, nil 
}





