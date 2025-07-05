package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/suhas-developer07/Authentification_in_Go/config"
)

var DB *pgxpool.Pool

func InitDB (){
	var err error

	DB,err =pgxpool.New(context.Background(),config.Conifg.POSTGERS_URL)

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