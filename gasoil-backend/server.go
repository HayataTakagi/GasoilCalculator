package server

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"firebase.google.com/go/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/HayataTakagi/GasoilCalculator/controller"
	db2 "github.com/HayataTakagi/GasoilCalculator/db"
	"github.com/HayataTakagi/GasoilCalculator/firebase"
	"github.com/HayataTakagi/GasoilCalculator/middleware"
	"github.com/HayataTakagi/GasoilCalculator/sample"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {
	authClient, err := firebase.InitAuthClient()
	if err != nil {
		log.Fatalf("failed init auth client. %s", err)
	}
	s.authClient = authClient

	db := db2.NewDB(datasource)
	dbcon, err := db.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	authMiddleware := middleware.NewAuth(s.authClient, s.db)
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	authChain := commonChain.Append(
		authMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))
	r.Methods(http.MethodGet).Path("/private").Handler(authChain.Then(sample.NewPrivateHandler(s.db)))

	gasOilController := controller.NewGaOil()
	r.Methods(http.MethodGet).Path("/gasoil").Handler(commonChain.Then(AppHandler{gasOilController.Index}))

	carController := controller.NewCar(s.db)
	r.Methods(http.MethodGet).Path("/car").Handler(commonChain.Then(AppHandler{carController.Index}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
