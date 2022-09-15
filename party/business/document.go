package business

import (
	"bytes"
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"image"
	"image/png"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-uuid"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/domain"
	pb "github.com/kkjhamb01/courier-management/party/proto"
	"github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"github.com/nfnt/resize"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Service) Upload(ctx context.Context, in *pb.UploadDocumentRequest) (*pb.UploadDocumentResponse, error) {
	logger.Infof("Upload infoType = %v, docType = %v",
		in.GetDocument().GetInfoType(), in.GetDocument().GetDocType())

	var validDocType = in.GetDocument().GetInfoType().AdditionalInfoType().ValidDocumentTypes()

	var docType = in.GetDocument().GetDocType()
	var isValid = false
	for _, valid := range validDocType {
		if valid == docType {
			isValid = true
			break
		}
	}

	if !isValid {
		logger.Debugf("Upload invalid document type %v", in.GetDocument().DocType)
		return nil, proto.InvalidArgument.ErrorMsg(fmt.Sprintf("invalid document type %v", in.GetDocument().DocType))
	}

	tokenUser := ctx.Value("user").(security.User)

	objectId, _ := uuid.GenerateUUID()

	err := s.db.Transaction(func(tx *gorm.DB) error {

		if in.GetDocument().GetInfoType() == pb.DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD {
			// delete pair id
			if in.GetDocument().GetDocType() == pb.DocumentType_DOCUMENT_TYPE_PASSPORT {
				if err := tx.Where("user_id = ? AND document_info_type = ?", tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD)).
					Delete(domain.Document{}).Error; err != nil {
					return proto.Internal.Error(err)
				}
			} else {
				if err := tx.Where("user_id = ? AND document_type = ?", tokenUser.Id, int(pb.DocumentType_DOCUMENT_TYPE_PASSPORT)).
					Delete(domain.Document{}).Error; err != nil {
					return proto.Internal.Error(err)
				}
			}
		}

		// delete existing documents
		if err := tx.Where("user_id = ? AND document_type = ?", tokenUser.Id, int(in.Document.DocType)).
			Delete(domain.Document{}).Error; err != nil {
			return proto.Internal.Error(err)
		}

		// insert document
		if err := tx.Create(&domain.Document{
			UserID:           tokenUser.Id,
			ObjectId:         objectId,
			DocumentInfoType: int(in.Document.InfoType),
			DocumentType:     int(in.Document.DocType),
			FileType:         sql.NullString{String: in.Document.FileType, Valid: in.Document.FileType != ""},
			CreationTime:     time.Now(),
		}).Error; err != nil {
			return proto.Internal.Error(err)
		}

		// insert document data
		if err := tx.Create(&domain.DocumentData{
			ObjectId: objectId,
			Data:     in.Document.Data,
		}).Error; err != nil {
			return proto.Internal.Error(err)
		}

		return nil
	})

	if err != nil {
		logger.Errorf("Upload error in uploading document", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("Upload uploaded successfully userId = %v, objectId = %v", tokenUser.Id, objectId)

	var infoType = in.GetDocument().GetInfoType().AdditionalInfoType()

	if infoType != pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		newStatus, err := s.profileInfoStatus(tokenUser.Id, infoType)
		if err != nil {
			return nil, proto.Internal.Error(err)
		}
		if newStatus != pb.ProfileAdditionalInfoStatus_UNKNOWN_PROFILE_ADDITIONAL_INFO_STATUS {
			err = s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "status_type"}},
				DoUpdates: clause.AssignmentColumns([]string{"status", "message"}),
			}).Create(&domain.CourierStatus{
				UserID:     tokenUser.Id,
				StatusType: int32(infoType),
				Status:     int32(newStatus),
				Message:    "",
			}).Error

			if err != nil {
				return nil, proto.Internal.Error(err)
			}
		}
	}

	return &pb.UploadDocumentResponse{
		ObjectId: objectId,
	}, err

}

func (s *Service) GetDocumentsOfUser(ctx context.Context, in *pb.GetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error) {
	tokenUser := ctx.Value("user").(security.User)
	return s.getDocumentsOfUserById(tokenUser.Id, in.GetType(), in.GetDataType())

}

