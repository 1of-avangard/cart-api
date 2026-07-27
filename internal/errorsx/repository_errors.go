package errorsx

import "errors"

var ErrCartNotFound = errors.New("cart not found")
var ErrItemNotFound = errors.New("item not found")
var ErrProductNotFound = errors.New("product name cannot be blank")
var ErrPriceNotFound = errors.New("price must be positive")
var ErrCartFull = errors.New("cart cannot contain more than 5 products")
