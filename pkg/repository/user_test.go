package repository

import (
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_userDatabase_FindByEmail(t *testing.T) {

	type args struct {
		email string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.UserLoginCheck
		wantErr    error
	}{
		{
			name: "success",
			args: args{email: "athul@gmail.com"},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`select id,name,email,ph_no as phno,password from users where email=$1 and role='user'`)).
					WithArgs("athul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "ph_no", "password"}).
						AddRow(1, "athul", "athul@gmail.com", "+919633513774", "$2a$10$MKaWljgSCiDRydVm9Gxhe.2d.xe9iOfmyH1gRvYD3tgTKM.DKw1Ni"))

			},
			want:    models.UserLoginCheck{ID: 1, Name: "athul", Email: "athul@gmail.com", PhNo: "+919633513774", Password: "$2a$10$MKaWljgSCiDRydVm9Gxhe.2d.xe9iOfmyH1gRvYD3tgTKM.DKw1Ni"},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{email: "athul@gmail.com"},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`select id,name,email,ph_no as phno,password from users where email=$1 and role='user'`)).
					WithArgs("athul@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{}).
						AddRow()).WillReturnError(errors.New("error in checking userdetails"))
			},
			want:    models.UserLoginCheck{},
			wantErr: errors.New("error in checking userdetails"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

			tt.beforeTest(mockSql)

			u := NewUserRepository(gormDB)

			got, err := u.FindByEmail(tt.args.email)
			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDatabase.FindByEmail()=%v want %v", got, tt.want)
			}
		})
	}
}

func Test_userDatabase_Save(t *testing.T) {
	type args struct {
		user models.UserSignUp
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       models.UserDetails
		wantErr    error
	}{
		{
			name: "successfull save",
			args: args{user: models.UserSignUp{Name: "aarathy", Email: "aarathy@gmail.com", PhNo: "+919745503907", Password: "1234"}},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`insert into users(name,email,ph_no,password,role) values($1,$2,$3,$4,$5) returning id,name,email,ph_no`)).WithArgs("aarathy", "aarathy@gmail.com", "+919745503907", "1234", "user").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "ph_no"}).AddRow(1, "aarathy", "aarathy@gmail.com", "+919745503907"))
			},
			want:    models.UserDetails{ID: 1, Name: "aarathy", Email: "aarathy@gmail.com", PhNo: "+919745503907"},
			wantErr: nil,
		},
		{
			name: "error in save",
			args: args{user: models.UserSignUp{Name: "aarathy", Email: "aarathy@gmail.com", PhNo: "+919745503907", Password: "1234"}},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`insert into users(name,email,ph_no,password,role) values($1,$2,$3,$4,$5) returning id,name,email,ph_no`)).WithArgs("aarathy", "aarathy@gmail.com", "+919745503907", "1234", "user").WillReturnRows(sqlmock.NewRows([]string{}).AddRow()).WillReturnError(errors.New("error saving in database"))
			},
			want:    models.UserDetails{},
			wantErr: errors.New("error saving in database"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

			tt.beforeTest(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.Save(tt.args.user)
			assert.Equal(t, tt.wantErr, err)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userDatabase_Save()=%v want %v", got, tt.want)
			}
		})
	}
}

func Test_userDatabase_UpdateUserDetails(t *testing.T) {
	type args struct {
		userId      uint
		userDetails models.UserUpdate
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlmock.Sqlmock)
		want       error
	}{
		{
			name: "successful updation",
			args: args{userId: 1, userDetails: models.UserUpdate{Name: "aarathy", Email: "aarathy@gmail.com", PhNo: "+919745503907"}},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(regexp.QuoteMeta(`UPDATE users SET name=$1,email=$2,ph_no=$3 WHERE id=$4`)).WithArgs("aarathy", "aarathy@gmail.com", "+919745503907", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			want: nil,
		},
		{
			name: "update error",
			args: args{userId: 1, userDetails: models.UserUpdate{Name: "aarathy", Email: "aarathy@gmail.com", PhNo: "+919745503907"}},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec(regexp.QuoteMeta(`UPDATE users SET name=$1,email=$2,ph_no=$3 WHERE id=$4`)).WithArgs("aarathy", "aarathy@gmail.com", "+919745503907", 1).WillReturnError(errors.New("error while updating"))
			},
			want: errors.New("error while updating"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

			tt.beforeTest(mockSQL)

			u := NewUserRepository(gormDB)

			err := u.UpdateUserDetails(tt.args.userId, tt.args.userDetails)

			assert.Equal(t, err, tt.want)
		})
	}

}

func TestIsBlocked(t *testing.T){
	type args struct{
		email string
	}
	tests:=[]struct{
		name string
		args args
		beforeTest func(sqlmock.Sqlmock)
		want bool
		wantErr error
	}{
		{
			name: "user is blocked",
			args: args{email: "aarathy@gmail.com"},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`select block from users where email=$1`)).WithArgs("aarathy@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"block"}).AddRow(true))
			},
			want: true,
			wantErr: nil,
		},
		{
			name: "user is not blocked",
			args: args{email: "aarathy@gmail.com"},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`select block from users where email=$1`)).WithArgs("aarathy@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"block"}).AddRow(false))
			},
			want: false,
			wantErr: nil,
		},
		{
			name: "error in fetching",
			args: args{email: "aarathy@gmail.com"},
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`select block from users where email=$1`)).WithArgs("aarathy@gmail.com").WillReturnRows(sqlmock.NewRows([]string{}).AddRow()).WillReturnError(errors.New("error in fetching block detail"))
			},
			want: false,
			wantErr: errors.New("error in fetching block detail"),
		},
	}

	for _,tt:=range tests{
		t.Run(tt.name,func(t *testing.T) {
			mockDB,mockSQL,_:=sqlmock.New()
			defer mockDB.Close()

			gormDB,_:=gorm.Open(postgres.New(postgres.Config{Conn: mockDB}),&gorm.Config{})

			tt.beforeTest(mockSQL)

			u:=NewUserRepository(gormDB)

			got,err:=u.IsBlocked(tt.args.email)

			assert.Equal(t,err,tt.wantErr)

			if !reflect.DeepEqual(got,tt.want){
				t.Errorf("userDatabse_IsBlocked()=%v want %v ",got,tt.want)
			}
		})
	}
}