func (s *Service) getDocumentsOfUserById(userId string, docType pb.DocumentInfoType, dataType pb.DocumentDataType) (*pb.GetDocumentsOfUserResponse, error) {
	logger.Infof("getDocumentsOfUserById userId = %v, type = %v", userId, docType)

	var documents []domain.Document
	var err error
	if docType == pb.DocumentInfoType_UNKNOWN_DOCUMENT_INFO_TYPE {
		err = s.db.Model(&domain.Document{}).Where("user_id=?", userId).Scan(&documents).Error
	} else {
		err = s.db.Model(&domain.Document{}).Where("user_id=? AND document_info_type=?", userId, docType).Scan(&documents).Error
	}
	if err != nil {
		logger.Errorf("getDocumentsOfUserById cannot find documents", err)
		return nil, proto.Internal.Error(err)
	}

	if len(documents) == 0 {
		return nil, proto.NotFound.ErrorMsg("no documents are of this user for this type")
	}

	var documentInfoList = make([]*pb.DocumentInfo, len(documents))

	for i, document := range documents {
		documentInfoList[i] = &pb.DocumentInfo{
			InfoType: pb.DocumentInfoType(document.DocumentInfoType),
			DocType:  pb.DocumentType(document.DocumentType),
			FileType: document.FileType.String,
			ObjectId: document.ObjectId,
		}
		if dataType == pb.DocumentDataType_DOCUMENT_DATA_TYPE_DATA ||
			dataType == pb.DocumentDataType_DOCUMENT_DATA_TYPE_THUMBNAIL {
			documentData := domain.DocumentData{}

			err = s.db.Model(&domain.DocumentData{}).Where("object_id=?", document.ObjectId).Find(&documentData).Error

			if err != nil {
				logger.Errorf("getDocumentsOfUserById cannot find document data", err)
				return nil, proto.Internal.Error(err)
			}

			var data = documentData.Data

			if dataType == pb.DocumentDataType_DOCUMENT_DATA_TYPE_THUMBNAIL {
				img := &Image{
					Data: documentData.Data,
				}
				img, err = img.thumbnail(s.config.Thumbnail.MaxDimension)
				if err != nil {
					logger.Errorf("getDocumentsOfUserById cannot create thumbnail", err)
					return nil, proto.Internal.Error(err)
				}
				data = img.Data
			}

			documentInfoList[i].Data = data
		}
	}

	return &pb.GetDocumentsOfUserResponse{
		Documents: documentInfoList,
	}, nil

}

func (s *Service) GetDocument(ctx context.Context, in *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	tokenUser := ctx.Value("user").(security.User)
	logger.Debugf("GetDocument userId = %v, objectId = %v ", tokenUser.Id, in.ObjectId)

	document := domain.Document{}

	err := s.db.Model(&domain.Document{}).Where("user_id=? AND object_id=?", tokenUser.Id, in.ObjectId).Find(&document).Error

	if err != nil {
		logger.Errorf("GetDocument cannot query document", err)
		return nil, proto.Internal.Error(err)
	}

	if document.UserID == "" {
		return nil, proto.NotFound.ErrorNoMsg()
	}

	result := &pb.GetDocumentResponse{
		DownloadLinkExpiration: s.config.Download.Expiration + time.Now().Unix(),
		DownloadLink:           s.generateDownloadLink(in.ObjectId),
	}

	logger.Debugf("GetDocument result = %v ", result)

	return result, nil

}

func (s *Service) generateDownloadLink(objectId string) string {
	exp := strconv.FormatInt(s.config.Download.Expiration+time.Now().Unix(), 10)
	hash := md5.Sum([]byte(objectId + "|" + exp + "|" + s.config.Download.ExpirationSecretKey))
	hashStr := hex.EncodeToString(hash[:])
	return "object=" + objectId + "&exp=" + exp + "&h=" + hashStr
}

