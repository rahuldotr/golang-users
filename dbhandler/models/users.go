package models

import (
	"fmt"
	"project3/dbhandler"
	"project3/objects"
	"project3/queries"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) ([]objects.User, error) {
	var result []objects.User

	err := dbhandler.DbClient.Query(context, queries.GetAllUsersQuery(), &result)
	fmt.Println(err, ">>>>>> from 'GetAllUsers' function")
	return result, err
}

func GetIndividualUser(context *gin.Context, params map[string]interface{}) (objects.User, error) {
	var result objects.User

	err := dbhandler.DbClient.QuerySingle(context, queries.GetIndividualUserQuery(), &result, params)
	fmt.Println(err, ">>>>>> from 'GetIndividualUser' function")
	return result, err
}

func CreateUser(context *gin.Context, params map[string]interface{}) (objects.CreateUser, error) {
	var result objects.CreateUser

	err := dbhandler.DbClient.QuerySingle(context, queries.GetInsertUserQuery(), &result, params)
	fmt.Println(err, ">>>>>> from 'CreateUser' function")
	return result, err
}

func UpdateUser(context *gin.Context, params map[string]interface{}) (objects.User, error) {
	var result objects.User

	err := dbhandler.DbClient.QuerySingle(context, queries.UpdateUserQuery(), &result, params)
	fmt.Println(err, ">>>>>> from 'UpdateUser' function")
	return result, err
}

func DeleteUser(context *gin.Context, params map[string]interface{}) (objects.User, error) {
	var result objects.User

	err := dbhandler.DbClient.QuerySingle(context, queries.DeleteUserQuery(), &result, params)
	fmt.Println(err, ">>>>>> from 'DeleteUser' function")
	return result, err
}

func GetUserByEmail(context *gin.Context, params map[string]interface{}) (bool, error) {
	var result int64

	err := dbhandler.DbClient.QuerySingle(context, queries.GetUserWithEmailCountQuery(), &result, params)
	fmt.Println(err, ">>>>>> from 'GetUserByEmail' function")
	var status bool
	if result > 0 {
		status = true
	}
	return status, err
}

func CheckEmailAssignedToOtherUser(context *gin.Context, params map[string]interface{}) (bool, error) {
	var result int64

	err := dbhandler.DbClient.QuerySingle(context, queries.CheckEmailAssignedToOtherUser(), &result, params)
	fmt.Println(err, ">>>>>> from 'CheckEmailAssignedToOtherUser' function")
	var status bool
	if result > 0 {
		status = true
	}
	return status, err
}
