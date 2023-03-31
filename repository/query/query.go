package query

const (
	AddBook = `
		INSERT INTO 
		db_go_sql.books (
			book_name,
			author
		) 
		VALUES ($1, $2)
		RETURNING *
	`
	GetBooks = `
		SELECT * FROM db_go_sql.books
	`
	GetBook = `
		SELECT *
		FROM db_go_sql.books 
		WHERE id=$1 
	`
	UpdateBook = `
		UPDATE db_go_sql.books SET "book_name"=$1, "author"=$2 WHERE "id"=$3
		RETURNING *
	`
	DeleteBook = `
		DELETE FROM db_go_sql.books WHERE id=$1	
	`
)
