package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"sushkof_test/internal/helper"
	"sushkof_test/internal/model"
)

type TaskRepoImpl struct {
	Db *sql.DB
}

func NewTaskRepository(Db *sql.DB) TaskRepo {
	return &TaskRepoImpl{Db: Db}
}

func (t *TaskRepoImpl) FindByName(ctx context.Context, title string) (model.Task, error) {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT id_task, title, description, status FROM tasks WHERE title=? and is_active=TRUE"
	result, err := tx.QueryContext(ctx, query, title)
	helper.PanicIfError(err)

	var task model.Task

	if result.Next() {
		err := result.Scan(&task.IdTask, &task.Title, &task.Description, &task.Status)
		helper.PanicIfError(err)
		return task, nil
	} else {
		return task, errors.New(fmt.Sprintf("Task with title %s not found", title))
	}
}

func (t *TaskRepoImpl) FindById(ctx context.Context, id int) (model.Task, error) {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT 1 FROM tasks WHERE id_task=? and is_active=TRUE"
	result, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)

	var task model.Task

	if result.Next() {
		err := result.Scan(&task.IdTask, &task.Title, &task.Status, &task.Description)
		helper.PanicIfError(err)
		return task, nil
	} else {
		return task, errors.New(fmt.Sprintf("Task with id %d not found", id))
	}
}

func (t *TaskRepoImpl) FindAll(ctx context.Context) []model.Task {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT * FROM tasks WHERE is_active=TRUE"
	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	defer result.Close()
	var tasks []model.Task

	for result.Next() {
		task := model.Task{}
		err = result.Scan(&task.IdTask, &task.Title, &task.Description, &task.Status, &task.IsActive, &task.UserId)
		helper.PanicIfError(err)

		tasks = append(tasks, task)
	}

	return tasks
}

func (t *TaskRepoImpl) Save(ctx context.Context, task model.Task) {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	task.IsActive = true
	query := "INSERT INTO tasks(title, description, status, is_active) VALUES (?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, task.Title, task.Description, task.Status, task.IsActive)
	helper.PanicIfError(err)

}

func (t *TaskRepoImpl) Update(ctx context.Context, task model.Task) {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE tasks SET title=?, description=?, status=?, updated_at=NOW() WHERE id_task=?"
	_, err = tx.ExecContext(ctx, query, task.Title, task.Description, task.Status, task.IdTask)
	helper.PanicIfError(err)

}

func (t *TaskRepoImpl) Delete(ctx context.Context, idTask int) {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE tasks SET is_active=FALSE WHERE id_task=?"
	_, err = tx.ExecContext(ctx, query, idTask)
	helper.PanicIfError(err)

	log.Info().Msgf("Task has been deleted")
}

func (t *TaskRepoImpl) GetTasksById(ctx context.Context, id int) []model.Task {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT title, description, status FROM tasks WHERE user_id = ? AND is_active=TRUE ORDER BY created_at DESC"
	result, err := tx.QueryContext(ctx, query, id)
	helper.PanicIfError(err)

	defer result.Close()

	var tasks []model.Task

	for result.Next() {
		task := model.Task{}
		err := result.Scan(&task.Title, &task.Description, &task.Status)
		helper.PanicIfError(err)

		tasks = append(tasks, task)
	}
	return tasks
}

func (t *TaskRepoImpl) AssignTaskToUser(ctx context.Context, title string, userId int) {
	tx, err := t.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE tasks SET user_id=? WHERE title=?"
	_, err = tx.ExecContext(ctx, query, userId, title)
	helper.PanicIfError(err)

	log.Info().Msgf("New task added to user")
}
