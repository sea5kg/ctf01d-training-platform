// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// Defines values for UserRequestRole.
const (
	UserRequestRoleAdmin  UserRequestRole = "admin"
	UserRequestRoleGuest  UserRequestRole = "guest"
	UserRequestRolePlayer UserRequestRole = "player"
)

// Defines values for UserResponseRole.
const (
	UserResponseRoleAdmin  UserResponseRole = "admin"
	UserResponseRoleGuest  UserResponseRole = "guest"
	UserResponseRolePlayer UserResponseRole = "player"
)

// GameRequest defines model for GameRequest.
type GameRequest struct {
	// Description A brief description of the game
	Description *string `json:"description,omitempty"`

	// EndTime The end time of the game
	EndTime time.Time `json:"end_time"`

	// StartTime The start time of the game
	StartTime time.Time `json:"start_time"`
}

// GameResponse defines model for GameResponse.
type GameResponse struct {
	// Description A brief description of the game
	Description *string `json:"description,omitempty"`

	// EndTime The end time of the game
	EndTime time.Time `json:"end_time"`

	// Id Unique identifier for the game
	Id int `json:"id"`

	// StartTime The start time of the game
	StartTime time.Time `json:"start_time"`
}

// ResultRequest defines model for ResultRequest.
type ResultRequest struct {
	// GameId Identifier of the game this result is for
	GameId string `json:"game_id"`

	// Rank The rank achieved by the team in this game
	Rank int `json:"rank"`

	// Score The score achieved by the team
	Score int `json:"score"`

	// TeamId Identifier of the team this result belongs to
	TeamId string `json:"team_id"`
}

// ResultResponse defines model for ResultResponse.
type ResultResponse struct {
	// GameId Identifier of the game this result is for
	GameId string `json:"game_id"`

	// Id Unique identifier for the result entry
	Id int `json:"id"`

	// Rank The rank achieved by the team in this game
	Rank int `json:"rank"`

	// Score The score achieved by the team
	Score int `json:"score"`

	// TeamId Identifier of the team this result belongs to
	TeamId string `json:"team_id"`
}

// ServiceRequest defines model for ServiceRequest.
type ServiceRequest struct {
	// Author Author of the service
	Author string `json:"author"`

	// Description A brief description of the service
	Description *string `json:"description,omitempty"`

	// IsPublic Boolean indicating if the service is public
	IsPublic bool `json:"is_public"`

	// LogoUrl URL to the logo of the service
	LogoUrl *string `json:"logo_url,omitempty"`

	// Name Name of the service
	Name string `json:"name"`
}

// ServiceResponse defines model for ServiceResponse.
type ServiceResponse struct {
	// Author Author of the service
	Author string `json:"author"`

	// Description A brief description of the service
	Description *string `json:"description,omitempty"`

	// Id Unique identifier for the service
	Id int `json:"id"`

	// IsPublic Boolean indicating if the service is public
	IsPublic bool `json:"is_public"`

	// LogoUrl URL to the logo of the service
	LogoUrl *string `json:"logo_url,omitempty"`

	// Name Name of the service
	Name string `json:"name"`
}

// TeamRequest defines model for TeamRequest.
type TeamRequest struct {
	// AvatarUrl URL to the team's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Description A brief description of the team
	Description *string `json:"description,omitempty"`

	// Name Name of the team
	Name string `json:"name"`

	// SocialLinks JSON string containing social media links of the team
	SocialLinks *string `json:"social_links,omitempty"`

	// UniversityId University or institution the team is associated with
	UniversityId int `json:"university_id"`
}

// TeamResponse defines model for TeamResponse.
type TeamResponse struct {
	// AvatarUrl URL to the team's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Description A brief description of the team
	Description *string `json:"description,omitempty"`

	// Id Unique identifier for the team
	Id int `json:"id"`

	// Name Name of the team
	Name string `json:"name"`

	// SocialLinks JSON string containing social media links of the team
	SocialLinks *string `json:"social_links,omitempty"`

	// University University or institution the team is associated with
	University *string `json:"university,omitempty"`
}

// UniversitiesResponse defines model for UniversitiesResponse.
type UniversitiesResponse = []UniversityResponse