func (s *Service) Download(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	logger.Debugf("Download %v ", in.DownloadLink)

	queryParams := strings.Split(in.DownloadLink, "&")
	if len(queryParams) != 3 {
		logger.Debug("Download invalid query params")
		return nil, proto.InvalidArgument.ErrorMsg("invalid download link")
	}

	queryParamsObject := strings.Split(queryParams[0], "=")
	if len(queryParamsObject) != 2 {
		logger.Debug("Download invalid query params of object")
		return nil, proto.InvalidArgument.ErrorMsg("invalid download link")
	}

	queryParamsExp := strings.Split(queryParams[1], "=")
	if len(queryParamsExp) != 2 {
		logger.Debug("Download invalid query params of exp")
		return nil, proto.InvalidArgument.ErrorMsg("invalid download link")
	}

	queryParamsHash := strings.Split(queryParams[2], "=")
	if len(queryParamsHash) != 2 {
		logger.Debug("Download invalid query params of hash")
		return nil, proto.InvalidArgument.ErrorMsg("invalid download link")
	}

	objectId := queryParamsObject[1]
	exp, err := strconv.ParseInt(queryParamsExp[1], 10, 64)
	if err != nil {
		logger.Errorf("Download invalid exp", err)
		return nil, proto.InvalidArgument.ErrorMsg("invalid download link")
	}
	hash := queryParamsHash[1]

	now := time.Now().Unix()
	if now > exp {
		logger.Debugf("Download link is expired")
		return nil, proto.InvalidArgument.ErrorMsg("link is expired")
	}

	hashToBe := md5.Sum([]byte(objectId + "|" + strconv.FormatInt(exp, 10) + "|" + s.config.Download.ExpirationSecretKey))
	hashStr := hex.EncodeToString(hashToBe[:])
	if hashStr != hash {
		logger.Debugf("Download invalid hash")
		return nil, proto.InvalidArgument.ErrorMsg("invalid download link")
	}

	var documentData domain.DocumentData
	err = s.db.Model(&domain.DocumentData{}).Where("object_id=?", objectId).Find(&documentData).Error

	if err != nil {
		logger.Errorf("Download cannot find document", err)
		return nil, proto.Internal.Error(err)
	}

	if documentData.ObjectId == "" {
		logger.Debugf("Download document not found %v", objectId)
		return nil, proto.NotFound.ErrorNoMsg()
	}

	return &pb.DownloadResponse{
		Data: documentData.Data,
	}, nil

}

func (s *Service) DirectDownload(ctx context.Context, in *pb.DirectDownloadRequest) (*pb.DirectDownloadResponse, error) {
	logger.Infof("DirectDownload objectId = %v", in.GetObjectId())

	document := domain.Document{}
	err := s.db.Model(&domain.Document{}).Where("object_id=?", in.ObjectId).Find(&document).Error

	if err != nil {
		logger.Errorf("DirectDownload cannot find document", err)
		return nil, proto.Internal.Error(err)
	}

	if document.UserID == "" {
		return nil, proto.NotFound.ErrorNoMsg()
	}

	documentData := domain.DocumentData{}
	err = s.db.Model(&domain.DocumentData{}).Where("object_id=?", in.ObjectId).Find(&documentData).Error

	if err != nil {
		logger.Errorf("DirectDownload cannot find document", err)
		return nil, proto.Internal.Error(err)
	}

	if documentData.ObjectId == "" {
		return nil, proto.NotFound.ErrorNoMsg()
	}

	return &pb.DirectDownloadResponse{
		InfoType: pb.DocumentInfoType(document.DocumentInfoType),
		DocType:  pb.DocumentType(document.DocumentType),
		FileType: document.FileType.String,
		Data:     documentData.Data,
	}, nil

	return nil, nil
}

func (i *Image) thumbnail(maxDim int) (*Image, error) {
	img, _, err := image.Decode(bytes.NewReader(i.Data))

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	d := float64(width) / float64(height)

	if d > 1 {
		height = int(float64(maxDim) / d)
		width = maxDim
	} else {
		width = int(float64(maxDim) * d)
		height = maxDim
	}

	thumbnail := resize.Thumbnail(uint(width), uint(height), img, resize.Lanczos3)

	data := new(bytes.Buffer)

	err = png.Encode(data, thumbnail)

	if err != nil {
		return nil, err
	}

	bs := data.Bytes()

	t := &Image{
		Data: bs,
	}

	return t, nil
}

type Image struct {
	Data []byte
}
