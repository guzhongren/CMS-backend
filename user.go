package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	RoleId     string `json:"roleId"`
	Password   string `json:"password"`
	CreateTime int64  `json:"createTime"`
	LoginTime  int64  `json:"loginTime"`
}

// 返回信息
type UserResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	RoleId     string `json:"roleId"`
	CreateTime int64  `json:"createTime"`
	LoginTime  int64  `json:"loginTime"`
}

// 更新用户
func (user User) UpdateUser(c echo.Context) error {
	u := new(User)
	u.ID = c.Param("id")
	if err := c.Bind(u); err != nil {
		log.Warn(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	log.Warn("没有该用户？", u)
	_, err := user.getOne(u.ID)
	if err != nil {
		return user.AddUser(c)
	}
	_, e := user.update(*u)
	if e != nil {
		log.Warn("更新用户信息错误", e)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "通过id更新用户错误",
		})
	}
	updatedUser, err := user.getOne(u.ID)
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  updatedUser,
		Message: "",
	})
}

func (user User) ResetPassword(c echo.Context) error {
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	_, err := user.getOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "没有该用户，请使用合理的用户！",
		})
	}
	err = user.updatePassword(id, u.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "重置密码错误，请联系管理员！",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  true,
		Message: "",
	})
}

// 删除用户
func (user User) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	innerUser, err := user.getOne(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "删除失败，没有该用户!",
		})
	}
	deletedId, e := user.delete(innerUser.ID)
	if e != nil {
		log.Warn("没有该用户错误", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "通过id删除用户错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  deletedId,
		Message: "",
	})
}

// 获取用户
func (user User) GetUser(c echo.Context) error {
	id := c.Param("id")
	u, err := user.getOne(id)
	log.Info("test", u)
	if err != nil {
		log.Warn("获取用户错误", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "通过id获取用户错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  u,
		Message: "",
	})
}

// 新增用户
func (user User) AddUser(c echo.Context) error {
	utils := Utils{}
	userResponse, err := utils.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	if userResponse.Role != "publish" {
		return c.JSON(http.StatusUpgradeRequired, &Response{
			Success: false,
			Result:  "",
			Message: "该角色没有添加用户的权限！",
		})
	}
	u := new(User)
	if err := c.Bind(u); err != nil {
		log.Warn("绑定数据错误", err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	innerUser, err := user.GetUserByName(u.Name)
	if err == nil {
		log.Info("已存在该用户，请使用新的用户名")
		_ = user.updatePassword(innerUser.ID, utils.CryptoStr(u.Password))
		_, _ = user.activeUser(innerUser.ID)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: true,
			Result:  innerUser.ID,
			Message: "已存在该用户，请使用该用户名和密码登录！",
		})
	}
	u.Password = utils.CryptoStr(u.Password)
	u.ID = utils.GetGUID()
	u.CreateTime = time.Now().Unix()
	insertedUser, err := user.insert(*u)
	if err != nil {
		log.Warn("插入数据库错误")
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "插入数据库错误",
		})
	}
	resUser, err := user.getOne(insertedUser.ID)
	if err != nil {
		log.Warn("插入又获取数据错误", err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "插入数据库错误",
		})
	}
	log.Info(resUser)
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  resUser,
		Message: "",
	})
}

// 插入用户
func (u User) insert(user User) (User, error) {
	stmt, err := db.Prepare(`insert into b_user(id,name,"roleId",password,"createTime") values($1,$2,$3,$4,$5)`)
	if err != nil {
		log.Warn("插入用户数据前错误", err)
		return User{}, err
	}
	defer stmt.Close()
	user.CreateTime = time.Now().Unix()
	_, err = stmt.Exec(user.ID, user.Name, user.RoleId, user.Password, user.CreateTime)
	if err != nil {
		log.Warn("插入用户错误", err)
		return User{}, err
	}
	return user, nil
}

// 删除用户
func (u User) delete(id string) (string, error) {
	stmt, err := db.Prepare("UPDATE b_user SET isdeleted=true WHERE id=$1")
	if err != nil {
		log.Warn("删除用户：操作数据库错误", err)
		return "", err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		log.Warn("执行删除用户错误", err)
		return id, err
	}
	log.Info("删除用户后", result)
	if count, _ := result.RowsAffected(); count <= 0 {
		return "", errors.New("没有该行数据")
	}
	return id, nil
}
func (u User) activeUser(id string) (string, error) {
	stmt, err := db.Prepare("UPDATE b_user SET isdeleted=false WHERE id=$1")
	if err != nil {
		log.Warn("激活用户：操作数据库错误", err)
		return "", err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		log.Warn("执行激活用户错误", err)
		return id, err
	}
	log.Info("激活用户后", result)
	if count, _ := result.RowsAffected(); count <= 0 {
		return "", errors.New("没有该行数据")
	}
	return id, nil
}

// 更新用户
func (u User) update(user User) (User, error) {
	stmt, err := db.Prepare(`UPDATE b_user set name=$2,"roleId"=$3 WHERE id=$1`)
	if err != nil {
		log.Warn("更新用户：操作数据库错误", err)
		return user, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.RoleId)
	if err != nil {
		log.Warn("执行更新用户错误", err)
		return user, err
	}
	return user, nil
}

// 重置密码
func (u User) updatePassword(id, newPW string) error {
	stmt, err := db.Prepare(`UPDATE b_user set password=$2 WHERE id=$1`)
	if err != nil {
		log.Warn("重置密码：操作数据库错误", err)
		return err
	}
	defer stmt.Close()
	utils := Utils{}
	password := utils.CryptoStr(newPW)
	_, err = stmt.Exec(id, password)
	if err != nil {
		log.Warn("执行重置密码错误", err)
		return err
	}
	return nil
}

// 通过用户id查询用户
func (user User) getOne(id string) (UserResponse, error) {
	u := UserResponse{}
	err := db.QueryRow(`select b.id, b.name, b."createTime", br.name, br.id from b_user b left join b_role br on b."roleId"= br.id where b.id=$1`, id).Scan(&u.ID, &u.Name, &u.CreateTime, &u.Role, &u.RoleId)
	if err != nil {
		log.Warn("查询用户出错", err)
		return UserResponse{}, err
	}
	return u, nil
}

// 通过用户名查询用户
func (user User) GetUserByName(userName string) (UserResponse, error) {
	u := UserResponse{}
	err := db.QueryRow(`select b.id, b.name, br.name from b_user b left join b_role br on b."roleId"=br.id where b.name=$1`, userName).Scan(&u.ID, &u.Name, &u.Role)
	if err != nil {
		log.Warn("根据用户名查询用户出错", err)
		return UserResponse{}, err
	}
	return u, nil
}

// 获取所有用户sql处理
func (user User) GetAll() ([]UserResponse, error) {
	rows, err := db.Query(`select b.id, b.name, b."createTime", br.name, br.id from b_user b left join b_role br on b."roleId" = br.id where b.isdeleted=false`)
	if err != nil {
		log.Warn("查询出错", err)
		return []UserResponse{}, err
	}
	var userList = []UserResponse{}
	for rows.Next() {
		user := UserResponse{}
		err := rows.Scan(&user.ID, &user.Name, &user.CreateTime, &user.Role, &user.RoleId)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return []UserResponse{}, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}

// 获取所有的用户
func (user User) GetAllUsers(c echo.Context) error {
	userList, err := user.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "所有用户获取错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  userList,
		Message: "",
	})
}
