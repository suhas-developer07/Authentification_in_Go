package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)



var DB *pgx.Conn

func InitDB (){
	var err error

	DB,err =pgx.Connect(context.Background(),"url")

	if err!=nil {
		fmt.Fprintf(os.Stderr,"Unable to connect to database : %v\n",err);
	}
	err = DB.Ping(context.Background())

	if err!=nil {
		log.Fatalf("error connecting database")
		os.Exit(1)
	}
	log.Printf("connected with database")
}