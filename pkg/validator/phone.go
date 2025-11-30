package validator

import (
	"manager/pkg/logger"

	"github.com/nyaruka/phonenumbers"
)

func IsPhoneValid(phone string) bool {
	// Phân tích số điện thoại, cung cấp "VN" làm mã vùng mặc định.
	// Nếu người dùng nhập số như "09...", thư viện sẽ hiểu đó là số của Việt Nam.
	parsedNumber, err := phonenumbers.Parse(phone, "VN")
	if err != nil {
		// Lỗi này xảy ra nếu chuỗi không có định dạng của một số điện thoại.
		logger.Debugf("Can not analyst number '%s': %v\n", phone, err)
		return false
	}

	// Kiểm tra xem số đã phân tích có hợp lệ theo quy tắc của khu vực đó không.
	isValid := phonenumbers.IsValidNumber(parsedNumber)

	if isValid {
		// Lấy định dạng E.164 (ví dụ: +84987654321) để xem kết quả.
		formattedNumber := phonenumbers.Format(parsedNumber, phonenumbers.E164)
		logger.Debugf("Number %v correct: %s\n", phone, formattedNumber)
		return true
	}

	logger.Debugf("Number '%s' is not correct.\n", phone)
	return false
}
