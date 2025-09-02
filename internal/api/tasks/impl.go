package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/model"
	"gitlab.ixcloud.ch/ZimmermannRoger/gin-todo/internal/repository"
)

type PollServer struct {
	polls *repository.Polls
}

var _ ServerInterface = PollServer{}

func NewPollServer() PollServer {
	p, _ := repository.NewPolls("./data/poll.sqlite")
	return PollServer{
		polls: p,
	}
}

// (GET /poll)
func (s PollServer) GetPoll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.ConcretePoll)
}

// (GET /vote)
func (s PollServer) GetVote(ctx *gin.Context) {
	cookie := ctx.GetString("voter_id")
	if cookie == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "voter_id cookie is missing"})
		return
	}

	fmt.Println("voter_id:", ctx.GetString("voter_id"))

	choiceId, err := s.polls.GetVote(model.PollID, ctx.GetString("voter_id"))
	if err != nil && err == repository.ErrNoRows {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve vote", "details": err.Error()})
		return
	}
	if choiceId == "" {
		ctx.JSON(http.StatusOK, gin.H{"message": "No vote found for this voter"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"choiceId": choiceId})
}

// (POST /vote)
func (s PollServer) PostVote(ctx *gin.Context) {
	var vote model.Vote

	if err := ctx.ShouldBindJSON(&vote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if choice is valid
	if _, ok := model.DefinedPollOptions[vote.ChoiceID]; !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid choice", "details": "choice must be one of the defined options",
			"choices": model.DefinedPollOptions})
		return
	}

	if err := s.polls.UpsertVote(model.PollID, ctx.GetString("voter_id"), vote.ChoiceID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not record vote", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vote received", "option": vote.ChoiceID})
}

// (GET /results)
func (s PollServer) GetResults(ctx *gin.Context) {
	// Here you would typically fetch the poll rawResults from a database
	rawResults, err := s.polls.GetResults(model.PollID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve results", "details": err.Error()})
		return
	}
	if rawResults == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve results"})
		return
	}

	totalVotes := 0
	for _, count := range rawResults {
		totalVotes += count
	}

	var options []model.OptionCounts = make([]model.OptionCounts, 0, len(model.DefinedPollOptions))
	for optionID, count := range rawResults {
		var option model.OptionCounts
		option.ID = optionID
		option.Label = model.DefinedPollOptions[optionID]
		option.Votes = count
		options = append(options, option)
	}

	results := &model.Results{
		PollID:     model.PollID,
		TotalVotes: totalVotes,
		Options:    options,
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
