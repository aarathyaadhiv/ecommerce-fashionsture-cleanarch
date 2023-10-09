package usecase

// import (
// 	"errors"
// 	"testing"

// 	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/mock"
// 	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/models"
// 	"github.com/golang/mock/gomock"
// )




// func Test_CouponUseCase_AddCoupon(t *testing.T) {
// 	ctrl:=gomock.NewController(t)
// 	defer ctrl.Finish()

// 	couponRepo:=mock.NewMockCouponRepository(ctrl)

// 	couponUseCase:=NewCouponUseCase(couponRepo)

// 	type args struct{
// 		coupon models.AddCoupon
// 	}
// 	tests:=[]struct{
// 		name string
// 		args args
// 		beforeTest func( mock.MockCouponRepository,models.AddCoupon)
// 		expected error
// 	}{
// 		{
// 			name: "adding coupon",
// 			args: args{coupon: models.AddCoupon{CouponId: "get18",Discount: 18,Usage: 3,Expiry: "2023-10-21",MinimumPurchase: 500,MaximumAmount: 100}},
// 			beforeTest: func(repo mock.MockCouponRepository,coupon models.AddCoupon) {
// 				gomock.InOrder(
// 					repo.EXPECT().IsExist(coupon.CouponId).Times(1).Return(false,nil),
// 					repo.EXPECT().AddCoupon(coupon).Times(1).Return(nil),
// 				)
// 			},
// 			expected: nil,
// 		},
// 		{
// 			name: "already existing coupon",
// 			args: args{coupon: models.AddCoupon{CouponId: "get18",Discount: 18,Usage: 3,Expiry: "2023-10-21",MinimumPurchase: 500,MaximumAmount: 100}},
// 			beforeTest: func(repo mock.MockCouponRepository,coupon models.AddCoupon) {
// 				gomock.InOrder(
// 					repo.EXPECT().IsExist(coupon.CouponId).Times(1).Return(true,nil),
// 					repo.EXPECT().AddCoupon(coupon).Times(1).Return(errors.New("already existing coupon")),
// 				)
// 			},
// 			expected: errors.New("already existing coupon"),
// 		},
// 	}
// 	for _,tt:=range tests{
// 		t.Run(tt.name,func(t *testing.T) {
// 			if tt.beforeTest!=nil{
// 				tt.beforeTest(*couponRepo,tt.args.coupon)
// 			}
// 			got:=couponUseCase.AddCoupon(tt.args.coupon)

// 			if got!=tt.expected{
// 				t.Errorf("CouponUseCase_AddCoupon()=%v want %v",got,tt.expected)
// 			}

// 		})
// 	}
// }