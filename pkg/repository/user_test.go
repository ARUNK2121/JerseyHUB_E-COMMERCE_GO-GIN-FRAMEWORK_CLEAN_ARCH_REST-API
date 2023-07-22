package repository

import (
	"reflect"
	"testing"

	"jerseyhub/pkg/utils/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_user_UserSignUp(t *testing.T) {

	type args struct {
		input models.UserDetails
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.UserDetailsResponse
		wantErr    error
	}{
		{
			name: "success signup user",
			args: args{
				input: models.UserDetails{Name: "Arun", Email: "arthurbishop120@gmail.com", Phone: "6282246077", Password: "12345"},
			},
			beforeTest: func(mockSQL sqlmock.Sqlmock) {

				expectedQuery := `^INSERT INTO users (.+)$`
				mockSQL.ExpectQuery(expectedQuery).WithArgs("Arun", "arthurbishop120@gmail.com", "12345", "6282246077", "12345").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "referral_code"}).AddRow(1, "Arun", "arthurbishop120@gmail.com", "6282246077", "12345"))

			},

			want:    models.UserDetailsResponse{Id: 1, Name: "Arun", Email: "arthurbishop120@gmail.com", Phone: "6282246077"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.beforeTest(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.UserSignUp(tt.args.input, "12345")

			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.UserSignUp() = %v, want %v", got, tt.want)
			}
		})
	}

}
