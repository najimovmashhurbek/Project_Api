package model

type User struct {
	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	FirstName    string    `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName"`
	LastName     string    `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName"`
	Bio          string    `protobuf:"bytes,6,opt,name=bio,proto3" json:"bio"`
	PhoneNumbers []string  `protobuf:"bytes,7,rep,name=phoneNumbers,proto3" json:"phoneNumbers"`
	Status       string    `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	CreatedAt    string    `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt"`
	UpdateAt     string    `protobuf:"bytes,10,opt,name=updateAt,proto3" json:"updateAt"`
	DeletedAt    string    `protobuf:"bytes,11,opt,name=deletedAt,proto3" json:"deletedAt"`
	Adress       []*Adress `protobuf:"bytes,12,rep,name=adress,proto3" json:"adress"`
	Post         []*Post   `protobuf:"bytes,13,rep,name=post,proto3" json:"post"`
}
type Adress struct {
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Country     string `protobuf:"bytes,3,opt,name=country,proto3" json:"country"`
	City        string `protobuf:"bytes,4,opt,name=city,proto3" json:"city"`
	District    string `protobuf:"bytes,5,opt,name=district,proto3" json:"district"`
	PostalCodes int64  `protobuf:"varint,6,opt,name=postalCodes,proto3" json:"postalCodes"`
}
type Post struct {
	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	UserId      string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Medias      []*Media `protobuf:"bytes,5,rep,name=medias,proto3" json:"medias"`
}
type Media struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`
	Link string `protobuf:"bytes,3,opt,name=link,proto3" json:"link"`
}

type JwtRequestModel struct {
	Token string `string:"token"`
}

type ResponseError struct {
	Error interface{} `json:"error"`
}

// ServerError ...
type ServerError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
