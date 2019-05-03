package 第0阶段

import "fmt"

type employees struct {
	id     string
	name   string
	age    int
	salary int
}

func (e employees) print() {
	fmt.Println("工号：%s", e.id)
	fmt.Println("姓名：%s", e.name)
	fmt.Println("年龄：%s", e.age)
	fmt.Println("薪资：%s", e.salary)
}
func main() {
	var e0 employees
	var e1 employees
	var e2 employees
	fmt.Println("请输入第1位员工的姓名 年龄 薪资")
	fmt.Scanf("%s %s %d %d", &e0.id, &e0.name, &e0.age, &e0.salary)
	fmt.Println("请输入第2位员工的姓名 年龄 薪资")
	fmt.Scanf("%s %s %d %d", &e1.id, &e1.name, &e1.age, &e1.salary)
	fmt.Println("请输入第3位员工的姓名 年龄 薪资")
	fmt.Scanf("%s %s %d %d", &e2.id, &e2.name, &e2.age, &e2.salary)
	arr := []employees{e0, e1, e2}
	for index := 0; index < len(arr); index++ {
		if arr[index].salary > 8000 {
			arr[index].print()
		}
	}
}
