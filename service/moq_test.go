// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"github.com/t-tazy/my_portfolio_api/entity"
	"github.com/t-tazy/my_portfolio_api/store"
	"sync"
)

// Ensure, that ExerciseAdderMock does implement ExerciseAdder.
// If this is not the case, regenerate this file with moq.
var _ ExerciseAdder = &ExerciseAdderMock{}

// ExerciseAdderMock is a mock implementation of ExerciseAdder.
//
//	func TestSomethingThatUsesExerciseAdder(t *testing.T) {
//
//		// make and configure a mocked ExerciseAdder
//		mockedExerciseAdder := &ExerciseAdderMock{
//			AddExerciseFunc: func(ctx context.Context, db store.Execer, e *entity.Exercise) error {
//				panic("mock out the AddExercise method")
//			},
//		}
//
//		// use mockedExerciseAdder in code that requires ExerciseAdder
//		// and then make assertions.
//
//	}
type ExerciseAdderMock struct {
	// AddExerciseFunc mocks the AddExercise method.
	AddExerciseFunc func(ctx context.Context, db store.Execer, e *entity.Exercise) error

	// calls tracks calls to the methods.
	calls struct {
		// AddExercise holds details about calls to the AddExercise method.
		AddExercise []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Execer
			// E is the e argument value.
			E *entity.Exercise
		}
	}
	lockAddExercise sync.RWMutex
}

// AddExercise calls AddExerciseFunc.
func (mock *ExerciseAdderMock) AddExercise(ctx context.Context, db store.Execer, e *entity.Exercise) error {
	if mock.AddExerciseFunc == nil {
		panic("ExerciseAdderMock.AddExerciseFunc: method is nil but ExerciseAdder.AddExercise was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Execer
		E   *entity.Exercise
	}{
		Ctx: ctx,
		Db:  db,
		E:   e,
	}
	mock.lockAddExercise.Lock()
	mock.calls.AddExercise = append(mock.calls.AddExercise, callInfo)
	mock.lockAddExercise.Unlock()
	return mock.AddExerciseFunc(ctx, db, e)
}

// AddExerciseCalls gets all the calls that were made to AddExercise.
// Check the length with:
//
//	len(mockedExerciseAdder.AddExerciseCalls())
func (mock *ExerciseAdderMock) AddExerciseCalls() []struct {
	Ctx context.Context
	Db  store.Execer
	E   *entity.Exercise
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Execer
		E   *entity.Exercise
	}
	mock.lockAddExercise.RLock()
	calls = mock.calls.AddExercise
	mock.lockAddExercise.RUnlock()
	return calls
}

// Ensure, that ExerciseListerMock does implement ExerciseLister.
// If this is not the case, regenerate this file with moq.
var _ ExerciseLister = &ExerciseListerMock{}

// ExerciseListerMock is a mock implementation of ExerciseLister.
//
//	func TestSomethingThatUsesExerciseLister(t *testing.T) {
//
//		// make and configure a mocked ExerciseLister
//		mockedExerciseLister := &ExerciseListerMock{
//			ListExercisesFunc: func(ctx context.Context, db store.Queryer) (entity.Exercises, error) {
//				panic("mock out the ListExercises method")
//			},
//		}
//
//		// use mockedExerciseLister in code that requires ExerciseLister
//		// and then make assertions.
//
//	}
type ExerciseListerMock struct {
	// ListExercisesFunc mocks the ListExercises method.
	ListExercisesFunc func(ctx context.Context, db store.Queryer) (entity.Exercises, error)

	// calls tracks calls to the methods.
	calls struct {
		// ListExercises holds details about calls to the ListExercises method.
		ListExercises []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
		}
	}
	lockListExercises sync.RWMutex
}

// ListExercises calls ListExercisesFunc.
func (mock *ExerciseListerMock) ListExercises(ctx context.Context, db store.Queryer) (entity.Exercises, error) {
	if mock.ListExercisesFunc == nil {
		panic("ExerciseListerMock.ListExercisesFunc: method is nil but ExerciseLister.ListExercises was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Queryer
	}{
		Ctx: ctx,
		Db:  db,
	}
	mock.lockListExercises.Lock()
	mock.calls.ListExercises = append(mock.calls.ListExercises, callInfo)
	mock.lockListExercises.Unlock()
	return mock.ListExercisesFunc(ctx, db)
}

// ListExercisesCalls gets all the calls that were made to ListExercises.
// Check the length with:
//
//	len(mockedExerciseLister.ListExercisesCalls())
func (mock *ExerciseListerMock) ListExercisesCalls() []struct {
	Ctx context.Context
	Db  store.Queryer
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Queryer
	}
	mock.lockListExercises.RLock()
	calls = mock.calls.ListExercises
	mock.lockListExercises.RUnlock()
	return calls
}

// Ensure, that UserRegisterMock does implement UserRegister.
// If this is not the case, regenerate this file with moq.
var _ UserRegister = &UserRegisterMock{}

// UserRegisterMock is a mock implementation of UserRegister.
//
//	func TestSomethingThatUsesUserRegister(t *testing.T) {
//
//		// make and configure a mocked UserRegister
//		mockedUserRegister := &UserRegisterMock{
//			RegisterUserFunc: func(ctx context.Context, db store.Execer, u *entity.User) error {
//				panic("mock out the RegisterUser method")
//			},
//		}
//
//		// use mockedUserRegister in code that requires UserRegister
//		// and then make assertions.
//
//	}
type UserRegisterMock struct {
	// RegisterUserFunc mocks the RegisterUser method.
	RegisterUserFunc func(ctx context.Context, db store.Execer, u *entity.User) error

	// calls tracks calls to the methods.
	calls struct {
		// RegisterUser holds details about calls to the RegisterUser method.
		RegisterUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Execer
			// U is the u argument value.
			U *entity.User
		}
	}
	lockRegisterUser sync.RWMutex
}

