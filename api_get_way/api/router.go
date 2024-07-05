package api

import (
	"api_get_way/api/handlers"
	"api_get_way/genproto"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func RouterAPi(conn1 *grpc.ClientConn, conn2 *grpc.ClientConn, conn3 *grpc.ClientConn, conn4 *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	compServ := genproto.NewCompositionServiceClient(conn1)
	collServ := genproto.NewCollaborationServiceClient(conn2)
	discServ := genproto.NewDiscoveryServiceClient(conn3)
	userprofile := genproto.NewUserProfilServerClient(conn4)

	h := handlers.NewHandler(compServ, collServ, discServ, userprofile)

	router.POST("/api/compositions", h.CreateComposition)
	router.GET("/api/compositions/:id", h.GetCompositionById)
	router.DELETE("/api/compositions/:id", h.DeleteComposition)
	router.PUT("/api/compositions/:id", h.UpdateComposition)
	router.GET("/api/users/:id/compositions", h.GetCompositionByUserId)
	router.POST("/api/compositions/:id/tracks", h.CreateTrack)
	router.GET("/api/compositions/:id/tracks", h.GetTrack)
	router.PUT("/api/compositions/:id/tracks/:track_id", h.PutTrack)
	router.DELETE("/api/compositions/:id/tracks/:trackId", h.DeleteTrack)

	router.POST("/api/collaborations/invite", h.CreateInvite)
	router.PUT("/api/collaborations/invite/:id/respond", h.UpdateInvite)
	router.POST("/api/composition/:id/collaborators", h.CreateCollaborations)
	router.GET("/api/compositions/:id/collaborators", h.GetCollaboration)
	router.PUT("/api/compositions/:id/collaborators/:userId", h.UpdateCollaboration)
	router.DELETE("/api/compositions/:id/collaborators/:userId", h.DeleteCollaboration)
	router.POST("/api/compositions/:id/comments", h.CreateComment)
	router.GET("/api/compositions/:id/comments", h.GetComment)

	router.GET("/api/discover/trending", h.GetTrending)                // Assuming a handler function for trending
	router.GET("/api/discover/recommended", h.GetRecommended)          // Assuming a handler function for recommended
	router.GET("/api/discover/genres/:genre", h.GetDiscoveryByGenre)   // Assuming a handler function for genres
	router.GET("/api/search", h.SearchDiscovery)                       // Assuming a handler function for search
	router.POST("/api/compositions/:id/like", h.CreateCompositionLike) // Assuming a handler function for liking compositions
	router.DELETE("/api/compositions/:id/like", h.DeleteCompositionLike)
	router.POST("/api/compositions/:id/listen", h.CreateCompositionListen)

	// api user profile
	router.POST("/api/userprofile", h.CreateUser)
	router.PUT("/api/userprofileU", h.UpdateUser)
	router.DELETE("/api/userprofileD", h.DeleteUser)
	router.GET("/api/userprofile", h.GetbyUserID)
	router.GET("/api/usersprofile", h.GetallUserP)

	return router

}
