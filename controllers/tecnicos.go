package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"tecnicos_service/database"
	"tecnicos_service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TecnicosRepo struct {
	Db *gorm.DB
}

func New() *TecnicosRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Base_Tecnicos{})
	return &TecnicosRepo{Db: db}
}

func (repository *TecnicosRepo) GetTecnicoId(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var tecnico models.Base_Tecnicos
	reply, err := database.GetRedis("tecnico_" + id)
	if err != nil {
		err := models.GetTecnicoId(repository.Db, &tecnico, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		email_lideres := (strings.Split(tecnico.Email_lideres, " "))
		m, err := json.Marshal(tecnico)
		database.SetRedis("tecnico_"+id, []byte(m))
		c.JSON(http.StatusOK, gin.H{"id_tecnico": tecnico.Id_tecnico,
			"nome_tecnico":  tecnico.Nome_tecnico,
			"login_tecnico": tecnico.Login_tecnico,
			"email_lideres": email_lideres})
	} else {
		var data models.Base_Tecnicos
		if err := json.Unmarshal(reply, &data); err != nil {
			panic(err)
		}
		email_lideres := (strings.Split(data.Email_lideres, " "))
		c.JSON(http.StatusOK, gin.H{"id_tecnico": data.Id_tecnico,
			"nome_tecnico":  data.Nome_tecnico,
			"login_tecnico": data.Login_tecnico,
			"email_lideres": email_lideres})
	}
}

func (repository *TecnicosRepo) GetTecnicoLogin(c *gin.Context) {
	login, _ := c.Params.Get("login")
	var tecnico models.Base_Tecnicos
	reply, err := database.GetRedis("tecnico_" + login)
	if err != nil {
		err := models.GetTecnicoLogin(repository.Db, &tecnico, login)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		email_lideres := (strings.Split(tecnico.Email_lideres, ","))
		m, err := json.Marshal(tecnico)
		database.SetRedis("tecnico_"+login, []byte(m))
		c.JSON(http.StatusOK, gin.H{"id_tecnico": tecnico.Id_tecnico,
			"nome_tecnico":  tecnico.Nome_tecnico,
			"login_tecnico": tecnico.Login_tecnico,
			"email_lideres": email_lideres})
	} else {
		var data models.Base_Tecnicos
		if err := json.Unmarshal(reply, &data); err != nil {
			panic(err)
		}
		email_lideres := (strings.Split(data.Email_lideres, ","))
		c.JSON(http.StatusOK, gin.H{"id_tecnico": data.Id_tecnico,
			"nome_tecnico":  data.Nome_tecnico,
			"login_tecnico": data.Login_tecnico,
			"email_lideres": email_lideres})
	}
}

func (repository *TecnicosRepo) GetTecnicos(c *gin.Context) {
	var tecnicos []models.Todos_Tecnicos
	reply, err := database.GetRedis("tecnicos")
	if err != nil {
		err := models.GetTecnicos(repository.Db, &tecnicos)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		m, err := json.Marshal(tecnicos)
		database.SetRedis("tecnicos", []byte(m))
		c.JSON(http.StatusOK, tecnicos)
	} else {
		var data interface{}
		if err := json.Unmarshal(reply, &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}
