package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetTrending(gn *gin.Context) {
	response, err := h.Discovery.GetCompositionTrending(gn, &pb.Void{})
	if err != nil {
		fmt.Println("+++++++++++________+", err)
		InternalServerError(gn, err)
	}
	gn.JSON(200, response)

}
func (h *Handler) GetRecommended(gn *gin.Context) {
	response, err := h.Discovery.GetCompositionRecommend(gn, &pb.Void{})
	if err != nil {
		fmt.Println("++++++++", err)
		InternalServerError(gn, err)
		return
	}
	gn.JSON(200, response)

}
func (h *Handler) GetDiscoveryByGenre(gn *gin.Context) {
	genre := gn.Param("genre")
	if genre == "" {
		gn.JSON(http.StatusBadRequest, gin.H{"error": "Genre parameter is required"})
		return
	}

	// Initialize the genre object
	genreObj := &pb.GetGenre{
		Genre: genre,
	}

	response, err := h.Discovery.GetCompositionGenre(gn, genreObj)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gn.JSON(http.StatusOK, response)
}

func (h *Handler) SearchDiscovery(gn *gin.Context) {
	limit1 := gn.Query("limit")
	if limit1 == "" {
		limit1 += "0"
	}

	limit, err := strconv.Atoi(limit1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	offset1 := gn.Query("offset")
	if offset1 == "" {
		offset1 += "0"
	}
	offset, err := strconv.Atoi(offset1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	like1 := gn.Query("like")
	if like1 == "" {
		like1 += "0"
	}
	likeCount, err := strconv.Atoi(like1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	listen1 := gn.Query("listen")
	if listen1 == "" {
		listen1 += "0"
	}
	listenCount, err := strconv.Atoi(listen1)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	discovery := pb.GetDiscoveryRequest{
		CompositionId: gn.Query("composition_id"),
		Genre:         gn.Query("param"),
		LikeCount:     int64(likeCount),
		ListenCount:   int64(listenCount),
		LimitOffset:   &pb.Filter{Limit: int32(limit), Offset: int32(offset)},
	}
	response, err := h.Discovery.GetDiscovery(gn, &discovery)
	if err != nil {
		InternalServerError(gn, err)
		fmt.Println("+++++++", err)
		return
	}
	gn.JSON(200, response)

}
func (h *Handler) CreateCompositionLike(gn *gin.Context) {
	var request pb.LikeRequest
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}
	request.CompositionId = gn.Param("id")

	_, err := h.Discovery.CreateCompositionLike(gn, &request)
	if err != nil {
		fmt.Println("+++++++++++", err)
		InternalServerError(gn, err)
		return
	}
	Created(gn, err)

}
func (h *Handler) DeleteCompositionLike(gn *gin.Context) {
	var request pb.LikeRequest
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}
	request.CompositionId = gn.Param("id")
	_, err := h.Discovery.DeleteCompositionLike(gn, &request)
	if err != nil {
		fmt.Println("+++++++++", err)
		InternalServerError(gn, err)
		return
	}

	OK(gn, err)
}
func (h *Handler) CreateCompositionListen(gn *gin.Context) {
	var request pb.LikeRequest
	if err := gn.ShouldBindJSON(&request); err != nil {
		BadRequest(gn, err)
		return
	}
	request.CompositionId = gn.Param("id")
	_, err := h.Discovery.CreateCompositionListen(gn, &request)
	if err != nil {
		InternalServerError(gn, err)
		return
	}
	Created(gn, err)

}