// UniversityResponse defines model for UniversityResponse.
type UniversityResponse struct {
	// Id The unique identifier of the university
	Id int `json:"id"`

	// Name The name of the university
	Name string `json:"name"`
}

// UserRequest defines model for UserRequest.
type UserRequest struct {
	// AvatarUrl URL to the user's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Password User password
	Password *string `json:"password,omitempty"`

	// Role The role of the user (admin, player or guest)
	Role *UserRequestRole `json:"role,omitempty"`

	// Status Status of the user (active, disabled)
	Status  *string `json:"status,omitempty"`
	TeamIds *[]int  `json:"team_ids,omitempty"`

	// UserName The name of the user
	UserName *string `json:"user_name,omitempty"`
}

// UserRequestRole The role of the user (admin, player or guest)
type UserRequestRole string

// UserResponse defines model for UserResponse.
type UserResponse struct {
	// AvatarUrl URL to the user's avatar
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// Id The unique identifier for the user
	Id *int `json:"id,omitempty"`

	// Role The role of the user (admin, player or guest)
	Role *UserResponseRole `json:"role,omitempty"`

	// Status Status of the user (active, disabled)
	Status *string `json:"status,omitempty"`

	// UserName The name of the user
	UserName *string `json:"user_name,omitempty"`
}

// UserResponseRole The role of the user (admin, player or guest)
type UserResponseRole string

// GetApiUniversitiesParams defines parameters for GetApiUniversities.
type GetApiUniversitiesParams struct {
	// Term Optional search term to filter universities by name.
	Term *string `form:"term,omitempty" json:"term,omitempty"`
}

