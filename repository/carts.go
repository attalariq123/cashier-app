package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
)

type CartRepository struct {
	db db.DB
}

func NewCartRepository(db db.DB) CartRepository {
	return CartRepository{db}
}

func (u *CartRepository) ReadCart() ([]model.Cart, error) {
	records, err := u.db.Load("carts")
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("Cart not found!")
	}

	var cart []model.Cart
	err = json.Unmarshal([]byte(records), &cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (u *CartRepository) UpdateCart(cart model.Cart) error {
	listCart, err := u.ReadCart()
	if err != nil {
		return err
	}

	targetCart, err := u.CartUserExist(cart.Name)
	if err != nil {
		return err
	}

	for idx, element := range listCart {
		if element.Name == targetCart.Name {
			listCart[idx] = cart
		}
	}

	jsonData, err := json.Marshal(listCart)
	if err != nil {
		return err
	}

	err = u.db.Save("carts", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *CartRepository) AddCart(cart model.Cart) error {

	listCart, err := u.ReadCart()
	if err != nil {
		return err
	}

	listCart = append(listCart, cart)
	cartData, err := json.Marshal(listCart)
	if err != nil {
		return err
	}

	err = u.db.Save("carts", cartData)
	if err != nil {
		return err
	}

	return nil
}

func (u *CartRepository) ResetCarts() error {
	err := u.db.Reset("carts", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *CartRepository) CartUserExist(name string) (model.Cart, error) {
	listcCart, err := u.ReadCart()
	if err != nil {
		return model.Cart{}, err
	}
	for _, element := range listcCart {
		if element.Name == name {
			return element, nil
		}
	}
	return model.Cart{}, fmt.Errorf("Cart Empty!")
}
