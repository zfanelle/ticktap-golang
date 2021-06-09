package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	pb "github.com/zfanelle/ticktap-golang/internal/app/mainservice/config/protos"
	"google.golang.org/grpc"
)

type AppConfig struct {
	DB         *sqlx.DB
	Router     *mux.Router
	GrpcClient *pb.TicketingServiceClient
}

func (appConfig *AppConfig) Configure() {

	appConfig.configureDatabase()
	appConfig.configureTicketingService()
}

func (appConfig *AppConfig) configureDatabase() {

	var db *sqlx.DB

	db, err := sqlx.Connect("mysql", "root:mysecretpassword@(localhost:3306)/ticktap")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()

	appConfig.DB = db

}

func (appConfig *AppConfig) configureTicketingService() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	//defer conn.Close()

	c := pb.NewTicketingServiceClient(conn)

	appConfig.GrpcClient = &c

	log.Printf("Response from server: %s")

	// response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})
	// if err != nil {
	// 	log.Fatalf("Error when calling SayHello: %s", err)
	// }
	// log.Printf("Response from server: %s", response.Body)

}
