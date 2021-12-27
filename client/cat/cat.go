package cat

import (
	"context"
	"fmt"
	cpb "grpc/proto/category"
	"io"

	"google.golang.org/grpc"
)


type Client struct{
	CatProtoBuf cpb.CategoryServiceClient
}

func NewClient(conn grpc.ClientConnInterface) Client{
	return Client{
		CatProtoBuf: cpb.NewCategoryServiceClient(conn),
	}
}

func (c *Client) GetCategory(id int64) (*cpb.GetCategoryResponse, error) {
	return c.CatProtoBuf.GetCategory(context.Background(), &cpb.GetCategoryRequest{
		ID: id,
	})
}

func (c *Client) GetCategories() error {
	stream, err := c.CatProtoBuf.GetCategories(context.Background(), &cpb.GetCategoriesRequest{})

	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil{
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println(res)
	}
}

func (c *Client) GetCat() error {
	stream, err := c.CatProtoBuf.GetCategories(context.Background(), &cpb.GetCategoriesRequest{})

	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if err != nil{
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println(res)
	}
}
