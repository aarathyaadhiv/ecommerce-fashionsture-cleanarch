package usecase

import (
	"errors"
	"strconv"
	"time"

	_ "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/helper"
	repo "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	services "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	adminRepo repo.AdminRepository
}

func NewAdminUseCase(repo repo.AdminRepository) services.AdminUseCase {
	return &AdminUseCase{
		adminRepo: repo,
	}
}

func (c *AdminUseCase) SignUp(admin models.AdminSignUp) (models.AdminTokenResponse, error) {
	if ok := c.adminRepo.CheckAdminAvailability(admin.Email); ok {
		return models.AdminTokenResponse{}, errors.New("already existing email")
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return models.AdminTokenResponse{}, errors.New("error in password hashing")
	}
	admin.Password = string(hashPassword)

	adminDetails, err := c.adminRepo.Save(admin)
	if err != nil {
		return models.AdminTokenResponse{}, errors.New("error in saving user data")
	}

	tokenString, err := helper.GenerateAdminToken(adminDetails)
	if err != nil {
		return models.AdminTokenResponse{}, err
	}
	return models.AdminTokenResponse{AdminDetails: adminDetails, Token: tokenString}, nil
}

func (c *AdminUseCase) Login(admin models.AdminLogin) (models.AdminTokenResponse, error) {
	if ok := c.adminRepo.CheckAdminAvailability(admin.Email); !ok {
		return models.AdminTokenResponse{}, errors.New("no such user exist")
	}

	adminCompare, err := c.adminRepo.FindByEmail(admin.Email)
	if err != nil {
		return models.AdminTokenResponse{}, errors.New("error in fetching userdata")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(adminCompare.Password), []byte(admin.Password)); err != nil {
		return models.AdminTokenResponse{}, errors.New("password is incorrect")
	}

	var adminDetails models.AdminDetails
	if err := copier.Copy(&adminDetails, &adminCompare); err != nil {
		return models.AdminTokenResponse{}, err
	}
	tokenString, err := helper.GenerateAdminToken(adminDetails)
	if err != nil {
		return models.AdminTokenResponse{}, err
	}

	return models.AdminTokenResponse{AdminDetails: adminDetails, Token: tokenString}, nil
}

func (c *AdminUseCase) BlockUser(id string) error {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if ok := c.adminRepo.IsBlocked(uint(userId)); ok {
		return errors.New("already blocked user")
	}
	return c.adminRepo.BlockUser(uint(userId))
}

func (c *AdminUseCase) UnblockUser(id string) error {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if ok := c.adminRepo.IsBlocked(uint(userId)); !ok {
		return errors.New("already unblocked user")
	}
	return c.adminRepo.UnblockUser(uint(userId))
}

func (c *AdminUseCase) ListUsers(pages,counts string) ([]models.AdminUserResponse, error) {
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	return c.adminRepo.ListUsers(page,count)
}

func (c *AdminUseCase) AdminHome(id uint) (models.AdminDetails, error) {

	return c.adminRepo.AdminDetails(id)
}

func (c *AdminUseCase) Dashboard() (models.Dashboard, error) {
	revenue, err := c.adminRepo.DashboardRevenue()
	if err != nil {
		return models.Dashboard{}, err
	}
	orders, err := c.adminRepo.DashboardOrders()
	if err != nil {
		return models.Dashboard{}, err
	}
	amount, err := c.adminRepo.DashboardAmount()
	if err != nil {
		return models.Dashboard{}, err
	}
	users, err := c.adminRepo.DashboardUsers()
	if err != nil {
		return models.Dashboard{}, err
	}
	product, err := c.adminRepo.DashboardProduct()
	if err != nil {
		return models.Dashboard{}, err
	}
	return models.Dashboard{DashboardRevenue: revenue,
		DashboardOrders:  orders,
		DashboardAmount:  amount,
		DashboardUsers:   users,
		DashboardProduct: product}, nil
}


func (c *AdminUseCase) SalesReport(timeWord string)(models.SalesReport,error){
	var startDate,endDate time.Time
	if timeWord=="day"{
		startDate,endDate=time.Now().AddDate(0,0,-1),time.Now()
	}else if timeWord=="week"{
		startDate,endDate=time.Now().AddDate(0,0,-7),time.Now()
	}else if timeWord=="year"{
		startDate,endDate=time.Now().AddDate(-1,0,0),time.Now()
	}
	salesReport,err:=c.adminRepo.SalesReport(startDate,endDate)
	if err!=nil{
		return models.SalesReport{},err
	}
	return salesReport,nil

}