package users

import (
	"net/http"
	"strconv"

	"github.com/khaihoan0805/bookstore_users-api/domain/users"

	"github.com/khaihoan0805/bookstore_users-api/utils/errors"

	"github.com/khaihoan0805/bookstore_users-api/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/* be replaced by following code to check the error of the context went server receive a request
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle the JSON error
		fmt.Println(err.Error())
		return
	}
	*/
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: handle the JSON error
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: Handle User Creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		//TODO: Handle User Creation error
		return
	}
	c.JSON(http.StatusCreated, user)

}
