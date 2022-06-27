package v1

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	pb "github.com/najimovmashhurbek/Project_Api/api-gateway_first/genproto"
	l "github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

type User struct {
	//Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	FirstName    string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"first_name"`
	LastName     string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"last_name"`
	Bio          string   `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio"`
	PhoneNumbers []string `protobuf:"bytes,7,rep,name=phoneNumbers,proto3" json:"phone_numbers"`
	Status       string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	CreatedAt    string   `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"created_at"`
	UpdateAt     string   `protobuf:"bytes,10,opt,name=updateAt,proto3" json:"update_at"`
	DeletedAt    string   `protobuf:"bytes,11,opt,name=deletedAt,proto3" json:"deleted_at"`
	Adress       []Adress `protobuf:"bytes,12,rep,name=adress,proto3" json:"adress"`
	Post         []Post   `protobuf:"bytes,13,rep,name=post,proto3" json:"post"`
}
type Adress struct {
	//Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	//UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Country     string `protobuf:"bytes,3,opt,name=country,proto3" json:"country"`
	City        string `protobuf:"bytes,4,opt,name=city,proto3" json:"city"`
	District    string `protobuf:"bytes,5,opt,name=district,proto3" json:"district"`
	PostalCodes int64  `protobuf:"varint,6,opt,name=postalCodes,proto3" json:"postal_codes"`
}
type Post struct {
	//Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	//UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Medias []Media `protobuf:"bytes,5,rep,name=medias,proto3" json:"medias"`
}
type Media struct {
	//Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Link string `protobuf:"bytes,3,opt,name=link,proto3" json:"link"`
}

// CreateUser creates user
// @Summary Create user summary
// @Description This api is using for creating new user
// @Tags user
// Accept json
// @Produce json
// @Success 200 {string} Succes
// @Param user body User  true "user body"
// @Router /v1/users [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	fmt.Println(&body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUser gets user by id
// @Summary Get user summary
// @Description This api is using for getting user by id
// @Tags user
// Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} User
// @Router /v1/users/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetAllUser(
		ctx, &pb.GetAllById{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListUsers gets users list
// @Summary Get user summary
// @Description This api is using for getting users list
// @Tags user
// Accept json
// @Produce json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {string} User
// @Router /v1/users [get]
func (h *handlerV1) ListUsers(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().ListUsers(
		ctx, &pb.GetUsersReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by id
// @Summary Update user summary
// @Description This api is using for updating new user
// @Tags user
// Accept json
// @Produce json
// @Success 200 {string} Succes
// @Param id path string true "User ID"
// @Param user body User  true "user body"
// @Router /v1/users/{id} [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes user by id
// @Summary Delete user summary
// @Description This api is using for deleting user
// @Tags user
// Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} Succes!
// @Router /v1/users/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().DeleteUser(
		ctx, &pb.DeleteById{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

//reg

func (h *handlerV1) Register(c *gin.Context) {
	var (
		body        pb.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	//check password

	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	exists, err := h.serviceManager.UserService().CheckUniquess(ctx, &pb.CheckUniqReq{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check email uniquess", l.Error(err))
		return
	}

}




func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool
	const minPassLength = 8
	const maxPassLength = 32
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}
	if !lowercasePresent {
		appendError("lowercase letter missing")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing")
	}
	if !numberPresent {
		appendError("atleast one numeric character required")
	}
	if !specialCharPresent {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(fmt.Sprintf("password length must be between %d to %d characters long", minPassLength, maxPassLength))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}
