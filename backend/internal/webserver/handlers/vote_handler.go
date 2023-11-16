package handlers

import (
	"backend/internal/database"
	"backend/internal/dto"
	"backend/internal/entity"
	"encoding/json"
	"fmt"
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

	voteExists, err := h.VoteDB.FindByIds(uint(userId), vote.ItemID)
	if err == nil {
		voteExists.Value = voteData.Value;
		err = h.VoteDB.Update(voteExists)
		if err != nil {
			handleBadRequest(w, err.Error())
			return
		}
	} else {
		err = h.VoteDB.Create(vote)
		if err != nil {
			handleBadRequest(w, err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *VoteHandler) List(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		handleBadRequest(w, err.Error())
		return 
	}

	userId := uint(claims["userId"].(float64))

	votes, err := h.VoteDB.FindByUserId(userId)
	if err != nil {
		handleBadRequest(w, err.Error())
		return
	}

	fmt.Println(votes)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(votes)
}