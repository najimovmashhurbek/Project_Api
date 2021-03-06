package postgres

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	pb "github.com/najimovmashhurbek/Project_Api/user-service.first/genproto"
	"golang.org/x/crypto/bcrypt"
	//"google.golang.org/grpc/internal/status"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *pb.User) (*pb.User, error) {
	userr := pb.User{}
	query := "insert into users (id,firstName,lastName,bio,phoneNumbers,createdAt,status,email,username,password) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id,firstName,lastName,bio,createdAt,status,phoneNumbers,email,username"
	time := time.Now()
	id := uuid.New()
	err := r.db.QueryRow(query, id, user.FirstName, user.LastName, user.Bio, pq.Array(user.PhoneNumbers), time, user.Status, user.Email, user.Username, user.Password).Scan(
		&userr.Id,
		&userr.FirstName,
		&userr.LastName,
		&userr.Bio,
		&userr.CreatedAt,
		&userr.Status,
		pq.Array(&userr.PhoneNumbers),
		&userr.Email,
		&userr.Username,
	)
	if err != nil {
		return nil, err
	}
	//defer r.db.Close()
	adress := pb.Adress{}
	for _, adres := range user.Adress {
		adid := uuid.New()
		query1 := "insert into adress (id,users_id,country,city,district,postalCodes) values ($1,$2,$3,$4,$5,$6) RETURNING id,users_id,country,city,district,postalCodes"
		err := r.db.QueryRow(query1, adid, id, adres.Country, adres.City, adres.District, adres.PostalCodes).Scan(
			&adress.Id,
			&adress.UserId,
			&adress.Country,
			&adress.City,
			&adress.District,
			&adress.PostalCodes,
		)
		if err != nil {
			return nil, err
		}
		userr.Adress = append(userr.Adress, &adress)
	}
	return &userr, nil
}

func (r *userRepo) DeleteUser(delete *pb.DeleteById) (*pb.DeleteUserRes, error) {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.Exec(query, delete.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserRes{Status: true}, nil
}
func (r *userRepo) UpdateUser(user *pb.User) (*pb.UpdateUserRes, error) {
	query := "UPDATE users set username=$1,firstname=$2,lastname=$3,bio=$4,updateat=$5,status=$6,phonenumbers=$7,email=$8 where id=$9"
	query1 := "UPDATE adress set country=$1,city=$2,district=$3,postalcodes=$4 where users_id=$5"
	time := time.Now()
	_, err := r.db.Exec(query, user.Username, user.FirstName, user.LastName, user.Bio, time, user.Status, pq.Array(user.PhoneNumbers), user.Email, user.Id)
	if err != nil {
		return nil, err
	}
	for _, adres := range user.Adress {
		_, err := r.db.Exec(query1, adres.Country, adres.City, adres.District, adres.PostalCodes, adres.UserId)
		if err != nil {
			return nil, err
		}
	}
	return &pb.UpdateUserRes{Status: true}, nil

}
func (r *userRepo) GetAllUser(get *pb.GetAllById) (*pb.User, error) {
	var res pb.User
	//var resp []*pb.User
	query1 := "select id,country,city,district,postalcodes from adress where users_id=$1"
	query := "select id,username,firstname,lastname,bio,createdat,status,email from users where id=$1"
	rows, err := r.db.Query(query, get.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&res.Id,
			&res.Username,
			&res.FirstName,
			&res.LastName,
			&res.Bio,
			&res.CreatedAt,
			&res.Status,
			&res.Email,
		)
		if err != nil {
			return nil, err
		}

	}
	row, err := r.db.Query(query1, get.Id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var adres pb.Adress
		err = row.Scan(
			&adres.Id,
			&adres.Country,
			&adres.City,
			&adres.District,
			&adres.PostalCodes,
		)
		if err != nil {
			return nil, err
		}
		res.Adress = append(res.Adress, &adres)
	}
	return &res, nil
}

func (r *userRepo) ListUsers(limit, page int64) ([]*pb.User, int64, error) {
	var (
		users []*pb.User
		count int64
	)
	offset := (page - 1) * limit
	query := "select id,firstname,lastname,bio,status,username,username,phonenumbers from users order by firstname OFFSET $1 LIMIT $2"
	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var user pb.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Bio,
			&user.Status,
			&user.Username,
			&user.Email,
			pq.Array(&user.PhoneNumbers),
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, &user)
	}
	countquery := "SELECT count(*) FROM users"
	err = r.db.QueryRow(countquery).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return users, count, nil
}
func (r *userRepo) CheckUniquess(field, value string) (bool, error) {
	var exists int64
	query := "SELECT count(1) FROM users WHERE $1=$2"
	err := r.db.QueryRow(query, field, value).Scan(
		&exists,
	)
	if err != nil {
		return false, err
	}
	if exists > 0 {
		return true, nil
	}
	return false, nil
}

func (r *userRepo) LoginUser(get *pb.LoginRequest) (*pb.User, error) {
	var res pb.User
	query1 := "select id,country,city,district,postalcodes from adress where users_id=$1"
	query := "select id,username,firstname,lastname,bio,createdat,status,email,password from users where email=$1"
	err := r.db.QueryRow(query, get.Email).Scan(&res.Id,
		&res.Username,
		&res.FirstName,
		&res.LastName,
		&res.Bio,
		&res.CreatedAt,
		&res.Status,
		&res.Email,
		&res.Password)
	
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(get.Password))
	if err != nil {
		return nil, err
	}

	fmt.Println(res.Id)
	row, err := r.db.Query(query1, res.Id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var adres pb.Adress
		err = row.Scan(
			&adres.Id,
			&adres.Country,
			&adres.City,
			&adres.District,
			&adres.PostalCodes,
		)
		if err != nil {
			return nil, err
		}
		res.Adress = append(res.Adress, &adres)
	}
	return &res, nil
}
