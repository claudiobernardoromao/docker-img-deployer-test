package main

import (
	"io"
	"io/ioutil"
	"os"
	"log"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)

func main() {
	fmt.Println(" teste do console")
	logTo("init.out")
	log.Println(" Create context")
	ctx := context.Background()
	log.Println(" Create NewEnvClient ")
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	log.Println(" Client Created")
	ref := "nginx:latest"

	err = imagePull(cli,ref)
}
func imagePull(cli *client.Client, ref string) error {
	log.Printf("Pulling %q from the registry...\n", ref)
	resp, err := cli.ImagePull(context.Background(), ref, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer resp.Close()
	if _, err = io.Copy(ioutil.Discard, resp); err != nil {
		return err
	}
	log.Println("Image pull complete")
	return nil
}

func logTo(fileName string) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
	return f
}
