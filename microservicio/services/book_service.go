package services

import (
	itemDao "microservicio/daos/item"
	"microservicio/dtos"
	model "microservicio/models"
	e "microservicio/utils/errors"
)

type itemService struct{}

type itemServiceInterface interface {
	//GetBook(id string) (dtos.ItemDto, e.ApiError)

	InsertItem(bookDto dtos.ItemDto) (dtos.ItemDto, e.ApiError)
}

var (
	ItemService itemServiceInterface
)

func init() {
	ItemService = &itemService{}
}

func (s *itemService) InsertItem(itemDto dtos.ItemDto) (dtos.ItemDto, e.ApiError) {

	var item model.Item

	item.Tittle = itemDto.Tittle
	item.Seller = itemDto.Seller
	item.Price = itemDto.Price
	item.Currency = itemDto.Currency
	item.Pictures = itemDto.Pictures
	item.Description = itemDto.Description
	item.State = itemDto.State
	item.City = itemDto.City
	item.Street = itemDto.Street
	item.Number = itemDto.Number

	item = itemDao.InsertItem(item)

	if item.Id.Hex() == "000000000000000000000000" {
		return itemDto, e.NewBadRequestApiError("error in insert")
	}

	itemDto.Id = item.Id.Hex()

	return itemDto, nil
}

/*func (s *bookService) GetBook(id string) (dtos.BookDto, e.ApiError) {

	var book model.Book = bookDao.GetById(id)
	var bookDto dtos.BookDto

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, e.NewBadRequestApiError("book not found")
	}
	bookDto.Name = book.Name
	bookDto.Id = book.Id.Hex()
	return bookDto, nil
}



func (s *bookService) InsertBook(bookDto dtos.BookDto) (dtos.BookDto, e.ApiError) {

	var book model.Book

	book.Name = bookDto.Name

	book = bookDao.Insert(book)

	if book.Id.Hex() == "000000000000000000000000" {
		return bookDto, e.NewBadRequestApiError("error in insert")
	}
	bookDto.Id = book.Id.Hex()

	return bookDto, nil
}*/
