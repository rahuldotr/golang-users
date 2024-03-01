package handlers

import (
	"project3/dbhandler/models"
	"project3/objects"
	"project3/utils"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
)

func Status(context *gin.Context) { // Test endpoint for checking server status
	utils.SendSuccessResponse(context, 200, "success", "Oohhh!! Hell yeaaaah!! The server is up!!", nil)
}

func CreateUser(context *gin.Context) {
	var userDetails objects.User

	if err := context.BindJSON(&userDetails); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", err.Error())
		return
	}
	if hasAnyValidationErrors := hasValidationErrors(context, userDetails); hasAnyValidationErrors {
		return
	}

	getEmailParams := map[string]interface{}{ //create parameters for edgedb
		"email": userDetails.Email,
	}
	if isEmailExist, _ := models.GetUserByEmail(context, getEmailParams); isEmailExist {
		utils.SendErrorResponse(context, 409, "error", "", "email already exists!")
		return
	}

	params := map[string]interface{}{ //create parameters for edgedb
		"name":             userDetails.Name,
		"email":            userDetails.Email,
		"company":          userDetails.Company,
		"bio":              userDetails.Bio,
		"location":         userDetails.Location,
		"twitter_username": userDetails.TwitterUsername,
	}
	if result, createErr := models.CreateUser(context, params); createErr != nil {
		utils.SendErrorResponse(context, 400, "error", "", "error occured while creating user")
		return
	} else {
		var res = &objects.CreateUser{
			Id: result.Id,
		}
		utils.SendSuccessResponse(context, 200, "success", "", res)
	}
}

func GetAllUsers(context *gin.Context) {

	if result, err := models.GetAllUsers(context); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", "error occured while getting users")
		return
	} else {
		utils.SendSuccessResponse(context, 200, "success", "", result)
	}
}

func GetIndividualUser(context *gin.Context) {

	uuidObj, uuidErr := edgedb.ParseUUID(context.Param("user-id")) // parse uuid from the params into edgedb object
	if uuidErr != nil {
		utils.SendErrorResponse(context, 400, "error", "", utils.UuidParseError)
		return
	}
	params := map[string]interface{}{
		"id": uuidObj, //create parameters for edgedb
	}

	if result, err := models.GetIndividualUser(context, params); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", "no data exist for the user id")
		return
	} else {
		utils.SendSuccessResponse(context, 200, "success", "", result)
	}
}

func UpdateUser(context *gin.Context) {

	var user objects.User

	if err := context.BindJSON(&user); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", err.Error())
		return
	}
	if hasAnyValidationErrors := hasValidationErrors(context, user); hasAnyValidationErrors {
		return
	}

	uuidObj, uuidErr := edgedb.ParseUUID(context.Param("user-id"))
	if uuidErr != nil {
		utils.SendErrorResponse(context, 400, "error", "", utils.UuidParseError)
		return
	}
	uuidParams := map[string]interface{}{
		"id": uuidObj, //create parameters for edgedb
	}
	if _, err := models.GetIndividualUser(context, uuidParams); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", "user id doesn't exist")
		return
	}

	getEmailParams := map[string]interface{}{ //create parameters for edgedb
		"id":    uuidObj,
		"email": user.Email,
	}

	if isEmailExist, _ := models.CheckEmailAssignedToOtherUser(context, getEmailParams); isEmailExist {
		utils.SendErrorResponse(context, 409, "error", "", "email already taken by another user!")
		return
	}

	params := map[string]interface{}{ //create parameters for edgedb
		"id":               uuidObj,
		"name":             user.Name,
		"email":            user.Email,
		"company":          user.Company,
		"bio":              user.Bio,
		"location":         user.Location,
		"twitter_username": user.TwitterUsername,
	}
	if _, updateErr := models.UpdateUser(context, params); updateErr != nil {
		utils.SendErrorResponse(context, 400, "error", "", "error occured while updating user details")
		return
	}
	if result, getErr := models.GetIndividualUser(context, params); getErr != nil {
		utils.SendErrorResponse(context, 400, "error", "", "error occured while getting updated user details")
		return
	} else {
		utils.SendSuccessResponse(context, 200, "success", "", result)
	}
}

func DeleteUser(context *gin.Context) {

	uuidObj, uuidErr := edgedb.ParseUUID(context.Param("user-id"))
	if uuidErr != nil {
		utils.SendErrorResponse(context, 400, "error", "", utils.UuidParseError)
		return
	}

	params := map[string]interface{}{ //create parameters for edgedb
		"id": uuidObj,
	}

	if _, err := models.GetIndividualUser(context, params); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", "user id doesn't exist")
		return
	}

	if _, err := models.DeleteUser(context, params); err != nil {
		utils.SendErrorResponse(context, 400, "error", "", "error occured while deleting user")
		return
	} else {
		utils.SendSuccessResponse(context, 200, "success", "user deleted successfuly", nil)
	}
}

func hasValidationErrors(context *gin.Context, user objects.User) bool {

	hasValidationError := false
	if name, isName := user.Name.Get(); !isName || name == "" {
		utils.SendErrorResponse(context, 400, "error", "", "name is required")
		hasValidationError = true
	}
	if email, isEmail := user.Email.Get(); !isEmail || email == "" {
		utils.SendErrorResponse(context, 400, "error", "", "email is required")
		hasValidationError = true
	} else {
		if isValid := utils.IsValidEmail(email); !isValid {
			utils.SendErrorResponse(context, 400, "error", "email should be in the format 'test@test.com'", "email is not valid")
			hasValidationError = true
		}
	}
	return hasValidationError
}
