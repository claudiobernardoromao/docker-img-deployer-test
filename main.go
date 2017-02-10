package main

import (
//	"io"
	"os"
	"log"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)

func main() {

	logTo("init.out")
	log.Println(" Create context")
	ctx := context.Background()
	log.Println(" Create NewEnvClient ")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	log.Println(" Client Created")
	
	_, err = cli.ImagePull(ctx, "nginx", types.ImagePullOptions{})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	
}
func logTo(fileName string) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
	return f
}
