package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/piyushgarg878/students-api/internal/config"
)

func main(){
	//first load config 
	cfg:=config.MustLoad()
	//add router and route
	router:=http.NewServeMux()
	router.HandleFunc("GET /",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("Welcome to students api"))
	})

	server:=http.Server{
		Addr: cfg.Addr,
		Handler:router,
	}
	done:=make(chan os.Signal,1)
	
	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
	slog.Info("server started",slog.String("address",cfg.Addr))
	go func(){
		err:=server.ListenAndServe()
			if err!=nil{
				log.Fatal("error",err)
	}
	}()
	
	<-done

	slog.Info("Shutting down the server")
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()


	if err:=server.Shutdown(ctx);err!=nil{
		slog.Error("failed to shiut down server",slog.String("error",err.Error()))
	}

	slog.Info("shut down successfully")












	//setup database
	//setup http server
}
