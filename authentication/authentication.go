package authentication

import (
	"database/sql"
	"fmt"
)

func Login(loginName, pwd string) (string, error) {
	db, err := sql.Open("mysql", "root:RxDragDb@tcp(localhost:3306)/rxdrag")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var password string
	err = db.QueryRow("select password from rx_user where loginName = ?", loginName).Scan(&password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)

	return loginName, nil
}

func Logout() {

}
