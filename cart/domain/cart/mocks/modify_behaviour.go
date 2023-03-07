// Code generated by mockery v2.21.1. DO NOT EDIT.

package mocks

import (
	context "context"

	cart "flamingo.me/flamingo-commerce/v3/cart/domain/cart"

	mock "github.com/stretchr/testify/mock"
)

// ModifyBehaviour is an autogenerated mock type for the ModifyBehaviour type
type ModifyBehaviour struct {
	mock.Mock
}

// AddToCart provides a mock function with given fields: ctx, _a1, deliveryCode, addRequest
func (_m *ModifyBehaviour) AddToCart(ctx context.Context, _a1 *cart.Cart, deliveryCode string, addRequest cart.AddRequest) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode, addRequest)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, cart.AddRequest) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, deliveryCode, addRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, cart.AddRequest) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode, addRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, cart.AddRequest) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode, addRequest)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, cart.AddRequest) error); ok {
		r2 = rf(ctx, _a1, deliveryCode, addRequest)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ApplyVoucher provides a mock function with given fields: ctx, _a1, couponCode
func (_m *ModifyBehaviour) ApplyVoucher(ctx context.Context, _a1 *cart.Cart, couponCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, couponCode)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, couponCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, couponCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CleanCart provides a mock function with given fields: ctx, _a1
func (_m *ModifyBehaviour) CleanCart(ctx context.Context, _a1 *cart.Cart) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart) *cart.Cart); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CleanDelivery provides a mock function with given fields: ctx, _a1, deliveryCode
func (_m *ModifyBehaviour) CleanDelivery(ctx context.Context, _a1 *cart.Cart, deliveryCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, deliveryCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, deliveryCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteItem provides a mock function with given fields: ctx, _a1, itemID, deliveryCode
func (_m *ModifyBehaviour) DeleteItem(ctx context.Context, _a1 *cart.Cart, itemID string, deliveryCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, itemID, deliveryCode)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, string) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, itemID, deliveryCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, itemID, deliveryCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, itemID, deliveryCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, string) error); ok {
		r2 = rf(ctx, _a1, itemID, deliveryCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RemoveVoucher provides a mock function with given fields: ctx, _a1, couponCode
func (_m *ModifyBehaviour) RemoveVoucher(ctx context.Context, _a1 *cart.Cart, couponCode string) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, couponCode)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, couponCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string) *cart.Cart); ok {
		r0 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, couponCode)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string) error); ok {
		r2 = rf(ctx, _a1, couponCode)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateAdditionalData provides a mock function with given fields: ctx, _a1, additionalData
func (_m *ModifyBehaviour) UpdateAdditionalData(ctx context.Context, _a1 *cart.Cart, additionalData *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, additionalData)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, additionalData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *cart.AdditionalData) *cart.Cart); ok {
		r0 = rf(ctx, _a1, additionalData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, *cart.AdditionalData) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, additionalData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, *cart.AdditionalData) error); ok {
		r2 = rf(ctx, _a1, additionalData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateBillingAddress provides a mock function with given fields: ctx, _a1, billingAddress
func (_m *ModifyBehaviour) UpdateBillingAddress(ctx context.Context, _a1 *cart.Cart, billingAddress cart.Address) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, billingAddress)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.Address) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, billingAddress)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.Address) *cart.Cart); ok {
		r0 = rf(ctx, _a1, billingAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, cart.Address) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, billingAddress)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, cart.Address) error); ok {
		r2 = rf(ctx, _a1, billingAddress)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateDeliveryInfo provides a mock function with given fields: ctx, _a1, deliveryCode, deliveryInfo
func (_m *ModifyBehaviour) UpdateDeliveryInfo(ctx context.Context, _a1 *cart.Cart, deliveryCode string, deliveryInfo cart.DeliveryInfoUpdateCommand) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode, deliveryInfo)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, deliveryCode, deliveryInfo)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode, deliveryInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode, deliveryInfo)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, cart.DeliveryInfoUpdateCommand) error); ok {
		r2 = rf(ctx, _a1, deliveryCode, deliveryInfo)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateDeliveryInfoAdditionalData provides a mock function with given fields: ctx, _a1, deliveryCode, additionalData
func (_m *ModifyBehaviour) UpdateDeliveryInfoAdditionalData(ctx context.Context, _a1 *cart.Cart, deliveryCode string, additionalData *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, deliveryCode, additionalData)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, deliveryCode, additionalData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) *cart.Cart); ok {
		r0 = rf(ctx, _a1, deliveryCode, additionalData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, deliveryCode, additionalData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, string, *cart.AdditionalData) error); ok {
		r2 = rf(ctx, _a1, deliveryCode, additionalData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateItem provides a mock function with given fields: ctx, _a1, itemUpdateCommand
func (_m *ModifyBehaviour) UpdateItem(ctx context.Context, _a1 *cart.Cart, itemUpdateCommand cart.ItemUpdateCommand) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, itemUpdateCommand)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, itemUpdateCommand)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) *cart.Cart); ok {
		r0 = rf(ctx, _a1, itemUpdateCommand)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, itemUpdateCommand)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, cart.ItemUpdateCommand) error); ok {
		r2 = rf(ctx, _a1, itemUpdateCommand)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateItems provides a mock function with given fields: ctx, _a1, itemUpdateCommands
func (_m *ModifyBehaviour) UpdateItems(ctx context.Context, _a1 *cart.Cart, itemUpdateCommands []cart.ItemUpdateCommand) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, itemUpdateCommands)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, itemUpdateCommands)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) *cart.Cart); ok {
		r0 = rf(ctx, _a1, itemUpdateCommands)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, itemUpdateCommands)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, []cart.ItemUpdateCommand) error); ok {
		r2 = rf(ctx, _a1, itemUpdateCommands)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePaymentSelection provides a mock function with given fields: ctx, _a1, paymentSelection
func (_m *ModifyBehaviour) UpdatePaymentSelection(ctx context.Context, _a1 *cart.Cart, paymentSelection cart.PaymentSelection) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, paymentSelection)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.PaymentSelection) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, paymentSelection)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, cart.PaymentSelection) *cart.Cart); ok {
		r0 = rf(ctx, _a1, paymentSelection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, cart.PaymentSelection) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, paymentSelection)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, cart.PaymentSelection) error); ok {
		r2 = rf(ctx, _a1, paymentSelection)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePurchaser provides a mock function with given fields: ctx, _a1, purchaser, additionalData
func (_m *ModifyBehaviour) UpdatePurchaser(ctx context.Context, _a1 *cart.Cart, purchaser *cart.Person, additionalData *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error) {
	ret := _m.Called(ctx, _a1, purchaser, additionalData)

	var r0 *cart.Cart
	var r1 cart.DeferEvents
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) (*cart.Cart, cart.DeferEvents, error)); ok {
		return rf(ctx, _a1, purchaser, additionalData)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) *cart.Cart); ok {
		r0 = rf(ctx, _a1, purchaser, additionalData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cart.Cart)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) cart.DeferEvents); ok {
		r1 = rf(ctx, _a1, purchaser, additionalData)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(cart.DeferEvents)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *cart.Cart, *cart.Person, *cart.AdditionalData) error); ok {
		r2 = rf(ctx, _a1, purchaser, additionalData)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewModifyBehaviour interface {
	mock.TestingT
	Cleanup(func())
}

// NewModifyBehaviour creates a new instance of ModifyBehaviour. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewModifyBehaviour(t mockConstructorTestingTNewModifyBehaviour) *ModifyBehaviour {
	mock := &ModifyBehaviour{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
