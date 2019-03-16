package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// 物料
type Material struct {
	ID           string  `json:"id"`
	OwnerID      string  `json:"ownerId"`
	UpdateUserID string  `json:"updateUserId"`
	Name         string  `json:"name"`
	Location     string  `json:"location"`
	Type         int     `json:"type"`
	Count        int16   `json:"count"`
	Provider     string  `json:"provider"`
	ProviderLink string  `json:"providerLink"`
	Images       string  `json:"images"`
	CreateTime   int64   `json:"createTime"`
	UpdateTime   int64   `json:"updateTime"`
	Price        float64 `json:"price"`
}

// 物料类型
type MaterialType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 用户简略信息
type UserSimpleInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 返回数据
type MaterialResponse struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Location     sql.NullString  `json:"location"`
	TypeName     string          `json:"typeName"`
	Count        sql.NullInt64   `json:"count"`
	Provider     sql.NullString  `json:"provider"`
	ProviderLink sql.NullString  `json:"providerLink"`
	Images       sql.NullString  `json:"images"`
	CreateTime   int64           `json:"createTime"`
	UpdateTime   sql.NullInt64   `json:"updateTime"`
	Price        sql.NullFloat64 `json:"price"`
	Owner        UserSimpleInfo  `json:"owner"`
	UpdateUser   UserSimpleInfo  `json:"updateUser"`
}

// 获取所有的物料
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
		log.Warn("添加：获取用户数据错！", err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "添加：获取用户数据错！",
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
	m.OwnerID = userId
	m.UpdateUserID = userId
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
	response.Location = sql.NullString{
		String: m.Location,
		Valid:  m.Location == "",
	}
	mType, err := material.getMaterialTypeById(m.Type)
	if err != nil {
		log.Warn("获取物料类型错误", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "服务器内部错误",
		})
	}
	response.TypeName = mType.Name
	response.Provider = sql.NullString{
		String: insertedMaterial.Provider,
		Valid:  insertedMaterial.Provider == "",
	}
	response.ProviderLink = sql.NullString{
		String: insertedMaterial.ProviderLink,
		Valid:  insertedMaterial.ProviderLink == "",
	}
	response.Images = sql.NullString{
		String: insertedMaterial.Images,
		Valid:  insertedMaterial.Images == "",
	}
	response.CreateTime = insertedMaterial.CreateTime
	response.Owner.ID = userId
	response.Owner.Name = user.Name
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  response,
		Message: "",
	})
}

// 删除物料路由
func (material Material) Delete(c echo.Context) error {
	id := c.Param("id")
	m, err := material.getOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "删除失败，没有该物料!",
		})
	}
	isDeleted := material.delete(m.ID)
	if !isDeleted {
		log.Warn("没有该物料错误", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "通过id删除物料错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  isDeleted,
		Message: "",
	})
}

// 更新用户信息
func (material Material) Update(c echo.Context) error {
	id := c.Param("id")
	m := new(Material)
	if err := c.Bind(m); err != nil {
		log.Warn("更新用户参数错误", err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	dbMaterial, err := material.getOne(id)
	if err != nil {
		log.Warn("更新物料: 获取要更新的物料信息错误", err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "没有该物料，请确认该物料存在！",
		})
	}
	utils := Utils{}
	user, err := utils.GetUserFromContext(c)
	if err != nil {
		log.Warn("更新：获取物料数据错！", err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "更新：获取物料数据错！",
		})
	}
	m.ID = id
	m.OwnerID = dbMaterial.Owner.ID
	m.UpdateUserID = user.ID
	m.UpdateTime = time.Now().Unix()
	_, err = material.update(*m)
	if err != nil {
		log.Warn("更新物料数据错误！", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "更新物料数据错误！",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  id,
		Message: "",
	})
}

func (material Material) GetOne(c echo.Context) error {
	id := c.Param("id")
	m, err := material.getOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "查询失败，没有该物料!",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  m,
		Message: "",
	})
}

// 获取物料类型路由
func (Material Material) GetMaterialTypeById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Warn("参数问题")
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	materialTypeList, err := Material.getMaterialTypeById(id)
	if err != nil {
		log.Warn("查询物料类型出错", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "服务器内部错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  materialTypeList,
		Message: "",
	})
}

// 获取物料类型路由
func (material Material) GetMaterialType(c echo.Context) error {
	materialTypeList, err := material.getMaterialTypes()
	if err != nil {
		log.Warn("查询物料类型出错了")
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "服务器内部错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  materialTypeList,
		Message: "",
	})
}

