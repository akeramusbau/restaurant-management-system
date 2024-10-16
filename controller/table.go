package controller

import (
	"fmt"
	"net/http"

	"github.com/akeramusbau/restaurant-management-system/services"
	"github.com/gin-gonic/gin"
)

type TableController struct{}

type createTableInput struct {
	Username string `json:"username" binding:"required"`
	RoomID   int    `json:"room_id" binding:"required"`
}

func (controller TableController) JoinTable(c *gin.Context) {
	// * Input Validation
	var input createTableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var response = ErrorResponse{
			Msg: "Validation Error",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	// * PLayer Join Room
	/*isPlayerJoined, err := services.MatchService.FindUserInRoom(input.Username, uint(input.RoomID))
	if  err != nil {
		var response = ErrorResponse{
			Msg: "Failed to check player",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	fmt.Println(isPlayerJoined)*/
	match, err := services.TableService.PlayerJoinRoom(input.Username, uint(input.RoomID))
	if err != nil {
		var response = ErrorResponse{
			Msg: "Failed to Create Player",
			Err: err,
		}
		c.AbortWithStatusJSON(http.StatusOK, response)
		return
	}

	var responseString = fmt.Sprintf("Successfully joined room %s", string(match.RoomID))
	var response = SuccessResponse{
		Status: true,
		Msg:    responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}
