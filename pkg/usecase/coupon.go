package usecase

import (
	"strconv"

	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/domain"
	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository/interface"
	interfaces "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase/interface"
	
)




type CouponUseCase struct{
	Repo repository.CouponRepository
}

func NewCouponUseCase(repo repository.CouponRepository)interfaces.CouponUseCase{
	return &CouponUseCase{repo}
}

func (c *CouponUseCase) AddCoupon(couponId string,discount int,usage uint)error{
	exist,err:=c.Repo.IsExist(couponId)
	if err!=nil{
		return err
	}
	if exist{
		expired,err:=c.Repo.IsExpired(couponId)
		if err!=nil{
			return err
		}
		if expired{
			return c.Repo.UpdateExpiry(couponId)
			
		}
	}
	return c.Repo.AddCoupon(couponId,discount,usage)
}

func (c *CouponUseCase) ExpireCoupon(id string)error{
	co_id,err:=strconv.Atoi(id)
	if err!=nil{
		return err
	}
	return c.Repo.ExpireCoupon(uint(co_id))
}

func (c *CouponUseCase) BlockCoupon(id string) error{
	co_id,err:=strconv.Atoi(id)
	if err!=nil{
		return err
	}
	return c.Repo.BlockCoupon(uint(co_id))
}

func (c *CouponUseCase) UnBlockCoupon(id string) error{
	co_id,err:=strconv.Atoi(id)
	if err!=nil{
		return err
	}
	return c.Repo.UnBlockCoupon(uint(co_id))
}

func (c *CouponUseCase) GetCoupon(pages,counts string)([]domain.Coupon,error){
	page,err:=strconv.Atoi(pages)
	if err!=nil{
		return nil,err
	}
	count,err:=strconv.Atoi(counts)
	if err!=nil{
		return nil,err
	}
	return c.Repo.GetCoupon(page,count)
}