// 插入
func (material Material) insert(m Material) (Material, error) {
	stmt, err := db.Prepare(`INSERT INTO "public"."b_material" ("id", "ownerId", "location", "type", "count", "provider", "providerLink", "images", "name", "createTime", "updateTime", "price") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING "id"`)
	if err != nil {
		log.Warn("插入物料前错误，", err)
		return Material{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(material.ID, material.OwnerID, material.Location, material.Type, material.Count, material.Provider, material.ProviderLink, material.Images, material.Name, material.CreateTime, material.UpdateTime, material.Price)
	if err != nil {
		log.Warn("插入物料时错误！", err)
		return Material{}, err
	}
	return m, nil
}

// 删除
func (material Material) delete(id string) bool {
	stmt, err := db.Prepare("DELETE FROM b_material WHERE id=$1")
	if err != nil {
		log.Warn("删除物料：操作数据库错误", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Warn("执行删除物料错误", err)
		return false
	}
	return true
}

// 更新物料信息
// TODO: 处理文件存储
func (material Material) update(m Material) (string, error) {
	m.UpdateTime = time.Now().Unix()
	stmt, err := db.Prepare(`UPDATE public.b_material SET "ownerId"=$2, location=$3, type=$4, count=$5, provider=$6, "providerLink"=$7, images=$8, name=$9, "updateTime"=$10, price=$11, "updateUserId"=$12 WHERE id=$1`)
	if err != nil {
		log.Warn("更新物料：操作数据库错误", err)
		return "", err
	}
	defer stmt.Close()
	if err != nil {
		log.Warn("update: 获取当前用户出错", err)
		return "", err
	}
	_, err = stmt.Exec(m.ID, m.OwnerID, m.Location, m.Type, m.Count, m.Provider, m.ProviderLink, m.Images, m.Name, m.UpdateTime, m.Price, m.UpdateUserID)
	if err != nil {
		log.Warn("执行更新用户错误", err)
		return "", err
	}
	return m.ID, nil
}

// 查询单个物料
func (material Material) getOne(id string) (MaterialResponse, error) {
	m := MaterialResponse{}
	err := db.QueryRow(`select m3.id, m3.name, m3.location, m3.typeName, m3.count, m3.provider, m3."providerLink", m3.images, m3."createTime", m3."updateTime", m3.price, m3."ownerId", m3.ownerName, m3."updateUserId", m3.updateUserName from (select m2.*, b2.name as updateUserName from (select m1.*, bu.name as ownerName from (select a.*, b.name as typeName from b_material a  left join b_material_type b on a.type = b.id where a.type=b.id) as m1 left join b_user bu on bu.id=m1."ownerId") as m2 left join b_user b2 on m2."updateUserId"=b2.id) as m3 where m3.id=$1`, id).Scan(&m.ID, &m.Name, &m.Location, &m.TypeName, &m.Count, &m.Provider, &m.ProviderLink, &m.Images, &m.CreateTime, &m.UpdateTime, &m.Price, &m.Owner.ID, &m.Owner.Name, &m.UpdateUser.ID, &m.UpdateUser.Name)
	if err != nil {
		log.Warn("查询物料出错", err)
		return MaterialResponse{}, err
	}
	return m, nil
}

// 获取所有物料及信息
func (material Material) selectAll() ([]MaterialResponse, error) {
	rows, err := db.Query(`select m.id, m.name, m.location, m.materialType, m.count, m.provider, m."providerLink", m.images, m."createTime", m."updateTime", cast(m.price as float), u.id as ownerId, u.name from (select m1.id,m1."ownerId", m1.name, m1.location, bmt.name as materialType, m1.count, m1.provider,m1."providerLink",m1.images, m1."createTime",m1."updateTime", m1.price from b_material m1 left join b_material_type bmt on m1.type = bmt.id) as m left join b_user u on m."ownerId"=u.id`)
	if err != nil {
		log.Warn("查询所有物料出错", err)
		return []MaterialResponse{}, err
	}
	var materialList = []MaterialResponse{}
	for rows.Next() {
		materialResponse := MaterialResponse{}
		err := rows.Scan(&materialResponse.ID, &materialResponse.Name, &materialResponse.Location, &materialResponse.TypeName, &materialResponse.Count, &materialResponse.Provider, &materialResponse.ProviderLink, &materialResponse.Images, &materialResponse.CreateTime, &materialResponse.UpdateTime, &materialResponse.Price, &materialResponse.Owner.ID, &materialResponse.Owner.Name)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return []MaterialResponse{}, err
		}
		materialList = append(materialList, materialResponse)
	}
	return materialList, nil
}

// 根据id获取特定的物料类型
func (material Material) getMaterialTypeById(id int) (MaterialType, error) {
	var ID string
	var name string
	err := db.QueryRow(`select id,name from b_material_type where id=$1`, id).Scan(&ID, &name)
	if err != nil {
		if err != nil {
			log.Warn("查询出错", err)
			return MaterialType{}, err
		}
	}

	return MaterialType{ID, name}, nil
}

// 获取所有的物料类型
func (material Material) getMaterialTypes() ([]MaterialType, error) {
	rows, err := db.Query(`select id,name from b_material_type`)
	if err != nil {
		if err != nil {
			log.Warn("查询所有类型出错", err)
			return []MaterialType{}, err
		}
	}
	var materialTypeList = []MaterialType{}
	for rows.Next() {
		t := MaterialType{}
		err := rows.Scan(&t.ID, &t.Name)
		if err != nil {
			log.Warn("处理查询物料结果出错", err)
			return []MaterialType{}, err
		}
		materialTypeList = append(materialTypeList, t)
	}
	return materialTypeList, nil
}
