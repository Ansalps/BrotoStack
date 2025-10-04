package repo

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ansalps/BrotoStack/pkg/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}
func (r *userRepo) Store_Unverified_User(signuprequest models.UserSignUpRequest) error {
	user := models.Users{
		Username:     signuprequest.Username,
		Email:        signuprequest.Email,
		PasswordHash: signuprequest.Confirmpassword,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func(r *userRepo)CheckIfEmailExistsInOtp(email string)(bool,error){
	var exists bool
	err:=r.db.Raw(`select exists(select 1 from otps where email =?)`,email).Scan(&exists).Error
	if err!=nil{
		return  exists,err
	}
	return exists,nil
}

func (r *userRepo) Store_Otp_For_User(Otp, email string) error {
	var otp models.Otps
	otp.Email = email
	otp.Data = Otp
	if err := r.db.Create(&otp).Error; err != nil {
		return err
	}
	return nil
}
func (r *userRepo) Overwrite_Otp_To_Email(Otp,email string)error{
	err:=r.db.Model(&models.Otps{}).Where("email=?",email).Update("data",Otp).Error
	if err!=nil{
		return err
	}
	return nil
}
func (r *userRepo)Delete_Unverified_User_With_Same_Email(email string) error{
	query1 := `delete from users where email=? and is_verified=false`
	err := r.db.Exec(query1, email).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Does_Email_Exist_In_DB(email string) (bool,error) {
	var exists bool
	query1:=`select exists(select 1 from users where email=? and is_verified=true)`
	err:=r.db.Raw(query1,email).Scan(&exists).Error
	if err!=nil{
		fmt.Println("is there any error in does email exi")
		return exists,err
	}
	return exists,err
}

func (r *userRepo) Does_Username_Exist_In_DB(usename string) error {
	query := `select  username from users where username = ?`
	var Username string
	err := r.db.Raw(query, usename).Scan(&Username).Error
	if err != nil {
		return err
	}
	if Username != "" {
		return fmt.Errorf("username %s already exists choose another username!", usename)
	}
	return nil
}

func (r *userRepo) RemoveUnverifiedUsersOlderThan3Minutes(age time.Duration) error {
	cutoffTime := time.Now().Add(-age)
	query := `delete from users where updated_at< ? and is_verified=false`
	err := r.db.Exec(query, cutoffTime).Error
	if err != nil {
		return err
	}
	query2:=`delete from otps where updated_at<? or is_valid=false`
	err=r.db.Exec(query2,cutoffTime).Error
	if err!=nil{
		return err
	}
	return nil
}


func (r *userRepo)CheckForUnverifiedUserInDB(email string)error  {
	var emailExists bool
	err:=r.db.Raw(`select exists (select email from users where email=? and is_verified=false)`,email).Scan(&emailExists).Error
	if err!=nil{
		return err
	}
	if !emailExists{
		return errors.New("try to sign up before otp verification")
	}
	return nil
}

func (r *userRepo)CheckIfOtpExists(otp,email string)(string,error){
	var OTP string
	err := r.db.Raw(`select data from otps where data =? and email=?`, otp, email).Scan(&OTP).Error
	if err != nil {
		return "",err
	}
	if OTP==""{
		return "",errors.New("otp is not stored in database")
	}
	return OTP,nil
}
func (r *userRepo)CheckIfOtpExpired(otp,email string) error{
	var exists bool
	var updated_at,expired_at time.Time
	r.db.Raw(`select updated_at from otps where data=? and email=?`,otp,email).Scan(&updated_at)
	fmt.Println("created_at ---",updated_at)
	r.db.Raw(`select NOW()-INTERVAL '3 minutes'`).Scan(&expired_at)
	fmt.Println("expired---",expired_at)
		err:=r.db.Raw(`select exists ( select 1 from otps where updated_at < NOW()-INTERVAL '3 minutes' and data=? and email=?)`,otp,email).Scan(&exists).Error
		if err!=nil{
			return err
		}
		if exists{
			return errors.New("otp is expired, and more than 3 minutes old.")
		}
		return nil
}
func(r *userRepo)VerifyUser(email string)error{
	err := r.db.Exec(`update users set is_verified=true where email=?`, email).Error
		if err != nil {
		return err
	}
	return nil
}
func (r *userRepo)InvalidateOtp(otp,email string)error{
	err:=r.db.Exec(`update otps set is_valid=false where data=? and email=?`,otp,email).Error
	if err!=nil{
		return err
	}
	return nil
}
func(r *userRepo)ResetPassword(email,password string)error{
	err:=r.db.Exec(`update users set password_hash = ? where email = ? and is_verified=true`,password,email).Error
	if err!=nil{
		return err
	}
	return nil
}

func (r *userRepo)FetchStoredHashFromExistingUser(email string)(string,error){
	var passwordHash string
	err:=r.db.Raw(`select password_hash from users where email=?`,email).Scan(&passwordHash).Error
	if err!=nil{
		return "",err
	}
	fmt.Println("is there error here",err)
	return passwordHash,nil
}
func (r *userRepo)FetchDetailsForExistingUser(email string)(models.Users,error){
	var user models.Users
	err:=r.db.Raw(`select * from users where email=?`,email).Scan(&user).Error
	if err!=nil{
		return user,err
	}
	fmt.Println("is there error here",err)
	return user,nil
}