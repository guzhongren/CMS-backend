package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Material struct {
	ID           string         `json:"id"`
	UserId       string         `json:"userId"`
	Name         string         `json:"name"`
	Location     sql.NullString `json:"location"`
	Type         int32          `json:"type"`
	Count        sql.NullInt64  `json:"count"`
	Provider     sql.NullString `json:"provider"`
	ProviderLink sql.NullString `json:"providerLink"`
	Images       sql.NullString `json:"images"`
	CreateTime   int64          `json:"createTime"`
	UpdateTime   sql.NullInt64  `json:"updateTime"`
	Price        sql.NullInt64  `json:"price"`
}

// 用户简略信息
type UserSimpleInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 返回数据
type MaterialResponse struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Location     sql.NullString `json:"location"`
	TypeName     string         `json:"typeName"`
	Count        sql.NullInt64  `json:"count"`
	Provider     sql.NullString `json:"provider"`
	ProviderLink sql.NullString `json:"providerLink"`
	Images       sql.NullString `json:"images"`
	CreateTime   int64          `json:"createTime"`
	UpdateTime   sql.NullInt64  `json:"updateTime"`
	Price        sql.NullInt64  `json:"price"`
	User         UserSimpleInfo `json:"user"`
}

func (material Material) GetAll(c echo.Context) error {
	materialList, err := material.selectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "获取所有物料错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  materialList,
		Message: "",
	})
}

// 添加物料
func (material Material) Add(c echo.Context) error {
	m := new(Material)
	utils := Utils{}
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误，请填写正确的参数！",
		})
	}

	userId := user.ID
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	m.ID = utils.GetGUID()
	m.CreateTime = time.Now().Unix()
	m.UserId = userId
	insertedMaterial, err := m.insert(*m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "服务器内部错误",
		})
	}
	response := MaterialResponse{}
	response.ID = insertedMaterial.ID
	response.Name = insertedMaterial.Name
	response.Location = insertedMaterial.Location
	// TODO: 处理物料类型
	// response.TypeName = insertedMaterial.Type
	response.Provider = insertedMaterial.Provider
	response.ProviderLink = insertedMaterial.ProviderLink
	response.Images = insertedMaterial.Images
	response.CreateTime = insertedMaterial.CreateTime
	response.User.ID = userId
	response.User.Name = user.Name
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  response,
		Message: "",
	})
}

// 插入
func (material Material) insert(m Material) (Material, error) {
	stmt, err := db.Prepare(`INSERT INTO "public"."b_material" ("id", "userId", "location", "type", "count", "provider", "providerLink", "images", "name", "createTime", "updateTime", "price") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING "id"`)
	if err != nil {
		log.Warn("插入物料前错误，", err)
		return Material{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(material.ID, material.UserId, material.Location, material.Type, material.Count, material.Provider, material.ProviderLink, material.Images, material.Name, material.CreateTime, material.UpdateTime, material.Price)
	if err != nil {
		log.Warn("插入物料时错误！", err)
		return Material{}, err
	}
	return m, nil
}

// 获取所有物料及信息
func (material Material) selectAll() ([]MaterialResponse, error) {
	rows, err := db.Query(`select m.id, m.name, m.location, m.materialType, m.count, m.provider, m."providerLink", m.images, m."createTime", m."updateTime", m.price, u.id as userId, u.name from (select m1.id,m1."userId", m1.name, m1.location, bmt.name as materialType, m1.count, m1.provider,m1."providerLink",m1.images, m1."createTime",m1."updateTime",m1.price from b_material m1 left join b_material_type bmt on m1.type = bmt.id) as m left join b_user u on m."userId"=u.id`)
	if err != nil {
		log.Warn("查询出错", err)
		return []MaterialResponse{}, err
	}
	var materialList = []MaterialResponse{}
	for rows.Next() {
		materialResponse := MaterialResponse{}
		err := rows.Scan(&materialResponse.ID, &materialResponse.Name, &materialResponse.Location, &materialResponse.TypeName, &materialResponse.Count, &materialResponse.Provider, &materialResponse.ProviderLink, &materialResponse.Images, &materialResponse.CreateTime, &materialResponse.UpdateTime, &materialResponse.Price, &materialResponse.User.ID, &materialResponse.User.Name)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return []MaterialResponse{}, err
		}
		materialList = append(materialList, materialResponse)
	}
	return materialList, nil
}

// func (material Material) selectById(id string) (MaterialResponse, error) {

// }
