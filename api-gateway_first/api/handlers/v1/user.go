package v1

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	_ "github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/google/uuid"

	//"github.com/gin-gonic/gin/internal/json"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/auth"
	pb "github.com/najimovmashhurbek/Project_Api/api-gateway_first/genproto"
	l "github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
	gomail "gopkg.in/mail.v2"
)

type User struct {
	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	FirstName    string    `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName"`
	LastName     string    `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName"`
	Bio          string    `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio"`
	PhoneNumbers []string  `protobuf:"bytes,5,rep,name=phoneNumbers,proto3" json:"phoneNumbers"`
	Status       string    `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	CreatedAt    string    `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt"`
	UpdateAt     string    `protobuf:"bytes,8,opt,name=updateAt,proto3" json:"updateAt"`
	DeletedAt    string    `protobuf:"bytes,9,opt,name=deletedAt,proto3" json:"deletedAt"`
	Username     string    `protobuf:"bytes,10,opt,name=username,proto3" json:"username"`
	Email        string    `protobuf:"bytes,11,opt,name=email,proto3" json:"email"`
	Password     string    `protobuf:"bytes,12,opt,name=password,proto3" json:"password"`
	Adress       []*Adress `protobuf:"bytes,13,rep,name=adress,proto3" json:"adress"`
	Post         []*Post   `protobuf:"bytes,14,rep,name=post,proto3" json:"post"`
	Code         string    `protobuf:"bytes,15,opt,name=code,proto3" json:"code"`
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
type Emailver struct {
	Email string `json:"Email"`
	Code  string `json:"Code"`
}
type RegisterResponse struct {
	UserID       string
	Accesstoken  string
	Refreshtoken string
}

// CreateUser creates user
// @Summary Create user summary
// @Description This api is using for creating new user
// @Tags user
// @Accept json
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
// @Accept json
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
// @Accept json
// @Produce json
// @Param limit query int true "limit"
// @Param page query int true "page"
// @Success 200 {string} User
// @Router /v1/users [get]
func (h *handlerV1) ListUsers(c *gin.Context) {
	p := c.Query("page")
	l := c.Query("limit")

	CheckClaims(h, c)

	page, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	limit, err := strconv.ParseInt(l, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to parsing limit or page to conv")
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().ListUsers(
		ctx, &pb.GetUsersReq{
			Page:  page,
			Limit: limit,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to getting user listing")
		return
	}
	c.JSON(http.StatusOK, response)
}

// UpdateUser updates user by id
// @Summary Update user summary
// @Description This api is using for updating new user
// @Tags user
// @Accept json
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

// RegisterUser register user
// @Summary Register user summary
// @Description This api is using for registering new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body User  true "user_body"
// @Success 200 {string} Succes
// @Router /v1/users/register [post]
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
	err = verifyPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		h.log.Error("your password doesn't respond to requests", l.Error(err))
		return
	}

	//hashing password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), len(body.Password))

	body.Password = string(hashedPassword)
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	exists, err := h.serviceManager.UserService().CheckUniquess(ctx, &pb.CheckUniqReq{
		Field: "email",
		Value: body.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check email uniquess", l.Error(err))
		return
	}
	if exists.IsExist {
		c.JSON(http.StatusConflict, gin.H{
			"error": "this email already in use, please use another email",
		})
		h.log.Error("failed to check email uniquess",
			l.Error(err))
		return
	}
	exists, err = h.serviceManager.UserService().CheckUniquess(ctx, &pb.CheckUniqReq{
		Field: "username",
		Value: body.Username,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check username uniquess", l.Error(err))
		return
	}
	if exists.IsExist {
		c.JSON(http.StatusConflict, gin.H{
			"eror": "this usrname already in use, please use another username",
		})
		h.log.Error("failed to check username uniquess", l.Error(err))
		return
	}
	//code generate
	min := 99999
	max := 1000000
	rand.Seed(time.Now().UnixNano())
	gen := rand.Intn((max - min) + min)
	code := strconv.Itoa(gen)

	body.Code = code

	bodyByte, err := json.Marshal(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to set to redis", l.Error(err))
		return
	}
	//writing redis
	err = h.redisStorage.SetWithTTL(body.Email, string(bodyByte), int64(time.Second*150))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to set redis", l.Error(err))
		return
	}
	SendEmail(body.Email, code)
}

// VerifyUser verify user
// @Description This api using for verifying registered user
// @Tags user
// @Accept json
// @Produce json
// @Param user body Emailver true "user body"
// @Success 200 {string} success
// @Router /v1/users/verfication [post]
func (h *handlerV1) VerifyUser(c *gin.Context) {
	var (
		dataemail   Emailver
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&dataemail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	dataemail.Email = strings.TrimSpace(dataemail.Email)
	dataemail.Email = strings.ToLower(dataemail.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	getRedis, err := redis.String(h.redisStorage.Get(dataemail.Email))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to getting redis for write db", l.Error(err))
		return
	}
	var redisBody *pb.User
	_ = json.Unmarshal([]byte(getRedis), &redisBody)

	if dataemail.Code == redisBody.Code {
		_, err := h.serviceManager.UserService().CreateUser(ctx, redisBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed to writting db", l.Error(err))
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("your code is wrong", l.Error(err))
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while generating uuid",
		})
		h.log.Error("error generate new uuid", l.Error(err))
		return
	}
	h.jwtHandler = auth.JwtHandler{
		Sub:  id.String(),
		Iss:  "client",
		Role: "authorized",
		Log:  h.log,
	}
	access, refresh, err := h.jwtHandler.GenerateJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error while generating jwt",
		})
		h.log.Error("error generate new jwt tokens", l.Error(err))
		return

	}

	c.JSON(http.StatusOK, &RegisterResponse{
		UserID:       id.String(),
		Accesstoken:  access,
		Refreshtoken: refresh,
	})
}

// Login login user
// @Description This api using for logging registered user
// @Tags user
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Param password path string true "Password"
// @Succes 200 {string} LoginResponse
// @Router /v1/users/login/{email}/{password} [post]
func (h *handlerV1) Login(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	email := c.Param("email")
	password := c.Param("password")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	userData, err := h.serviceManager.UserService().LoginUser(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to getting datas", l.Error(err))
		return
	}
	userData.Password = ""
	c.JSON(http.StatusOK, userData)
}

func SendEmail(email, code string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", "testapigomail@gmail.com")

	// Set E-Mail receivers
	m.SetHeader("To", email)

	// Set E-Mail subject
	m.SetHeader("code:", "dfsdfdsf")

	m.SetBody("text/plain", code)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "testapigomail@gmail.com", "cpebajsbmuddenig")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
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
