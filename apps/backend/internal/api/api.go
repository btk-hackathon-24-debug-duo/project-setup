package api

import (
	"database/sql"
	"net/http"

	"github.com/btk-hackathon-24-debug-duo/project-setup/internal/middleware"
	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	db           *sql.DB
	mongoClient  *mongo.Collection
	geminiClient *genai.GenerativeModel
}

func NewRouter(db *sql.DB, mongo *mongo.Collection, gemini *genai.GenerativeModel) *Router {
	return &Router{
		db:           db,
		mongoClient:  mongo,
		geminiClient: gemini,
	}
}

func (r *Router) NewRouter() *mux.Router {
	h := NewUserHandlers(r.db)
	c := NewChatHandlers(r.mongoClient, r.geminiClient, r.db)

	router := mux.NewRouter()

	router.Use(middleware.CorsMiddleware)

	router.HandleFunc("/user/login", h.LoginHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/user/register", h.RegisterHandler).Methods(http.MethodPost, http.MethodOptions)

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.EnsureValidToken)

	protected.HandleFunc("/chat/message", c.SendMessageHandler).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/chat/message", c.GetMessages).Methods(http.MethodGet, http.MethodOptions)

	protected.HandleFunc("/chat/firstmessage", c.SendFirstMessageHandler).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/chat/name", c.UpdateChatNameHandler).Methods(http.MethodPut, http.MethodOptions)

	protected.HandleFunc("/chat", c.NewChat).Methods(http.MethodPost, http.MethodOptions)
	protected.HandleFunc("/chat", c.GetChats).Methods(http.MethodGet, http.MethodOptions)

	protected.HandleFunc("/user/update", h.UpdateUserHandler).Methods(http.MethodPut, http.MethodOptions)

	return router
}
