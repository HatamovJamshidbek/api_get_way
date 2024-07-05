package handlers

import "api_get_way/genproto"

type Handler struct {
	Composition   genproto.CompositionServiceClient
	Discovery     genproto.DiscoveryServiceClient
	Collaboration genproto.CollaborationServiceClient
	UserProfile   genproto.UserProfilServerClient
}

func NewHandler(cm genproto.CompositionServiceClient, cl genproto.CollaborationServiceClient, disc genproto.DiscoveryServiceClient, userprofile genproto.UserProfilServerClient) *Handler {
	return &Handler{
		Composition:   cm,
		Collaboration: cl,
		Discovery:     disc,
		UserProfile: userprofile,
	}
}
