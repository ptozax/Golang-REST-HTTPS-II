package models

type Constants struct {
	Result  bool    `json:"result"`
	Message message `json:"message"`
}

type message struct {
	Th string `json:"th"`
	En string `json:"en"`
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX TOKEN XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func Authentication_nil() Constants {
	return Constants{
		false,
		message{
			"ไม่มีข้อมูลการยืนยันตัวตน",
			"nil authentication",
		},
	}
}
func Token_invalid() Constants {
	return Constants{
		false,
		message{
			"โทเค็นไม่ถูกต้อง",
			"invalid token",
		},
	}
}

func Token_expired() Constants {
	return Constants{
		false,
		message{
			"โทเค็นหมดอายุ",
			"token expired",
		},
	}
}

func Token_not_match() Constants {
	return Constants{
		false,
		message{
			"โทเค็นไม่ตรงกัน",
			"token mismatch",
		},
	}
}

func Token_not_found() Constants {
	return Constants{
		false,
		message{
			"ไม่เจอโทเค็นนี้",
			"token not found",
		},
	}
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX CRUD XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func Insert_error() Constants {
	return Constants{
		false,
		message{
			"บันทึกข้อมูลผิดพลาด",
			"insert error",
		},
	}
}

func Insert_success() Constants {
	return Constants{
		true,
		message{
			"บันทึกข้อมูลสำเร็จ",
			"insert success",
		},
	}
}

func Edit_error() Constants {
	return Constants{
		false,
		message{
			"แก้ไขข้อมูลผิดพลาด",
			"edit data error",
		},
	}
}

func Edit_success() Constants {
	return Constants{
		true,
		message{
			"แก้ไขข้อมูลสำเร็จ",
			"edit data success",
		},
	}
}

func Delete_error() Constants {
	return Constants{
		false,
		message{
			"ลบข้อมูลผิดพลาด",
			"delete error",
		},
	}
}

func Delete_success() Constants {
	return Constants{
		true,
		message{
			"ลบข้อมูลสำเร็จ",
			"delete success",
		},
	}
}

func Get_Data_success() Constants {
	return Constants{
		true,
		message{
			"ดึงข้อมูลสำเร็จ",
			"get data success",
		},
	}
}

func Get_Data_Error() Constants {
	return Constants{
		false,
		message{
			"ดึงข้อมูลไม่สำเร็จ",
			"get data error",
		},
	}
}

func NotFound_OR_MoerThanOne() Constants {
	return Constants{
		false,
		message{
			"ไม่พบข้อมูล หรือ เจอมากกว่าหนึ่ง",
			"not found or moer than one",
		},
	}
}

func Data_Alardy_Added() Constants {
	return Constants{
		false,
		message{
			"มีข้อมูลนี้แล้ว",
			"data alardy added",
		},
	}
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX LOGIN LOGOUT SIGHUP  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
func Signup_success() Constants {
	return Constants{
		true,
		message{
			"สมัครใช้งานสำเร็จ",
			"Signup success",
		},
	}
}

func Signup_error() Constants {
	return Constants{
		false,
		message{
			"สมัครใช้งานไม่สำเร็จ",
			"Signup error",
		},
	}
}
func Logout_success() Constants {
	return Constants{
		true,
		message{
			"ออกจากระบบสำเร็จ",
			"logout success",
		},
	}
}

func Logout_error() Constants {
	return Constants{
		false,
		message{
			"ออกจากระบบไม่สำเร็จ",
			"logout error",
		},
	}
}

func Login_error() Constants {
	return Constants{
		false,
		message{
			"เข้าสู่ระบบไม่สำเร็จ",
			"login error",
		},
	}
}

func Username_Or_Email_AlreadyUsed() Constants {
	return Constants{
		false,
		message{
			"ชื่อผู้ช้งานหรืออีเมลนี้ถูกใช้ไปแล้ว",
			"username or email already used",
		},
	}
}

func Incorrect_Username_or_Password() Constants {
	return Constants{
		false,
		message{
			"ชื้อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
			"incorrect username or password",
		},
	}
}

func Baned_User() Constants {
	return Constants{
		false,
		message{
			"ผู้ใช้นี้ถูกระงับการใช้งาน",
			"user has been suspended",
		},
	}
}

//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX server  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

func Decode_error() Constants {
	return Constants{
		false,
		message{
			"ถอดรหัสไใาสำเร็จ",
			"decode error",
		},
	}
}
func Invalid_syntax() Constants {
	return Constants{
		false,
		message{
			"Syntax ไม่ถูกต้อง",
			"Invalid Syntax",
		},
	}
}
