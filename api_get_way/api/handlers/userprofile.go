package handlers

import (
	pb "api_get_way/genproto"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(gn *gin.Context) {
	var req pb.UserRequest
	err := gn.ShouldBind(&req)
	if err != nil {
		BadRequest(gn, err)
	}

	resp, err := h.UserProfile.UpdateUserProfiles(gn, &req)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	OK(gn, err)
	fmt.Println(resp)
}

func (h *Handler) CreateUser(gn *gin.Context) {
	var req pb.UserRequest

	err := gn.Bind(&req)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	fmt.Println("++++++++++", &req)

	//_, err = uuid.Parse(req.UserID)
	//if err != nil {
	//	BadRequest(gn, err)
	//	return
	//}
	//_, err = uuid.Parse(req.Id)
	//if err != nil {
	//	BadRequest(gn, err)
	//	return
	//}

	resp, err := h.UserProfile.CreateUserProfiles(gn, &req)
	if err != nil {
		fmt.Println("++++++++++", err)
		InternalServerError(gn, err)
		return
	}
	OK(gn, err)

	fmt.Println(resp)
}

func (h *Handler) DeleteUser(gn *gin.Context) {
	var user pb.UserID

	user.UserID = gn.Param("id")

	resp, err := h.UserProfile.DeleteUserProfiles(gn, &user)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	OK(gn, err)
	fmt.Println(resp)
}

func (h *Handler) GetbyUserID(gn *gin.Context) {
	var req pb.UserID

	req.UserID = gn.Param("id")

	resp, err := h.UserProfile.GetByUserID(gn, &req)
	if err != nil {
		InternalServerError(gn, err)
	}

	gn.JSON(200, resp)
}

func (h *Handler) GetallUserP(gn *gin.Context) {
	var req pb.UserFilter

	req.Fulname = gn.Query("fulname")

	req.Location = gn.Query("location")

	req.UpdatedAt = gn.Query("updated_at")

	req.UserRole = gn.Query("user_role")

	req.UserID = gn.Query("user_id")

	l := gn.Query("limmit")

	if len(l) == 0 {
		l = "0"
	}

	limmint, err := strconv.Atoi(l)
	if err != nil {
		BadRequest(gn, err)
		return
	}

	o := gn.Query("offset")
	if len(o) == 0 {
		o = "0"
	}

	offset, err := strconv.Atoi(o)
	if err != nil {
		BadRequest(gn, err)
		return
	}
	req.GetUser.Limit = int32(limmint)
	req.GetUser.Offset = int32(offset)

	resp, err := h.UserProfile.GetAllUser(gn, &req)
	if err != nil {
		InternalServerError(gn, err)
		return
	}

	gn.JSON(200, resp)
}
