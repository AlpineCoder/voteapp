package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/model"
)

type PollServer struct {
}

var _ ServerInterface = PollServer{}

func NewPollServer() PollServer {
	return PollServer{}
}

// (GET /poll)
func (s PollServer) GetPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.DefinedPollOptions)
}

// (GET /vote)
func (s PollServer) GetVote(ctx *gin.Context) {
	cookie := ctx.GetString("voter_id")
	if cookie == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "voter_id cookie is missing"})
		return
	}

	fmt.Println("voter_id:", ctx.GetString("voter_id"))

	ctx.JSON(http.StatusOK, gin.H{"message": "Vote endpoint"})
}

// (POST /vote)
func (s PollServer) PostVote(ctx *gin.Context) {
	var vote struct {
		Option string `json:"option" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&vote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would typically process the vote, e.g., store it in a database

	ctx.JSON(http.StatusOK, gin.H{"message": "Vote received", "option": vote.Option})
}

// (GET /results)
func (s PollServer) GetResults(ctx *gin.Context) {
	// Here you would typically fetch the poll results from a database
	results := map[string]int{
		"optionA": 10,
		"optionB": 5,
		"optionC": 3,
	}

	ctx.JSON(http.StatusOK, gin.H{"results": results})
}

// (GET /tasks)
// func (s PollServer) GetTasks(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, taskStore.List())
// }

// // (GET /tasks/{id})
// func (s PollServer) GetTasksId(ctx *gin.Context, id string) {
// 	if id == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
// 		return
// 	}

// 	task, ok := taskStore.Get(id)
// 	if !ok {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, task)
// }

// // PostTasks
// func (s PollServer) PostTasks(ctx *gin.Context) {
// 	var task Task
// 	if err := ctx.ShouldBindJSON(&task); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	resp := Task{
// 		Title:     task.Title,
// 		Completed: ptr.To(false),
// 	}

// 	id := taskStore.Add(resp)

// 	resp.Id = id

// 	ctx.JSON(http.StatusCreated, resp)
// }

// func (s PollServer) DeleteTasksId(ctx *gin.Context, id string) {
// 	if id == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
// 		return
// 	}

// 	if _, ok := taskStore.tasks[id]; !ok {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
// 		return
// 	}
// 	taskStore.Delete(id)

// 	ctx.JSON(http.StatusNoContent, nil)
// }

// func (s PollServer) PatchTasksId(ctx *gin.Context, id string) {
// 	if id == "" {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
// 		return
// 	}

// 	task, ok := taskStore.Get(id)
// 	if !ok {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
// 		return
// 	}

// 	var taskPatch TaskPatch
// 	if err := ctx.ShouldBindJSON(&taskPatch); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if taskPatch.Completed != nil {
// 		task.Completed = taskPatch.Completed
// 	}

// 	taskStore.tasks[id] = task

// 	ctx.JSON(http.StatusOK, task)
// }
