package handlers

import (
	"backend/internal/database"
	"backend/internal/dto"
	"backend/internal/entity"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type VoteHandler struct {
	VoteDB database.VoteInterface
}

func NewVoteHandler(voteDB database.VoteInterface) *VoteHandler {
	return &VoteHandler{
		VoteDB: voteDB,
	}
}

func (h *VoteHandler) Create(w http.ResponseWriter, r *http.Request){
	var voteData dto.VoteInput
	err := json.NewDecoder(r.Body).Decode(&voteData)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		handleBadRequest(w, err.Error())
		return 
	}

	userId := claims["userId"].(float64)


	vote := entity.NewVote(uint(userId), voteData.ItemId, voteData.Value)

	_, err = h.VoteDB.FindByIds(uint(userId), vote.ItemID)
	if err == nil {
		handleBadRequest(w, "você já votou neste item")
		return
	}

	err = h.VoteDB.Create(vote)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}