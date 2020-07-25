package users

import (
	"net/http"
	"strconv"

	"github.com/khaihoan0805/bookstore_users-api/domain/users"

	"github.com/khaihoan0805/bookstore_users-api/utils/errors"

	"github.com/khaihoan0805/bookstore_users-api/services"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
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

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		//TODO: Handle User Creation error
		return
	}
	c.JSON(http.StatusCreated, user)
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: handle the JSON error
		restErr := errors.NewBadRequestError("invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, saveErr := services.UpdateUser(isPartial, user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: Handle User Creation error
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	err := services.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		//TODO: Handle User Creation error
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
