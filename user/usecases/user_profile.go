package usecases

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-playground/validator/v10"
	"github.com/muhammedarifp/user/commonhelp/requests"
	"github.com/muhammedarifp/user/commonhelp/response"
	"github.com/muhammedarifp/user/config"
)

func (u *userUseCase) FetchUserProfileUsingID(userid string) (response.UserProfileValue, error) {
	userProfile, repoErr := u.userRepo.FetchUserProfileUsingID(userid)
	if repoErr != nil {
		return response.UserProfileValue{}, repoErr
	}

	return userProfile, nil
}

func (u *userUseCase) UpdateUserProfile(profile requests.UserPofileUpdate, userid string) (response.UserProfileValue, error) {
	// validate struct
	var userProfilrEmpty response.UserProfileValue
	if err := validator.New().Struct(&profile); err != nil {
		return userProfilrEmpty, err
	}

	if profile.Github != "" {
		if !govalidator.IsURL(profile.Github) {
			return userProfilrEmpty, errors.New("github link error")
		}
	}

	if profile.Linkedin != "" {
		if !govalidator.IsURL(profile.Linkedin) {
			return userProfilrEmpty, errors.New("linkedin link error")
		}
	}

	//
	userProfile, repoErr := u.userRepo.UpdateUserProfile(profile, userid)
	if repoErr != nil {
		return userProfilrEmpty, repoErr
	}

	return userProfile, nil
}

func (u *userUseCase) UploadNewProfilePhoto(photo multipart.File, header multipart.FileHeader, userid string) (response.UserProfileValue, error) {
	var userProfileEmpty response.UserProfileValue
	cfg := config.GetConfig()
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(cfg.AWS_ACCESS_KEYID, cfg.AWS_SECRET_ACCESS_KEY, ""),
	}))
	filename_slice := strings.Split(header.Filename, ".")
	ext := filename_slice[len(filename_slice)-1]
	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.BUCKET_NAME),
		Key:    aws.String("profile/" + userid + "." + ext),
		ACL:    aws.String("public-read"),
		Body:   photo,
	})

	if err != nil {
		return userProfileEmpty, err
	}

	userVal, usecaseErr := u.userRepo.UploadProfileImage(result.Location, userid)
	if usecaseErr != nil {
		return userProfileEmpty, usecaseErr
	}

	return userVal, nil
}
