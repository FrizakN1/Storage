package database

import (
	"database/sql"
	"storage/utils"
)

var query map[string]*sql.Stmt

type Product struct {
	Name     string
	Category string
	Amount   int
	Date     string
}

func prepareRequest() []string {
	query = make(map[string]*sql.Stmt)
	errors := make([]string, 0)
	var e error

	query["GetCategory"], e = Link.Prepare(`SELECT "Name" FROM "Category"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetProducts"], e = Link.Prepare(`SELECT "Name", "Category" FROM "Product"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddProductStorage"], e = Link.Prepare(`INSERT INTO "Storage"("Product", "Amount") VALUES ($1, $2)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["UpdateProductStorage"], e = Link.Prepare(`UPDATE "Storage" SET "Amount" = $2 WHERE "Product" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetProductStorage"], e = Link.Prepare(`SELECT "Amount" FROM "Storage" WHERE "Product" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddProductJournal"], e = Link.Prepare(`INSERT INTO "Journal"("Product", "Motion", "Date") VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetProductsAmount"], e = Link.Prepare(`SELECT "Amount" FROM "Storage" WHERE "Product" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetJournal"], e = Link.Prepare(`SELECT "Product","Category","Motion", "Date" FROM "Journal" AS j INNER JOIN "Product" AS p ON p."Name" = j."Product" ORDER BY "ID" DESC `)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetJournal() []Product {
	var journal []Product

	stmt, ok := query["GetJournal"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	for rows.Next() {
		var journalElement Product
		e = rows.Scan(&journalElement.Name, &journalElement.Category, &journalElement.Amount, &journalElement.Date)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}
		journalElement.Date = journalElement.Date[0:10]

		journal = append(journal, journalElement)
	}

	return journal
}

func GetStorage() []Product {
	var storage []Product

	stmt, ok := query["GetProducts"]
	if !ok {
		return nil
	}

	stmt1, ok := query["GetProductsAmount"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	for rows.Next() {
		var product Product
		e = rows.Scan(&product.Name, &product.Category)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}

		row := stmt1.QueryRow(product.Name)
		e = row.Scan(&product.Amount)
		if e != nil {
			product.Amount = 0
		}

		storage = append(storage, product)
	}

	return storage
}

func (p *Product) MotionProduct(motion string) string {
	stmt, ok := query["GetProductStorage"]
	if !ok {
		return "Ошибка"
	}

	amount := 0
	check := true
	row := stmt.QueryRow(p.Name)
	e := row.Scan(&amount)
	if e != nil {
		check = false
	}

	if check {
		if motion == "del" {
			if amount > p.Amount {
				p.Amount *= -1
			} else {
				return "Вы пытаетесь использовать больше товара, чем есть на складе"
			}
		}

		stmt, ok = query["UpdateProductStorage"]
		if !ok {
			return "Ошибка"
		}
	} else {
		stmt, ok = query["AddProductStorage"]
		if !ok {
			return "Ошибка"
		}
	}

	_, e = stmt.Exec(p.Name, amount+p.Amount)
	if e != nil {
		utils.Logger.Println(e)
		return "Ошибка"
	}

	stmt, ok = query["AddProductJournal"]
	if !ok {
		return "Ошибка"
	}

	_, e = stmt.Exec(p.Name, p.Amount)
	if e != nil {
		utils.Logger.Println(e)
		return "Ошибка"
	}

	return "Товар использован"
}

func GetProducts() []Product {
	var products []Product

	stmt, ok := query["GetProducts"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	for rows.Next() {
		var product Product
		e = rows.Scan(&product.Name, &product.Category)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}

		products = append(products, product)
	}

	return products
}

func GetCategory() []string {
	var categories []string

	stmt, ok := query["GetCategory"]
	if !ok {
		return nil
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}

	for rows.Next() {
		var category string
		e = rows.Scan(&category)
		if e != nil {
			utils.Logger.Println(e)
			return nil
		}

		categories = append(categories, category)
	}

	return categories
}