// PostApiV1AuthSigninJSONBody defines parameters for PostApiV1AuthSignin.
type PostApiV1AuthSigninJSONBody struct {
	Password *string `json:"password,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// CreateGameJSONRequestBody defines body for CreateGame for application/json ContentType.
type CreateGameJSONRequestBody = GameRequest

// UpdateGameJSONRequestBody defines body for UpdateGame for application/json ContentType.
type UpdateGameJSONRequestBody = GameRequest

// CreateResultJSONRequestBody defines body for CreateResult for application/json ContentType.
type CreateResultJSONRequestBody = ResultRequest

// CreateServiceJSONRequestBody defines body for CreateService for application/json ContentType.
type CreateServiceJSONRequestBody = ServiceRequest

// UpdateServiceJSONRequestBody defines body for UpdateService for application/json ContentType.
type UpdateServiceJSONRequestBody = ServiceRequest

// CreateTeamJSONRequestBody defines body for CreateTeam for application/json ContentType.
type CreateTeamJSONRequestBody = TeamRequest

// UpdateTeamJSONRequestBody defines body for UpdateTeam for application/json ContentType.
type UpdateTeamJSONRequestBody = TeamRequest

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = UserRequest

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = UserRequest

// PostApiV1AuthSigninJSONRequestBody defines body for PostApiV1AuthSignin for application/json ContentType.
type PostApiV1AuthSigninJSONRequestBody PostApiV1AuthSigninJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all games
	// (GET /api/games)
	ListGames(w http.ResponseWriter, r *http.Request)
	// Create a new game
	// (POST /api/games)
	CreateGame(w http.ResponseWriter, r *http.Request)
	// Delete a game
	// (DELETE /api/games/{id})
	DeleteGame(w http.ResponseWriter, r *http.Request, id int)
	// Get a game by ID
	// (GET /api/games/{id})
	GetGameById(w http.ResponseWriter, r *http.Request, id int)
	// Update a game
	// (PUT /api/games/{id})
	UpdateGame(w http.ResponseWriter, r *http.Request, id int)
	// List all results
	// (GET /api/results)
	ListResults(w http.ResponseWriter, r *http.Request)
	// Create a new result
	// (POST /api/results)
	CreateResult(w http.ResponseWriter, r *http.Request)
	// Get a result by ID
	// (GET /api/results/{id})
	GetResultById(w http.ResponseWriter, r *http.Request, id int)
	// List all services
	// (GET /api/services)
	ListServices(w http.ResponseWriter, r *http.Request)
	// Create a new service
	// (POST /api/services)
	CreateService(w http.ResponseWriter, r *http.Request)
	// Delete a service
	// (DELETE /api/services/{id})
	DeleteService(w http.ResponseWriter, r *http.Request, id int)
	// Get a service by ID
	// (GET /api/services/{id})
	GetServiceById(w http.ResponseWriter, r *http.Request, id int)
	// Update a service
	// (PUT /api/services/{id})
	UpdateService(w http.ResponseWriter, r *http.Request, id int)
	// List all teams
	// (GET /api/teams)
	ListTeams(w http.ResponseWriter, r *http.Request)
	// Create a new team
	// (POST /api/teams)
	CreateTeam(w http.ResponseWriter, r *http.Request)
	// Delete a team
	// (DELETE /api/teams/{id})
	DeleteTeam(w http.ResponseWriter, r *http.Request, id int)
	// Get a team by ID
	// (GET /api/teams/{id})
	GetTeamById(w http.ResponseWriter, r *http.Request, id int)
	// Update a team
	// (PUT /api/teams/{id})
	UpdateTeam(w http.ResponseWriter, r *http.Request, id int)
	// Retrieves a list of universities
	// (GET /api/universities)
	GetApiUniversities(w http.ResponseWriter, r *http.Request, params GetApiUniversitiesParams)
	// List all users
	// (GET /api/users)
	ListUsers(w http.ResponseWriter, r *http.Request)
	// Create a new user
	// (POST /api/users)
	CreateUser(w http.ResponseWriter, r *http.Request)
	// Delete a user
	// (DELETE /api/users/{id})
	DeleteUser(w http.ResponseWriter, r *http.Request, id int)
	// Get a user by ID
	// (GET /api/users/{id})
	GetUserById(w http.ResponseWriter, r *http.Request, id int)
	// Update a user
	// (PUT /api/users/{id})
	UpdateUser(w http.ResponseWriter, r *http.Request, id int)
	// Login user
	// (POST /api/v1/auth/signin)
	PostApiV1AuthSignin(w http.ResponseWriter, r *http.Request)
	// Logout user
	// (POST /api/v1/auth/signout)
	PostApiV1AuthSignout(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// List all games
// (GET /api/games)
func (_ Unimplemented) ListGames(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new game
// (POST /api/games)
func (_ Unimplemented) CreateGame(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a game
// (DELETE /api/games/{id})
func (_ Unimplemented) DeleteGame(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a game by ID
// (GET /api/games/{id})
func (_ Unimplemented) GetGameById(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a game
// (PUT /api/games/{id})
func (_ Unimplemented) UpdateGame(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List all results
// (GET /api/results)
func (_ Unimplemented) ListResults(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new result
// (POST /api/results)
func (_ Unimplemented) CreateResult(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a result by ID
// (GET /api/results/{id})
func (_ Unimplemented) GetResultById(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List all services
// (GET /api/services)
func (_ Unimplemented) ListServices(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new service
// (POST /api/services)
func (_ Unimplemented) CreateService(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a service
// (DELETE /api/services/{id})
func (_ Unimplemented) DeleteService(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a service by ID
// (GET /api/services/{id})
func (_ Unimplemented) GetServiceById(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a service
// (PUT /api/services/{id})
func (_ Unimplemented) UpdateService(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List all teams
// (GET /api/teams)
func (_ Unimplemented) ListTeams(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new team
// (POST /api/teams)
func (_ Unimplemented) CreateTeam(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a team
// (DELETE /api/teams/{id})
func (_ Unimplemented) DeleteTeam(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a team by ID
// (GET /api/teams/{id})
func (_ Unimplemented) GetTeamById(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a team
// (PUT /api/teams/{id})
func (_ Unimplemented) UpdateTeam(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Retrieves a list of universities
// (GET /api/universities)
func (_ Unimplemented) GetApiUniversities(w http.ResponseWriter, r *http.Request, params GetApiUniversitiesParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// List all users
// (GET /api/users)
func (_ Unimplemented) ListUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new user
// (POST /api/users)
func (_ Unimplemented) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a user
// (DELETE /api/users/{id})
func (_ Unimplemented) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a user by ID
// (GET /api/users/{id})
func (_ Unimplemented) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a user
// (PUT /api/users/{id})
func (_ Unimplemented) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Login user
// (POST /api/v1/auth/signin)
func (_ Unimplemented) PostApiV1AuthSignin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Logout user
// (POST /api/v1/auth/signout)
func (_ Unimplemented) PostApiV1AuthSignout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// ListGames operation middleware
func (siw *ServerInterfaceWrapper) ListGames(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListGames(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateGame operation middleware
func (siw *ServerInterfaceWrapper) CreateGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateGame(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteGame operation middleware
func (siw *ServerInterfaceWrapper) DeleteGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteGame(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetGameById operation middleware
func (siw *ServerInterfaceWrapper) GetGameById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGameById(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateGame operation middleware
func (siw *ServerInterfaceWrapper) UpdateGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateGame(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListResults operation middleware
func (siw *ServerInterfaceWrapper) ListResults(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListResults(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateResult operation middleware
func (siw *ServerInterfaceWrapper) CreateResult(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateResult(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetResultById operation middleware
func (siw *ServerInterfaceWrapper) GetResultById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetResultById(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListServices operation middleware
func (siw *ServerInterfaceWrapper) ListServices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListServices(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateService operation middleware
func (siw *ServerInterfaceWrapper) CreateService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateService(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteService operation middleware
func (siw *ServerInterfaceWrapper) DeleteService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteService(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetServiceById operation middleware
func (siw *ServerInterfaceWrapper) GetServiceById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetServiceById(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateService operation middleware
func (siw *ServerInterfaceWrapper) UpdateService(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateService(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListTeams operation middleware
func (siw *ServerInterfaceWrapper) ListTeams(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListTeams(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateTeam operation middleware
func (siw *ServerInterfaceWrapper) CreateTeam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTeam(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteTeam operation middleware
func (siw *ServerInterfaceWrapper) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTeam(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTeamById operation middleware
func (siw *ServerInterfaceWrapper) GetTeamById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTeamById(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateTeam operation middleware
func (siw *ServerInterfaceWrapper) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateTeam(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetApiUniversities operation middleware
func (siw *ServerInterfaceWrapper) GetApiUniversities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiUniversitiesParams

	// ------------- Optional query parameter "term" -------------

	err = runtime.BindQueryParameter("form", true, false, "term", r.URL.Query(), &params.Term)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "term", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetApiUniversities(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ListUsers operation middleware
func (siw *ServerInterfaceWrapper) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListUsers(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUser(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserById operation middleware
func (siw *ServerInterfaceWrapper) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserById(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", chi.URLParam(r, "id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateUser(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1AuthSignin operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1AuthSignin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1AuthSignin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostApiV1AuthSignout operation middleware
func (siw *ServerInterfaceWrapper) PostApiV1AuthSignout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiV1AuthSignout(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/games", wrapper.ListGames)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/games", wrapper.CreateGame)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/games/{id}", wrapper.DeleteGame)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/games/{id}", wrapper.GetGameById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/games/{id}", wrapper.UpdateGame)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/results", wrapper.ListResults)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/results", wrapper.CreateResult)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/results/{id}", wrapper.GetResultById)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/services", wrapper.ListServices)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/services", wrapper.CreateService)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/services/{id}", wrapper.DeleteService)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/services/{id}", wrapper.GetServiceById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/services/{id}", wrapper.UpdateService)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/teams", wrapper.ListTeams)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/teams", wrapper.CreateTeam)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/teams/{id}", wrapper.DeleteTeam)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/teams/{id}", wrapper.GetTeamById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/teams/{id}", wrapper.UpdateTeam)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/universities", wrapper.GetApiUniversities)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/users", wrapper.ListUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/users", wrapper.CreateUser)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/api/users/{id}", wrapper.DeleteUser)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/api/users/{id}", wrapper.GetUserById)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/api/users/{id}", wrapper.UpdateUser)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/auth/signin", wrapper.PostApiV1AuthSignin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api/v1/auth/signout", wrapper.PostApiV1AuthSignout)
	})

	return r
}
