package models

type Constants struct {
	Result  bool    `json:"result"`
	Message message `json:"message"`
}

type message struct {
	Th string `json:"th"`
	En string `json:"en"`
	Bu string `json:"bu"`
}

func Invalid_syntax() Constants {
	return Constants{
		false,
		message{
			"Syntax ไม่ถูกต้อง",
			"Invalid Syntax",
			"",
		},
	}
}

func User_not_found() Constants {
	return Constants{
		false,
		message{
			"ไม่เจอผู้ใช้งานนี้",
			"user not found",
			"",
		},
	}
}

func Password_Incorrect() Constants {
	return Constants{
		false,
		message{
			"พาสเวิร์ดไม่ถูกต้อง",
			"password incorrect",
			"",
		},
	}
}

func Email_invalid() Constants {
	return Constants{
		false,
		message{
			"อีเมลล์หรือพาสเวิร์ดไม่ถูกต้อง",
			"Invalid e-mail or password",
			"",
		},
	}
}

func Invalid_token() Constants {
	return Constants{
		false,
		message{
			"โทเค็นไม่ถูกต้อง",
			"invalid token",
			"",
		},
	}
}

func Token_expired() Constants {
	return Constants{
		false,
		message{
			"โทเค็นหมดอายุ",
			"token expired",
			"",
		},
	}
}

func Token_not_match() Constants {
	return Constants{
		false,
		message{
			"โทเค็นไม่ตรงกัน",
			"token mismatch",
			"",
		},
	}
}

func Token_not_found() Constants {
	return Constants{
		false,
		message{
			"ไม่เจอโทเค็นนี้",
			"token not found",
			"",
		},
	}
}

func Logout_success() Constants {
	return Constants{
		true,
		message{
			"ออกจากระบบสำเร็จ",
			"logout success",
			"",
		},
	}
}

func Get_data_error() Constants {
	return Constants{
		false,
		message{
			"ส่งข้อมูลผิดพลาด",
			"sent data error",
			"",
		},
	}
}

func Username_Pass() Constants {
	return Constants{
		true,
		message{
			"ชื่อผู้ใช้นี้สามารถใช้งานได้",
			"Username is available",
			"",
		},
	}
}

func Username_Not_Pass() Constants {
	return Constants{
		true,
		message{
			"ชื่อผู้ใช้นี้ถูกใช้งานไปแล้ว",
			"Username is already in use",
			"",
		},
	}
}

func Get_data_success() Constants {
	return Constants{
		true,
		message{
			"รับข้อมูลถูกต้อง",
			"get data success",
			"",
		},
	}
}

func Delete_picture_success() Constants {
	return Constants{
		true,
		message{
			"ลบรูปภาพสำเร็จ",
			"delete picture success",
			"",
		},
	}
}

func Save_picture_error() Constants {
	return Constants{
		false,
		message{
			"บันทึกรูปภาพผิดพลาด",
			"save picture error",
			"",
		},
	}
}

func Insert_error() Constants {
	return Constants{
		false,
		message{
			"บันทึกข้อมูลผิดพลาด",
			"insert error",
			"",
		},
	}
}

func Insert_success() Constants {
	return Constants{
		true,
		message{
			"บันทึกข้อมูลสำเร็จ",
			"insert success",
			"",
		},
	}
}

func Update_error() Constants {
	return Constants{
		false,
		message{
			"บันทึกข้อมูลผิดพลาด",
			"save data error",
			"",
		},
	}
}

func File_error() Constants {
	return Constants{
		false,
		message{
			"บันทึกไฟล์ผิดพลาด",
			"save file error",
			"",
		},
	}
}

func Image_Invalid() Constants {
	return Constants{
		false,
		message{
			"รูปภาพไม่ถูกต้อง",
			"image invalid",
			"",
		},
	}
}

func EmailAlreadyUsed() Constants {
	return Constants{
		false,
		message{
			"อีเมลล์นี้ถูกใช้ไปแล้ว",
			"email already used",
			"",
		},
	}
}

func Username_Or_Email_AlreadyUsed() Constants {
	return Constants{
		false,
		message{
			"ชื่อผู้ช้งานหรืออีเมลนี้ถูกใช้ไปแล้ว",
			"username or email already used",
			"",
		},
	}
}

func Data_not_found() Constants {
	return Constants{
		false,
		message{
			"ไม่พบข้อมูล",
			"data not found",
			"",
		},
	}
}

func Update_success() Constants {
	return Constants{
		true,
		message{
			"บันทึกข้อมูลสำเร็จ",
			"save data success",
			"",
		},
	}
}

func Delete_file_error() Constants {
	return Constants{
		false,
		message{
			"ลบไฟล์ล้มเหลว",
			"delete file error",
			"",
		},
	}
}
func Delete_file_success() Constants {
	return Constants{
		true,
		message{
			"ลบไฟล์สำเร็จ",
			"delete file success",
			"",
		},
	}
}

func Edit_error() Constants {
	return Constants{
		false,
		message{
			"แก้ไขข้อมูลผิดพลาด",
			"edit data error",
			"",
		},
	}
}

func Edit_success() Constants {
	return Constants{
		true,
		message{
			"แก้ไขข้อมูลสำเร็จ",
			"edit data success",
			"",
		},
	}
}

func Delete_error() Constants {
	return Constants{
		false,
		message{
			"ลบข้อมูลผิดพลาด",
			"delete error",
			"",
		},
	}
}

func Delete_success() Constants {
	return Constants{
		true,
		message{
			"ลบข้อมูลสำเร็จ",
			"delete success",
			"",
		},
	}
}

func Create_token_error() Constants {
	return Constants{
		true,
		message{
			"สร้างโทเค็นผิดพลาด",
			"create token error",
			"",
		},
	}
}

func Change_password_success() Constants {
	return Constants{
		true,
		message{
			"เปลี่ยนพาสเวิร์ดสำเร็จ",
			"change password success",
			"",
		},
	}
}

func Change_password_error() Constants {
	return Constants{
		false,
		message{
			"เปลี่ยนพาสเวิร์ดผิดพลาด",
			"change password error",
			"",
		},
	}
}

func Password_not_match() Constants {
	return Constants{
		false,
		message{
			"พาสเวิร์ดไม่ตรงกัน",
			"Password not match",
			"",
		},
	}
}

func Member_not_found() Constants {
	return Constants{
		false,
		message{
			"ไม่พบสมาชิก",
			"member not found",
			"",
		},
	}
}

func Signup_success() Constants {
	return Constants{
		true,
		message{
			"สมัครใช้งานสำเร็จ",
			"Signup success",
			"",
		},
	}
}
