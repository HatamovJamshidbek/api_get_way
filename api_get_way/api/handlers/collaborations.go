package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
)

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func (h *Handler) CreateCollaborations(gn *gin.Context) {
	var collaboration pb.CreateCollaborationRequest
	err := gn.ShouldBindJSON(&collaboration)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	compositionId := gn.Param("id")
	if !isValidUUID(compositionId) {
		BadRequest(gn, fmt.Errorf("invalid composition ID"))
		return
	}

	collaboration.CompositionId = compositionId
	_, err = h.Collaboration.CreateCollaborators(gn, &collaboration)
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(gn, err)
		return
	}

	Created(gn, err)
}

func (h *Handler) GetCollaboration(gn *gin.Context) {
	limit1 := gn.Query("limit")
	if limit1 == "" {
		limit1 = "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, fmt.Errorf("invalid limit value"))
		return
	}

	offset1 := gn.Query("offset")
	if offset1 == "" {
		offset1 = "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, fmt.Errorf("invalid offset value"))
		return
	}

	compositionId := gn.Param("id")
	if !isValidUUID(compositionId) {
		BadRequest(gn, fmt.Errorf("invalid composition ID"))
		return
	}

	fmt.Printf("Request Parameters - Limit: %d, Offset: %d, Role: %s, UserID: %s, CompositionID: %s\n", limit, offset, gn.Query("role"), gn.Query("user_id"), compositionId)

	collaboration := pb.GetCollaboratorsRequest{
		LimitOffset:   &pb.Filter{Limit: int32(limit), Offset: int32(offset)},
		Role:          gn.Query("role"),
		UserId:        gn.Query("user_id"),
		CompositionId: compositionId,
	}

	fmt.Printf("gRPC Request: %+v\n", collaboration)

	response, err := h.Collaboration.GetCollaborators(gn, &collaboration)
	if err != nil {
		fmt.Println("+++++++++++++", err)
		InternalServerError(gn, err)
		return
	}
	fmt.Printf("gRPC Response: %+v\n", response)
	gn.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCollaboration(gn *gin.Context) {
	fmt.Println("+++++++++")
	var collaboration *pb.UpdateCollaborationRequest
	err := gn.ShouldBindJSON(&collaboration)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	collaboration.CompositionId = gn.Param("id")
	collaboration.Userid = gn.Param("userId")
	response, err := h.Collaboration.UpdateCollaborators(gn, collaboration)
	if err != nil {
		fmt.Println("++++++", err)
		InternalServerError(gn, err)
		return
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) DeleteCollaboration(gn *gin.Context) {
	collaboration := pb.DeleteCollaborationRequest{
		CompositionId: gn.Param("id"),
		Userid:        gn.Param("userId"),
	}
	response, err := h.Collaboration.DeleteCollaborators(gn, &collaboration)
	if err != nil {
		fmt.Println("+++++++++++", err)
		InternalServerError(gn, err)
		return
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) CreateComment(gn *gin.Context) {
	var commit *pb.CreateCommitRequest
	err := gn.ShouldBindJSON(&commit)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	commit.CompositionId = gn.Param("id")
	response, err := h.Collaboration.CreateComment(gn, commit)
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(gn, err)
		return
	}
	fmt.Println(response)
	OK(gn, err)

}
func (h *Handler) GetComment(gn *gin.Context) {
	limitStr := gn.Query("limit")
	if limitStr == "" {
		limitStr = "0"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	offsetStr := gn.Query("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	commit := pb.GetCommitRequest{
		CompositionId: gn.Param("id"),
		LimitOffset:   &pb.Filter{Limit: int32(limit), Offset: int32(offset)},
		UserId:        gn.Query("user_id"),
		Content:       gn.Query("content"),
	}

	response, err := h.Collaboration.GetComment(gn, &commit)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(200, response)
}
