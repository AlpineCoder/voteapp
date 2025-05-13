package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/ptr"
)

// optional code omitted

type Server struct {
}

func NewServer() Server {
	return Server{}
}

// (GET /tasks)
// @Summary List all tasks
// @Produce json
// @Success 200 {array} Task
// @Router /tasks [get]
func (Server) GetTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, taskStore.List())
}

// @Summary Create a new task
// @Accept json
// @Produce json
// @Param task body NewTask true "New task"
// @Success 201 {object} Task
// @Router /tasks [post]
func (Server) PostTasks(ctx *gin.Context) {
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
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404
// @Router /tasks/{id} [delete]
func (Server) DeleteTasksId(ctx *gin.Context, id string) {
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
