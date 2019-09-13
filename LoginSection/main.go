package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

//Khoi tao bien session la bien toan cuc
//Cau hinh tuy chon cho session
var session = sessions.New(sessions.Config{
	Cookie:       "test", //Ten cookie
	AllowReclaim: true,
	Expires:      time.Hour * 1, //Thoi gian session
})

// User struct ...
type User struct {
	ID       int
	Name     string
	Password string
}

//Tao DS user la mang cac Struct User
var users = []User{
	{
		ID:       0,
		Name:     "thuhang",
		Password: "3110",
	},
	{
		ID:       1,
		Name:     "buihien",
		Password: "0109",
	},
	{
		ID:       2,
		Name:     "phamphuong",
		Password: "1608",
	},
	{
		ID:       3,
		Name:     "bichngoc",
		Password: "0709",
	},
}

func main() {
	app := iris.New()
	// Đăng kí các file HTML trong thư mục view
	tmpl := iris.HTML("./views", ".html")
	// Mỗi lần chỉnh sửa file HTML thì chỉ cần refresh lại trang
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.HandleDir("/resources", "./resources")

	// Method GET
	app.Get("/", HomeView)
	app.Get("/register", RegisterView)
	app.Get("/sign-in", SigninView)
	app.Get("/blog", BlogView)
	app.Get("/logout", LogoutView)
	app.Get("/detele/{id}", DeteleHandler)
	app.Get("/update/{id}", UpdateHandler)

	//Method POST
	app.Post("/register", RegisterHandler)
	app.Post("/sign-in", SigninHandler)
	app.Post("/update", UpdateUserHandler)

	//Config run pot for App
	app.Run(iris.Addr(":8080")) // defaults to that but you can change it.

}

// --------------Xu ly method GET-----------------
func HomeView(ctx iris.Context) {
	ctx.ViewData("users", users)
	var mess string = ""
	if len(users) == 0 {
		mess = `Danh sach trong!!`
	} else {
		mess = "Co " + strconv.Itoa(len(users)) + " user trong danh sach"
	}
	ctx.ViewData("mess", mess)
	ctx.View("home.html")
}

func RegisterView(ctx iris.Context) {
	ctx.View("register.html")
}

func SigninView(ctx iris.Context) {
	sess := session.Start(ctx)
	tmp := sess.Get("login")
	if tmp != nil {
		ctx.Redirect("blog")
		return
	}
	ctx.View("sign-in.html")
}
func BlogView(ctx iris.Context) {
	var name string
	sess := session.Start(ctx)
	name = sess.Get("login").(string)

	ctx.ViewData("name", name)
	ctx.View("blog.html")
}
func LogoutView(ctx iris.Context) {
	sess := session.Start(ctx)
	sess.Destroy()
	ctx.Redirect("/sign-in")
}

func DeteleHandler(ctx iris.Context) {
	ID := ctx.Params().Get("id")
	// fmt.Println("ID :", ID)
	userID, _ := strconv.Atoi(ID)
	// fmt.Println("userID :",userID)
	var userDelete User
	fmt.Println("userDelete :", userDelete)

	for i, user := range users {
		if user.ID == userID {
			// fmt.Println("i :", i)
			users = append(users[:i], users[i+1:]...)
			// fmt.Println(users)
			fmt.Println("Len: ", len(users))

			break
		}
	}
	ctx.ViewData("users", users)
	ctx.Redirect("/")

}

func UpdateHandler(ctx iris.Context) {
	ID := ctx.Params().Get("id")
	userID, _ := strconv.Atoi(ID)

	var userUpdate User

	for _, user := range users {
		if user.ID == userID {
			userUpdate = user
			break
		}
	}
	ctx.ViewData("userUpdate", userUpdate)
	ctx.View("update.html")
}

// --------------Xu ly method POST-----------------
func RegisterHandler(ctx iris.Context) {
	user := User{}
	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	user.ID = len(users)
	users = append(users, user)

	// ctx.Writef("User: %#v", user)
	ctx.Redirect("/")
}

func SigninHandler(ctx iris.Context) {
	user := User{}
	err := ctx.ReadForm(&user)
	if err != nil && !iris.IsErrPath(err) /* see: https://github.com/kataras/iris/issues/1157 */ {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}

	isRight := false
	var userID int
	fmt.Println(userID)
	for _, v := range users {
		if user.Name == v.Name && user.Password == v.Password {
			isRight = true
			userID = v.ID
			break
		}
	}
	if isRight {

		sess := session.Start(ctx)
		sess.Set("login", user.Name)
		ctx.Redirect("/blog")
	} else {
		ctx.Writef("Failed to sign in")
	}
}
func UpdateUserHandler(ctx iris.Context) {
	user := User{}
	err := ctx.ReadForm(&user)
	fmt.Println("user: ", user)
	fmt.Println("-------------------")
	if err != nil && !iris.IsErrPath(err) {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	for i, v := range users {
		if i == user.ID {
			user = v
			v.Name = user.Name
			v.Password = user.Password
			fmt.Println("ID: ", i)
			fmt.Println("name: ", user.Name)
			fmt.Println("pass: ", user.Password)
			fmt.Println("vname: ", v.Name)
			fmt.Println("vpass: ", v.Password)
			break
		}
	}
	fmt.Println("user: ", user)
	fmt.Println("users: ", users)
	ctx.ViewData("users", users)
	ctx.Redirect("/")
}
