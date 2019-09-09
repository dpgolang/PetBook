package main

import (
	"github.com/dpgolang/PetBook/gomigrations"
	"github.com/dpgolang/PetBook/pkg/authentication"
	"github.com/dpgolang/PetBook/pkg/controllers"
	"github.com/dpgolang/PetBook/pkg/driver"
	"github.com/dpgolang/PetBook/pkg/logger"
	_ "github.com/dpgolang/PetBook/pkg/logger"
	"github.com/dpgolang/PetBook/pkg/models"
	"github.com/dpgolang/PetBook/pkg/models/forum"
	"github.com/dpgolang/PetBook/pkg/models/search"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	db := driver.ConnectDB()
	err := gomigrations.Migrate(db)
	if err != nil {
		logger.FatalError(err, "Migration failed.\n")
	}

	router := mux.NewRouter()

	storeUser := models.UserStore{DB: db}
	storePet := models.PetStore{DB: db}
	storeRefreshToken := models.RefreshTokenStore{DB: db}
	storeForum := forum.ForumStore{DB: db}
	storeSearch := search.SearchStore{DB: db}
	storeBlog := models.BlogStore{DB: db}
	storeChat := models.ChatStore{DB: db}

	controller := controllers.Controller{
		PetStore:          &storePet,
		UserStore:         &storeUser,
		ForumStore:        &storeForum,
		SearchStore:       &storeSearch,
		RefreshTokenStore: &storeRefreshToken,
		BlogStore:         &storeBlog,
		ChatStore:         &storeChat,
	}

	router.HandleFunc("/register", controller.RegisterPostHandler()).Methods("POST")
	router.HandleFunc("/register", controller.RegisterGetHandler()).Methods("GET")

	router.HandleFunc("/login", controller.LoginPostHandler()).Methods("POST")
	router.HandleFunc("/login", controller.LoginGetHandler()).Methods("GET")
	router.HandleFunc("/logout", controller.LogoutGetHandler()).Methods("GET")

	subrouter := router.PathPrefix("/").Subrouter()
	subrouter.Use(authentication.ValidateTokenMiddleware(&storeRefreshToken, &storeUser))

	subrouter.HandleFunc("/mypage", controller.MyPageGetHandler()).Methods("GET")
	subrouter.HandleFunc("/petcabinet", controller.PetPostHandler()).Methods("POST")
	subrouter.HandleFunc("/petcabinet", controller.PetGetHandler()).Methods("GET")
	subrouter.HandleFunc("/search", controller.ViewSearchHandler()).Queries("section", "{section}").Methods("GET")
	subrouter.HandleFunc("/search", controller.RedirectSearchHandler()).Methods("GET")

	subrouter.HandleFunc("/forum", controller.TopicsGetHandler()).Methods("GET")
	subrouter.HandleFunc("/forum", controller.TopicsPostHandler()).Methods("POST")
	subrouter.HandleFunc("/forum/topic/{topicID}/comments", controller.CommentsGetHandler()).Methods("GET")
	subrouter.HandleFunc("/forum/topic/{topicID}/comments", controller.CommentsPostHandler()).Methods("POST")
	subrouter.HandleFunc("/forum/topic/{topicID}/comments/{commentID}/ratings", controller.CommentsLikeHandler()).Methods("POST")

	subrouter.HandleFunc("/chats", controller.ChatsGetHandler()).Methods("GET")
	subrouter.HandleFunc("/chats/{id}", controller.DeleteChatHandler()).Methods("POST") //does not work with method DELETE with overriding with js too
	subrouter.HandleFunc("/chats/{id}", controller.HandleChatConnectionGET()).Methods("GET")
	subrouter.HandleFunc("/chats/{id}/search/{date}", controller.HandleChatSearchConnection()).Methods("GET")
	subrouter.HandleFunc("/ws/{id}", controller.HandleChatConnection())

	//subrouter.HandleFunc("/search", controller.ViewSearchHandler()).Methods("GET")
	subrouter.HandleFunc("/", controller.MyPageGetHandler()).Methods("GET")

	subrouter.HandleFunc("/process", controller.CreateBlogHandler).Methods("Post")
	subrouter.HandleFunc("/delete", controller.DeleteBlogHandler)
	subrouter.HandleFunc("/upload", controllers.UploadFile)
	subrouter.HandleFunc("/edit", controller.EditHandler).Methods("GET")
	subrouter.HandleFunc("/edit", controller.UpdateHandler).Methods("POST")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./web/static/"))))

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	if err := http.ListenAndServe(":8080", loggedRouter); err != nil {
		logger.FatalError(err, "Error occurred, while trying to listen and serve a server")
	}
}
