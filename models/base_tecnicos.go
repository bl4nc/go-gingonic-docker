package models

import (
	"gorm.io/gorm"
)

type Base_Tecnicos struct {
	//*gorm.Model
	Id_tecnico    int    `json:"id_tecnico" gorm:"primary_key"`
	Login_tecnico string `json:"login_tecnico"`
	Nome_tecnico  string `json:"nome_tecnico"`
	Email_lideres string `json:"email_lideres" gorm:"type:text"`
}

type Todos_Tecnicos struct {
	//*gorm.Model
	Id_tecnico    int    `json:"id_tecnico" gorm:"primary_key"`
	Login_tecnico string `json:"login_tecnico"`
	Nome_tecnico  string `json:"nome_tecnico"`
}

func GetTecnicoId(db *gorm.DB, Base_Tecnicos *Base_Tecnicos, id_tecnico string) (err error) {
	err = db.Raw("SELECT nome_tecnico,id_tecnico,login_tecnico FROM tabela_tecnicos group by login_tecnico;").Scan(&Base_Tecnicos).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTecnicoLogin(db *gorm.DB, Base_Tecnicos *Base_Tecnicos, login_tecnico string) (err error) {
	err = db.Raw("SELECT id_tecnico,login_tecnico,nome_tecnico, group_concat(DISTINCT email_lider) as email_lideres FROM tabela_tecnicos  where login_tecnico = ?;", login_tecnico).Scan(&Base_Tecnicos).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTecnicos(db *gorm.DB, Todos_Tecnicos *[]Todos_Tecnicos) (err error) {
	if err := db.Raw("SELECT nome_tecnico,id_tecnico,login_tecnico FROM tabela_tecnicos group by login_tecnico;").Scan(&Todos_Tecnicos).Error; err != nil {
		return err
	}
	return nil
}
