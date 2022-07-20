package handle

import (
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Add(c *gin.Context, s *machinery.Server) {
	uid := uuid.New().String()

	signature := &tasks.Signature{
		UUID: uid,
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	asyncResult, err := s.SendTask(signature)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"add": err, "uuid": uid})
	fmt.Println(asyncResult)
}

func AddChain(c *gin.Context, s *machinery.Server) {
	signature1 := tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}

	signature2 := tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 5,
			},
			{
				Type:  "int64",
				Value: 5,
			},
		},
	}

	signature3 := tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 4,
			},
			{
				Type:  "int64",
				Value: 3,
			},
		},
	}

	chain, _ := tasks.NewChain(&signature1, &signature2, &signature3)
	chainAsyncResult, err := s.SendChain(chain)
	if err != nil {
		// failed to send the chain
		// do something with the error
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{"add": err})
	fmt.Println(chainAsyncResult)
}
