package db

import (
    "database/sql"
    "time"
)

type Task struct {
    ID      string  `json:"id"`
    Date    string `json:"date"`
    Title   string `json:"title"`
    Comment string `json:"comment"`
    Repeat  string `json:"repeat"`
}

// Механизм добавления задачи в таблицу scheduler
func AddTask(task *Task) (int64, error) {
    query := `INSERT INTO scheduler(date, title, comment, repeat) VALUES (?, ?, ?, ?)`
    res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}
// Механизм вывода задач из таблицы scheduler и поиска
func Tasks(limit int, search string) ([]*Task, error) {
    var rows *sql.Rows
    var err error  
	
	if search == "" {
		rows, err = DB.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?", limit)
    } else {
        // парсинг даты
        t, parseErr := time.Parse("02.01.2006", search)
        if parseErr == nil {
            date := t.Format("20060102")
            rows, err = DB.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? ORDER BY date ASC LIMIT ?", date, limit)
        } else {
            like := "%" + search + "%"
            rows, err = DB.Query("SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date ASC LIMIT ?", like, like, limit)
        }
    }
	
	if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []*Task
    for rows.Next() {
        var t Task
        if err := rows.Scan(&t.ID, &t.Date, &t.Title, &t.Comment, &t.Repeat); err != nil {
            return nil, err
        }
        tasks = append(tasks, &t)
    }

    if tasks == nil { // нет задач
        tasks = []*Task{} // пустой слайс
    }
    return tasks, nil
}