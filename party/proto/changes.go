package proto

func (x *IDCard) SetDocumentIds(documentIds []*DocumentInfo) {
	x.DocumentIds = documentIds
}

func (x *IDCard) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD)
}

func (x *DrivingLicense) SetDocumentIds(documentIds []*DocumentInfo) {
	x.DocumentIds = documentIds
}

func (x *DrivingLicense) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVING_LICENSE)
}

func (x *DriverBackground) SetDocumentIds(documentIds []*DocumentInfo) {
	x.DocumentIds = documentIds
}

func (x *DriverBackground) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVER_BACKGROUND)
}

func (x *ResidenceCard) SetDocumentIds(documentIds []*DocumentInfo) {
	x.DocumentIds = documentIds
}

func (x *ResidenceCard) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_RESIDENCE_CARD)
}

func (x *BankAccount) SetDocumentIds(documentIds []*DocumentInfo) {
	x.DocumentIds = documentIds
}

func (x *BankAccount) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_BANK_ACCOUNT)
}

func (x *Address) SetDocumentIds(documentIds []*DocumentInfo) {
	x.DocumentIds = documentIds
}

func (x *Address) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_ADDRESS)
}

func (x *InsuranceCertificate) SetDocumentIds(documentIds []*DocumentInfo) {
	if len(documentIds) > 0 {
		x.ObjectId = documentIds[0].ObjectId
		x.Data = documentIds[0].Data
	}
}

func (x *InsuranceCertificate) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_INSURANCE_CERTIFICATE)
}

func (x *InsuranceCertificate) GetDocumentIds() []*DocumentInfo{
	if x.ObjectId == ""{
		return nil
	}
	return []*DocumentInfo{{
		ObjectId: x.ObjectId,
		Data:     x.Data,
	}}
}

func (x *ProfilePicture) SetDocumentIds(documentIds []*DocumentInfo) {
	if len(documentIds) > 0 {
		x.ObjectId = documentIds[0].ObjectId
		x.Data = documentIds[0].Data
	}
}

func (x *ProfilePicture) GetDocumentInfoType() int32{
	return int32(DocumentInfoType_DOCUMENT_INFO_TYPE_PROFILE_PICTURE)
}

func (x *ProfilePicture) GetDocumentIds() []*DocumentInfo{
	if x.ObjectId == ""{
		return nil
	}
	return []*DocumentInfo{{
		ObjectId: x.ObjectId,
		Data:     x.Data,
	}}
}

func (x *Boolean) ToInt() int32{
	switch *x {
	case Boolean_BOOLEAN_FALSE:
		return 1
	case Boolean_BOOLEAN_TRUE:
		return 2
	default:
		return 0
	}
}

func (x *Boolean) FromInt(i int32) Boolean{
	switch i {
	case 1:
		return Boolean_BOOLEAN_FALSE
	case 2:
		return Boolean_BOOLEAN_TRUE
	default:
		return Boolean_UNKNOWN_BOOLEAN
	}
}

func (x *Boolean) Valid() bool{
	if *x == Boolean_UNKNOWN_BOOLEAN{
		return false
	}
	return true
}

func (x *Boolean) ToBool() bool {
	switch *x {
	case Boolean_BOOLEAN_FALSE:
		return false
	case Boolean_BOOLEAN_TRUE:
		return true
	default:
		return false
	}
}

func (x DocumentInfoType) AdditionalInfoType() AdditionalInfoType {
	switch x {
	case DocumentInfoType_DOCUMENT_INFO_TYPE_PROFILE_PICTURE:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE
	case DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD
	case DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVING_LICENSE:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE
	case DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVER_BACKGROUND:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND
	case DocumentInfoType_DOCUMENT_INFO_TYPE_RESIDENCE_CARD:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD
	case DocumentInfoType_DOCUMENT_INFO_TYPE_INSURANCE_CERTIFICATE:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE
	case DocumentInfoType_DOCUMENT_INFO_TYPE_ADDRESS:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS
	case DocumentInfoType_DOCUMENT_INFO_TYPE_BANK_ACCOUNT:
		return AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT
	}
	return AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE
}

func (x DocumentType) InArray(array []DocumentType) bool {
	for _,a := range array{
		if x == a{
			return true
		}
	}
	return false
}

func (x AdditionalInfoType) DocumentInfoType() DocumentInfoType{
	switch x {
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_PROFILE_PICTURE
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVING_LICENSE
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVER_BACKGROUND
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_RESIDENCE_CARD
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_INSURANCE_CERTIFICATE
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_ADDRESS
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT:
		return DocumentInfoType_DOCUMENT_INFO_TYPE_BANK_ACCOUNT
	}
	return DocumentInfoType_UNKNOWN_DOCUMENT_INFO_TYPE
}

func (x AdditionalInfoType) ValidDocumentTypes() []DocumentType {
	switch x {
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_PROFILE_PICTURE,
		}
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_PASSPORT,
			DocumentType_DOCUMENT_TYPE_NATIONAL_ID_FRONT,
			DocumentType_DOCUMENT_TYPE_NATIONAL_ID_BACK,
		}
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_DRIVING_LICENSE_FRONT,
			DocumentType_DOCUMENT_TYPE_DRIVING_LICENSE_BACK,
		}
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_DBS_CERTIFICATE_FRONT,
			DocumentType_DOCUMENT_TYPE_DBS_CERTIFICATE_BACK,
		}
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_RESIDENCE_CARD_FRONT,
			DocumentType_DOCUMENT_TYPE_RESIDENCE_CARD_BACK,
		}
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_INSURANCE_CERTIFICATE,
		}
	case AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS:
		return []DocumentType{
			DocumentType_DOCUMENT_TYPE_PROOF_OF_ADDRESS,
		}
	}
	return nil
}