package repository

import (
	"errors"
	"reflect"
	"testing"

	"jerseyhub/pkg/domain"
	"jerseyhub/pkg/utils/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_user_UserSignUp(t *testing.T) {

	type args struct {
		input models.UserDetails
	}

	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    models.UserDetailsResponse
		wantErr error
	}{
		{
			name: "success signup user",
			args: args{
				input: models.UserDetails{Name: "Arun", Email: "arthurbishop120@gmail.com", Phone: "6282246077", Password: "12345"},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("Arun", "arthurbishop120@gmail.com", "12345", "6282246077", "12345").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "referral_code"}).AddRow(1, "Arun", "arthurbishop120@gmail.com", "6282246077", "12345"))

			},

			want:    models.UserDetailsResponse{Id: 1, Name: "Arun", Email: "arthurbishop120@gmail.com", Phone: "6282246077"},
			wantErr: nil,
		},

		{
			name: "error signup user",
			args: args{
				input: models.UserDetails{Name: "", Email: "", Phone: "", Password: ""},
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("", "", "", "", "").
					WillReturnError(errors.New("text string"))

			},

			want:    models.UserDetailsResponse{},
			wantErr: errors.New("Query 'INSERT INTO users (name, email, password, phone,referral_code) VALUES ($1, $2, $3, $4,$5) RETURNING id, name, email, phone', arguments do not match: argument 4 expected [string - ] does not match actual [string - 12345]"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.UserSignUp(tt.args.input, "12345")

			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_CheckUserAvailability(t *testing.T) {

	tests := []struct {
		name string
		args string
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "user available",
			args: "arthurbishop120@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

			},

			want: true,
		},
		{
			name: "user not available",
			args: "arthurbishop120@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

			},

			want: false,
		},
		{
			name: "error from database",
			args: "arthurbishop120@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs("").
					WillReturnError(errors.New("text string"))

			},

			want: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result := u.CheckUserAvailability(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_UserBlockStatus(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    bool
		wantErr error
	}{
		{
			name: "user is blocked",
			args: "arthurbishop120@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select blocked from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"is_blocked"}).AddRow(true))

			},

			want:    true,
			wantErr: nil,
		},
		{
			name: "user is not blocked",
			args: "arthurbishop120@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select blocked from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"is_blocked"}).AddRow(false))

			},

			want:    false,
			wantErr: nil,
		},
		{
			name: "error from database",
			args: "arthurbishop120@gmail.com",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select blocked from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("text string"))

			},

			want:    false,
			wantErr: errors.New("text string"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.UserBlockStatus(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindUserByEmail(t *testing.T) {

	tests := []struct {
		name    string
		args    models.UserLogin
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "success",
			args: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "1234",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^SELECT \* FROM users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs("arthurbishop120@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "email", "phone", "password"}).AddRow(1, 1, "Arun K", "arthurbishop120@gmail.com", "6282246077", "4321"))

			},

			want: models.UserSignInResponse{
				Id:       1,
				UserID:   1,
				Name:     "Arun K",
				Email:    "arthurbishop120@gmail.com",
				Phone:    "6282246077",
				Password: "4321",
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: models.UserLogin{
				Email:    "arthurbishop120@gmail.com",
				Password: "1234",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^SELECT \* FROM users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).WithArgs("arthurbishop120@gmail.com").
					WillReturnError(errors.New("new error"))

			},

			want:    models.UserSignInResponse{},
			wantErr: errors.New("error checking user details"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindUserByEmail(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_AddAddress(t *testing.T) {

	tests := []struct {
		name    string
		args    models.AddAddress
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "success",
			args: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("INSERT INTO addresses").WithArgs(1, "Arun K", "nellikkal", "pallippuram", "cherthala", "kerala", "688541", true).WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: models.AddAddress{
				Name:      "Arun K",
				HouseName: "nellikkal",
				Street:    "pallippuram",
				City:      "cherthala",
				State:     "kerala",
				Pin:       "688541",
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("INSERT INTO addresses").WithArgs(1, "Arun K", "nellikkal", "pallippuram", "cherthala", "kerala", "688541", true).WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("could not add address"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.AddAddress(1, tt.args, true)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_CheckIfFirstAddress(t *testing.T) {

	tests := []struct {
		name string
		args int
		stub func(sqlmock.Sqlmock)
		want bool
	}{
		{
			name: "first address",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{gomock.Any().String()}).AddRow(2))

			},

			want: true,
		},
		{
			name: "error occured",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select count\(\*\) from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result := u.CheckIfFirstAddress(tt.args)

			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_GetAddresses(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    []domain.Address
		wantErr error
	}{
		{
			name: "successfully got all addresses",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select \* from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "house_name", "street", "city", "state", "pin", "default"}).AddRow(1, 1, "a", "b", "c", "d", "e", "f", true).AddRow(2, 1, "a", "b", "c", "d", "e", "f", false))

			},

			want: []domain.Address{
				{Id: 1,
					UserID:    1,
					Name:      "a",
					HouseName: "b",
					Street:    "c",
					City:      "d",
					State:     "e",
					Pin:       "f",
					Default:   true,
				}, {
					Id:        2,
					UserID:    1,
					Name:      "a",
					HouseName: "b",
					Street:    "c",
					City:      "d",
					State:     "e",
					Pin:       "f",
					Default:   false,
				},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select \* from addresses(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want:    []domain.Address{},
			wantErr: errors.New("error in getting addresses"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetAddresses(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_GetUserDetails(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    models.UserDetailsResponse
		wantErr error
	}{
		{
			name: "successfully got details",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				// expectedQuery := `^select \* from users(.+)$`,

				mockSQL.ExpectQuery(`^select id\,name\,email\,phone from users(.+)$`).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(1, "arun k", "arthurbishop120@gmail.com", "6282246077"))
			},

			want: models.UserDetailsResponse{
				Id:    1,
				Name:  "arun k",
				Email: "arthurbishop120@gmail.com",
				Phone: "6282246077",
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectQuery(`^select id\,name\,email\,phone from users(.+)$`).
					WillReturnError(errors.New("error"))
			},

			want:    models.UserDetailsResponse{},
			wantErr: errors.New("could not get user details"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetUserDetails(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_ChangePassword(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id       int
			password string
		}
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id       int
				password string
			}{id: 1, password: gomock.Any().String()},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("UPDATE users SET").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id       int
				password string
			}{id: 1, password: gomock.Any().String()},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("UPDATE users SET").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.ChangePassword(tt.args.id, tt.args.password)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_GetPassword(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select password from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(gomock.Any().String()))

			},

			want:    gomock.Any().String(),
			wantErr: nil,
		},
		{
			name: "failure",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select password from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want:    "",
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetPassword(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindIDFromPhone(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: "6282246077",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select id from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			},

			want:    1,
			wantErr: nil,
		},
		{
			name: "failure",
			args: "6282246077",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select id from users(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindIdFromPhone(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_EditName(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id   int
			name string
		}
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id   int
				name string
			}{id: 1, name: "Arun K"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id   int
				name string
			}{id: 1, name: "Arun K"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditName(tt.args.id, tt.args.name)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_EditEmail(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id    int
			email string
		}
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id    int
				email string
			}{id: 1, email: "arthurbishop120@gmail.com"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id    int
				email string
			}{id: 1, email: "arthurbishop120@gmail.com"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditEmail(tt.args.id, tt.args.email)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_EditPhone(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			id    int
			phone string
		}
		stub    func(sqlmock.Sqlmock)
		want    models.UserSignInResponse
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id    int
				phone string
			}{id: 1, phone: "6282246077"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				id    int
				phone string
			}{id: 1, phone: "6282246077"},
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("update users set").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.EditPhone(tt.args.id, tt.args.phone)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_GetCart(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    []models.GetCart
		wantErr error
	}{
		{
			name: "successfull",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select inventories\.product\_name\,cart\_products\.quantity\,cart\_products\.total\_price from cart\_products(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"product_name", "quantity", "total"}).AddRow("a", 1, 400.0))

			},

			want: []models.GetCart{{
				ProductName: "a",
				Quantity:    1,
				Total:       400.0,
			},
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^select inventories\.product\_name\,cart\_products\.quantity\,cart\_products\.total\_price from cart\_products(.+)$`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},

			want:    []models.GetCart{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetCart(tt.args)

			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, result)
		})
	}

}

func Test_RemoveFromCart(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `delete from cart_products`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `delete from cart_products`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			err := u.RemoveFromCart(tt.args)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_UpdateQuantityAdd(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			id     int
			inv_id int
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `UPDATE line_items
				SET quantity`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},

		{
			name: "error",
			args: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `UPDATE line_items
				SET quantity`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			err := u.UpdateQuantityAdd(tt.args.id, tt.args.inv_id)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_UpdateQuantityLess(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			id     int
			inv_id int
		}
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `UPDATE line_items
				SET quantity`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},

		{
			name: "error",
			args: struct {
				id     int
				inv_id int
			}{
				id:     1,
				inv_id: 1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `UPDATE line_items
				SET quantity`

				mockSQL.ExpectExec(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			err := u.UpdateQuantityLess(tt.args.id, tt.args.inv_id)

			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_GetCartID(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select id from carts`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select id from carts`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetCartID(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_GetProductsInCart(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    []int
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select inventory_id from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2))

			},
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select inventory_id from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    []int{},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetProductsInCart(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindProductsNames(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select product_name from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"product_name"}).AddRow("fc barcelona homekit"))

			},
			want:    "fc barcelona homekit",
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select product_name from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    "",
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindProductNames(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_CartQuantity(t *testing.T) {

	tests := []struct {
		name string
		args struct {
			cart_id int
			inv_id  int
		}
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: struct {
				cart_id int
				inv_id  int
			}{
				cart_id: 1,
				inv_id:  1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select quantity from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"quantity"}).AddRow(1))

			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "error",
			args: struct {
				cart_id int
				inv_id  int
			}{
				cart_id: 1,
				inv_id:  1,
			},
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select quantity from line_items`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindCartQuantity(tt.args.cart_id, tt.args.inv_id)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindPrice(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    float64
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select price from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"price"}).AddRow(400))

			},
			want:    400,
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select price from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindPrice(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindCategory(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select category_id from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"category_id"}).AddRow(1))

			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select category_id from inventories`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindCategory(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindOfferPercentage(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select discount_rate from offers`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"discount_rate"}).AddRow(10))

			},
			want:    10,
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `select discount_rate from offers`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindofferPercentage(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_FindUserFromReference(t *testing.T) {

	tests := []struct {
		name    string
		args    string
		stub    func(sqlmock.Sqlmock)
		want    int
		wantErr error
	}{
		{
			name: "success",
			args: "SRK123",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `SELECT id FROM users`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "error",
			args: "SRK123",
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `SELECT id FROM users`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    0,
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.FindUserFromReference(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}

func Test_CreditReferencePointsToWallet(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("Update wallets set amount").WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

			},
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectExec("Update wallets set amount").WithArgs().WillReturnError(errors.New("error"))

			},
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)
			u := NewUserRepository(gormDB)

			err := u.CreditReferencePointsToWallet(tt.args)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_GetReferralCodeFromID(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		stub    func(sqlmock.Sqlmock)
		want    string
		wantErr error
	}{
		{
			name: "success",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `SELECT referral_code FROM users`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows([]string{"referral_code"}).AddRow("SRK123"))

			},
			want:    "SRK123",
			wantErr: nil,
		},
		{
			name: "error",
			args: 1,
			stub: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `SELECT referral_code FROM users`

				mockSQL.ExpectQuery(expectedQuery).
					WillReturnError(errors.New("error"))

			},
			want:    "",
			wantErr: errors.New("error"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			result, err := u.GetReferralCodeFromID(tt.args)

			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}

}
