package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/ptr"
)

type TaskServer struct {
}

var _ ServerInterface = TaskServer{}

func NewTaskServer() TaskServer {
	return TaskServer{}
}

// (GET /tasks)
// @Summary List all tasks
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {array} Task
// @Router /tasks [get]
func (s TaskServer) GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, taskStore.List())
}

// (GET /tasks/{id})
// @Summary Get a task by ID
// @Security ApiKeyAuth
// @Param id path string true "Task ID"
// @Success 200 {object} Task
// @Failure 404
// @Router /tasks/{id} [get]
func (s TaskServer) GetTasksId(ctx *gin.Context, id string) {
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	task, ok := taskStore.Get(id)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// @Summary Create a new task
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param task body NewTask true "New task"
// @Success 201 {object} Task
// @Router /tasks [post]
func (s TaskServer) PostTasks(ctx *gin.Context) {
	var task Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := Task{
		Title:     task.Title,
		Completed: ptr.To(false),
	}

	id := taskStore.Add(resp)

	resp.Id = id

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Delete a task
// @Security ApiKeyAuth
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404
// @Router /tasks/{id} [delete]
func (s TaskServer) DeleteTasksId(ctx *gin.Context, id string) {
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	if _, ok := taskStore.tasks[id]; !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	taskStore.Delete(id)

	ctx.JSON(http.StatusNoContent, nil)
}

// @Summary Update a task
// @Security ApiKeyAuth
// @Param id path string true "Task ID"
// @Param task body TaskPatch true "TaskPatch"
// @Success 200 {object} Task
// @Failure 404
// @Router /tasks/{id} [patch]
func (s TaskServer) PatchTasksId(ctx *gin.Context, id string) {
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	task, ok := taskStore.Get(id)
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	var taskPatch TaskPatch
	if err := ctx.ShouldBindJSON(&taskPatch); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if taskPatch.Completed != nil {
		task.Completed = taskPatch.Completed
	}

	taskStore.tasks[id] = task

	ctx.JSON(http.StatusOK, task)
}
