package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
	pb "github.com/golanguzb70/grpc_client_server/client/genproto/post_service"
	grpclient "github.com/golanguzb70/grpc_client_server/client/grpcClient"
	"github.com/google/uuid"
)

func main() {
	client, err := grpclient.New()
	if err != nil {
		log.Printf("Error while connecting to gRPC servers")
		return
	}
	choice := 0

	for choice != 6 {
		fmt.Println("  I am client of Post Service   ")
		fmt.Println("1. Create Post")
		fmt.Println("2. Get Post")
		fmt.Println("3. List Post")
		fmt.Println("4. Update Post")
		fmt.Println("5. Delete Post")
		fmt.Println("6. Exit")
		fmt.Print(">>> ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			post, err := CreatePost(client)
			if err != nil {
				fmt.Println("Ooops something went wrong while creating post >>> ", err)
			} else {
				DisplayPost(post)
			}
		case 2:
			id := ""
			fmt.Print("Enter Id of post >>> ")
			fmt.Scan(&id)
			post, err := GetPost(client, id)
			if err != nil {
				fmt.Println("Ooops something went wrong while getting post >>> ", err)
			} else {
				DisplayPost(post)
			}
		case 3:
			page := int32(1)
			limit := int32(2)

			for {
				posts, err := FindPosts(client, page, limit)
				if err != nil {
					fmt.Println("Ooops something went wrong while getting posts >>> ", err)
				} else {
					for _, post := range posts.GetItems() {
						DisplayPost(post)
					}
				}
				if len(posts.Items) < int(limit) {
					break
				}

				fmt.Println("press 1 to see the next page")
				ch := "0"
				fmt.Scan(&ch)

				if ch != "1" {
					break
				}
				page++
			}
		case 4:
			id := ""
			fmt.Print("Enter Id of post >>> ")
			fmt.Scan(&id)

			post, err := UpdatePost(client, id)
			if err != nil {
				fmt.Println("Ooops something went wrong while updating post >>> ", err)
			} else {
				DisplayPost(post)
			}
		case 5:
			id := ""
			fmt.Print("Enter Id of post >>> ")
			fmt.Scan(&id)
			err := DeletePost(client, id)
			if err != nil {
				fmt.Println("Ooops something went wrong while updating post >>> ", err)
			} else {
				fmt.Println("Post successfully deleted")
			}
		}
	}

}

func CreatePost(client grpclient.ServiceManager) (*pb.Post, error) {
	post := &pb.CreateRequest{
		Id:      uuid.NewString(),
		Slug:    faker.Username(),
		Title:   faker.Word(),
		Body:    faker.Sentence(),
		OwnerId: uuid.NewString(),
	}

	return client.PostServiceClient().Create(context.Background(), post)
}

func GetPost(client grpclient.ServiceManager, id string) (*pb.Post, error) {
	return client.PostServiceClient().GetByKey(context.Background(), &pb.KeyRequest{Id: id})
}

func FindPosts(client grpclient.ServiceManager, page, limit int32) (*pb.Posts, error) {
	return client.PostServiceClient().Find(context.Background(), &pb.Filter{Page: page, Limit: limit})
}

func UpdatePost(client grpclient.ServiceManager, id string) (*pb.Post, error) {
	post := &pb.UpdateRequest{
		Id:    id,
		Slug:  faker.Username(),
		Title: faker.Word(),
		Body:  faker.Sentence(),
	}

	return client.PostServiceClient().Update(context.Background(), post)
}

func DeletePost(client grpclient.ServiceManager, id string) error {
	_, err := client.PostServiceClient().Delete(context.Background(), &pb.KeyRequest{Id: id})

	return err
}

func DisplayPost(post *pb.Post) {
	postByte, err := json.MarshalIndent(post, "", "\t")

	if err != nil {
		fmt.Println("Output: ", err)
		return
	}

	fmt.Println("Output: \n", string(postByte))
}
