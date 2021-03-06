package models

import (
	"fmt"
	"strconv"
)

type Item struct {
	Id       int
	UserId   int
	IsDone   int
	ItemName string
}

func GetAllItems() ([]Item, error) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	itemList := make([]Item, 0)
	for rows.Next() {
		var item Item
		if err = rows.Scan(&item.Id, &item.UserId, &item.IsDone, &item.ItemName); err != nil {
			return nil, err
		}
		itemList = append(itemList, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return itemList, err
}

func GetTotalCount(pageSize string) (int, error) {
	intPageSize, _ := strconv.Atoi(pageSize)
	var totalRows int
	rows, err := db.Query(fmt.Sprintf("SELECT COUNT(id) FROM items"))
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err = rows.Scan(&totalRows); err != nil {
			return 0, err
		}
	}

	if (totalRows % intPageSize) == 0 {
		return (totalRows / intPageSize), err
	}
	return (totalRows / intPageSize) + 1, err
}

func GetPageOfItems(page, pageSize string) ([]Item, error) {
	intPage, _ := strconv.Atoi(page)
	intPageSize, _ := strconv.Atoi(pageSize)

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM items LIMIT %d OFFSET %d", intPageSize, ((intPage - 1) * intPageSize)))
	if err != nil {
		return nil, err
	}
	itemList := make([]Item, 0)
	for rows.Next() {
		var item Item
		if err = rows.Scan(&item.Id, &item.UserId, &item.IsDone, &item.ItemName); err != nil {
			return nil, err
		}
		itemList = append(itemList, item)
	}
	return itemList, err
}

func AddItem(userId string, itemName string) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO items(user_id, item_name, is_done) VALUE('%s', '%s', '%d')", userId, itemName, 0))
	return err
}

func UpdateItem(isDone string, id string) error {
	_, err := db.Exec(fmt.Sprintf("Update items SET is_done=%s WHERE id=%s", isDone, id))
	return err
}

func DeleteAll() error {
	_, err := db.Exec(fmt.Sprintf("DELETE FROM items"))
	return err
}

func DeleteItem(itemId string) error {
	_, err := db.Exec(fmt.Sprintf("DELETE FROM items WHERE id=%s", itemId))
	return err
}
