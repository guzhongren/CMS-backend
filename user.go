package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	RoleId   string `json:"roleId"`
	Password string `json:"password"`
}

// 返回信息
type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// 更新用户
func (user User) UpdateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		log.Warn(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	innerUser, err := user.getOne(u.ID)
	if err != nil {
		insertUser, err := user.insert(innerUser)
		if err != nil {
			return c.JSON(http.StatusOK, &Response{
				Success: true,
				Result:  insertUser,
				Message: "",
			})
		}
		return c.JSON(http.StatusOK, &Response{
			Success: false,
			Result:  "",
			Message: "没有该用户，更新失败",
		})
	}
	user, e := user.update(*u)
	if e != nil {
		log.Warn("更新用户信息错误", e)
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "通过id更新用户错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  user,
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
	user, err := user.getOne(id)
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
		Result:  user,
		Message: "",
	})
}

// 新增用户
func (user User) AddUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		log.Warn(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	var utils = Utils{}
	u.Password = utils.CryptoStr(u.Password)
	u.ID = utils.GetGUID()
	user, err := user.insert(*u)
	if err != nil {
		log.Warn("插入数据库错误")
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "插入数据库错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  user,
		Message: "",
	})
}

// 插入用户
func (u User) insert(user User) (User, error) {
	stmt, err := db.Prepare("insert into b_user(id,name,roleId,password) values($1,$2,$3,$4)")
	if err != nil {
		log.Warn("插入用户数据前错误", err)
		return User{}, err
	}
	defer stmt.Close()
	log.Info(user.ID)
	_, err = stmt.Exec(user.ID, user.Name, user.RoleId, user.Password)
	if err != nil {
		log.Warn("插入用户错误", err)
		return User{}, err
	}
	return user, nil
}

// 删除用户
func (u User) delete(id string) (string, error) {
	stmt, err := db.Prepare("DELETE FROM b_user WHERE id=$1")
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

// 更新用户
func (u User) update(user User) (User, error) {
	stmt, err := db.Prepare("UPDATE b_user set name=$2,roleId=$3,password=$4 WHERE id=$1")
	if err != nil {
		log.Warn("更新用户：操作数据库错误", err)
		return user, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.RoleId, user.Password)
	if err != nil {
		log.Warn("执行更新用户错误", err)
		return user, err
	}
	return user, nil
}

// 通过用户名查询用户
func (user User) getOne(id string) (User, error) {
	err := db.QueryRow("select id, name, roleId from b_user where id=$1", id).Scan(&user.ID, &user.Name, &user.RoleId)
	if err != nil {
		log.Warn("查询用户出错", err)
		return User{}, err
	}
	return user, nil
}

// 通过用户名查询用户
func (user User) GetUserByName(userName string) (User, error) {
	err := db.QueryRow("select id, name, roleId from b_user where name=$1", userName).Scan(&user.ID, &user.Name, &user.RoleId)
	if err != nil {
		log.Warn("查询用户出错", err)
		return User{}, err
	}
	return user, nil
}
func (user User) GetAll() ([]User, error) {
	rows, err := db.Query("select id, roleId ,name from b_user")
	if err != nil {
		log.Warn("查询出错", err)
		return []User{}, err
	}
	var userList = []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.RoleId, &user.Name)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return []User{}, err
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