// RegisterUser calls RegisterUserFunc.
func (mock *UserRegisterMock) RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error {
	if mock.RegisterUserFunc == nil {
		panic("UserRegisterMock.RegisterUserFunc: method is nil but UserRegister.RegisterUser was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Db  store.Execer
		U   *entity.User
	}{
		Ctx: ctx,
		Db:  db,
		U:   u,
	}
	mock.lockRegisterUser.Lock()
	mock.calls.RegisterUser = append(mock.calls.RegisterUser, callInfo)
	mock.lockRegisterUser.Unlock()
	return mock.RegisterUserFunc(ctx, db, u)
}

// RegisterUserCalls gets all the calls that were made to RegisterUser.
// Check the length with:
//
//	len(mockedUserRegister.RegisterUserCalls())
func (mock *UserRegisterMock) RegisterUserCalls() []struct {
	Ctx context.Context
	Db  store.Execer
	U   *entity.User
} {
	var calls []struct {
		Ctx context.Context
		Db  store.Execer
		U   *entity.User
	}
	mock.lockRegisterUser.RLock()
	calls = mock.calls.RegisterUser
	mock.lockRegisterUser.RUnlock()
	return calls
}

// Ensure, that UserGetterMock does implement UserGetter.
// If this is not the case, regenerate this file with moq.
var _ UserGetter = &UserGetterMock{}

// UserGetterMock is a mock implementation of UserGetter.
//
//	func TestSomethingThatUsesUserGetter(t *testing.T) {
//
//		// make and configure a mocked UserGetter
//		mockedUserGetter := &UserGetterMock{
//			GetUserFunc: func(ctx context.Context, db store.Queryer, name string) (*entity.User, error) {
//				panic("mock out the GetUser method")
//			},
//		}
//
//		// use mockedUserGetter in code that requires UserGetter
//		// and then make assertions.
//
//	}
type UserGetterMock struct {
	// GetUserFunc mocks the GetUser method.
	GetUserFunc func(ctx context.Context, db store.Queryer, name string) (*entity.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetUser holds details about calls to the GetUser method.
		GetUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db store.Queryer
			// Name is the name argument value.
			Name string
		}
	}
	lockGetUser sync.RWMutex
}

// GetUser calls GetUserFunc.
func (mock *UserGetterMock) GetUser(ctx context.Context, db store.Queryer, name string) (*entity.User, error) {
	if mock.GetUserFunc == nil {
		panic("UserGetterMock.GetUserFunc: method is nil but UserGetter.GetUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Db   store.Queryer
		Name string
	}{
		Ctx:  ctx,
		Db:   db,
		Name: name,
	}
	mock.lockGetUser.Lock()
	mock.calls.GetUser = append(mock.calls.GetUser, callInfo)
	mock.lockGetUser.Unlock()
	return mock.GetUserFunc(ctx, db, name)
}

// GetUserCalls gets all the calls that were made to GetUser.
// Check the length with:
//
//	len(mockedUserGetter.GetUserCalls())
func (mock *UserGetterMock) GetUserCalls() []struct {
	Ctx  context.Context
	Db   store.Queryer
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Db   store.Queryer
		Name string
	}
	mock.lockGetUser.RLock()
	calls = mock.calls.GetUser
	mock.lockGetUser.RUnlock()
	return calls
}

// Ensure, that TokenGeneratorMock does implement TokenGenerator.
// If this is not the case, regenerate this file with moq.
var _ TokenGenerator = &TokenGeneratorMock{}

// TokenGeneratorMock is a mock implementation of TokenGenerator.
//
//	func TestSomethingThatUsesTokenGenerator(t *testing.T) {
//
//		// make and configure a mocked TokenGenerator
//		mockedTokenGenerator := &TokenGeneratorMock{
//			GenerateTokenFunc: func(ctx context.Context, u entity.User) ([]byte, error) {
//				panic("mock out the GenerateToken method")
//			},
//		}
//
//		// use mockedTokenGenerator in code that requires TokenGenerator
//		// and then make assertions.
//
//	}
type TokenGeneratorMock struct {
	// GenerateTokenFunc mocks the GenerateToken method.
	GenerateTokenFunc func(ctx context.Context, u entity.User) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// GenerateToken holds details about calls to the GenerateToken method.
		GenerateToken []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// U is the u argument value.
			U entity.User
		}
	}
	lockGenerateToken sync.RWMutex
}

// GenerateToken calls GenerateTokenFunc.
func (mock *TokenGeneratorMock) GenerateToken(ctx context.Context, u entity.User) ([]byte, error) {
	if mock.GenerateTokenFunc == nil {
		panic("TokenGeneratorMock.GenerateTokenFunc: method is nil but TokenGenerator.GenerateToken was just called")
	}
	callInfo := struct {
		Ctx context.Context
		U   entity.User
	}{
		Ctx: ctx,
		U:   u,
	}
	mock.lockGenerateToken.Lock()
	mock.calls.GenerateToken = append(mock.calls.GenerateToken, callInfo)
	mock.lockGenerateToken.Unlock()
	return mock.GenerateTokenFunc(ctx, u)
}

// GenerateTokenCalls gets all the calls that were made to GenerateToken.
// Check the length with:
//
//	len(mockedTokenGenerator.GenerateTokenCalls())
func (mock *TokenGeneratorMock) GenerateTokenCalls() []struct {
	Ctx context.Context
	U   entity.User
} {
	var calls []struct {
		Ctx context.Context
		U   entity.User
	}
	mock.lockGenerateToken.RLock()
	calls = mock.calls.GenerateToken
	mock.lockGenerateToken.RUnlock()
	return calls
}
