package usecase

import (
	"go-clean-architecture-firstapp/model"
	"go-clean-architecture-firstapp/repository"
	"go-clean-architecture-firstapp/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// user_usecaseに必要な、メソッドが、
// signUpとsignInの2である理由がわからない
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error) // signUpの返り値
	Login(user model.User) (string, error)              // Loginの返り値
}

// ! urがよくわからない
type userUsecase struct {
	// urは、フィールド
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

// signUPメソッド
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {

	// !バリデーション
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	// !登録実装
	// signupの引数であるuserには、クライアントがわから渡ってきた入力データが入る
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// byte(user.Password)をハッシュ化
	if err != nil {
		return model.UserResponse{}, err
		// 返値では、UserResponseの後ろに空{}を記述しないといけない
	}
	// user.Passwordで渡ってきた文字列をbcryptで暗号化して、newUserのハッシュ化
	newUser := model.User{Email: user.Email, Password: string(hash)}
	// ! := は、変数宣言と代入を同時に行う。つまり var hoge = fuga と同じ
	// var newUser = model.User{Email: user.Email, Password: string(hash)}

	// CreateUserを呼美出す
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	// user_repositoryで、レスポンスを構造として持ってる。返却値に値を代入してる
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	// !バリデーション
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	// !登録
	storedUser := model.User{}
	// db側で持ってる情報と入力された情報の照らし合わせ
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	// パスワードが一致したらjwtトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // アクセストークン時間
	})
	tokenString, err := token.SignedString(([]byte(os.Getenv("SECRET")))) // SignedStringでjwtトークンを生成できる
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
