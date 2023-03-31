package repository

import (
	"book-rest-api/model"
	"book-rest-api/repository/query"
)

type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.AddBook,
		in.Bookname,
		in.Author,
	).Scan(
		&res.ID,
		&res.Bookname,
		&res.Author,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	rows, err := r.db.Query(query.GetBooks)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err = rows.Scan(
			&book.ID,
			&book.Bookname,
			&book.Author,
		)
		if err != nil {
			return res, err
		}
		res = append(res, book)
	}

	if err := rows.Err(); err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBookById(id int64) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.GetBook,
		id,
	).Scan(
		&res.ID,
		&res.Bookname,
		&res.Author,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.UpdateBook,
		in.Bookname,
		in.Author,
		in.ID,
	).Scan(
		&res.ID,
		&res.Bookname,
		&res.Author,
	)

	if err != nil {
		return res, err
	}

	return in, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	_, err = r.db.Exec(query.DeleteBook, id)
	if err != nil {
		panic(err)
	}

	return nil
}
