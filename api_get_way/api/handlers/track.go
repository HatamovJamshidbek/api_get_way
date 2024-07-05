package handlers

import (
	pb "api_get_way/genproto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (h *Handler) CreateTrack(gn *gin.Context) {
	var track *pb.CreateTrackRequest
	err := gn.ShouldBindJSON(&track)
	if err != nil {
		BadRequest(gn, err)
	}
	track.CompositionId = gn.Param("id")
	response, err := h.Composition.CreateTrack(gn, track)
	if err != nil {
		InternalServerError(gn, err)
	}
	Created(gn, err)
	fmt.Println(response)

}
func (h *Handler) GetTrack(gn *gin.Context) {
	limit1 := gn.Query("limit")
	if len(limit1) == 0 {
		limit1 = limit1 + "0"
	}
	limit, err := strconv.Atoi(limit1)
	if err != nil {

		BadRequest(gn, err)
		return
	}
	offset1 := gn.Query("offset")
	if len(offset1) == 0 {
		offset1 = offset1 + "0"
	}

	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	track := pb.GetTrackRequest{
		Userid:        gn.Query("user_id"),
		LimitOffset:   &pb.Filter{Limit: int32(limit), Offset: int32(offset)},
		CompositionId: gn.Param("composition_id"),
		FileUrl:       gn.Query("file_url"),
		Title:         gn.Query("title"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := h.Composition.GetTrack(ctx, &track)
	if err != nil {
		fmt.Println("+++++++", err)
		InternalServerError(gn, err)
		return
	}

	gn.JSON(200, response)
}

func (h *Handler) PutTrack(gn *gin.Context) {
	compositionID := gn.Param("id")
	trackID := gn.Param("track_id")

	if compositionID == "" || trackID == "" {
		BadRequest(gn, fmt.Errorf("composition ID and track ID must be provided"))
		return
	}

	var trackUpdateRequest pb.UpdateTrackRequest
	if err := gn.BindJSON(&trackUpdateRequest); err != nil {
		BadRequest(gn, err)
		return
	}

	trackUpdateRequest.CompositionId = compositionID
	trackUpdateRequest.Id = trackID

	response, err := h.Composition.UpdateTrack(gn, &trackUpdateRequest)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	fmt.Println(response)
	OK(gn, err)
}

func (h *Handler) DeleteTrack(gn *gin.Context) {
	compositionID := gn.Param("id")
	trackID := gn.Param("trackId")

	if compositionID == "" || trackID == "" {
		BadRequest(gn, fmt.Errorf("composition ID and track ID must be provided"))
		return
	}

	var trackDelete pb.DeleteTrackRequest

	trackDelete.CompositionId = compositionID
	trackDelete.TrackId = trackID

	_, err := h.Composition.DeleteTrack(gn, &trackDelete)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)

}
