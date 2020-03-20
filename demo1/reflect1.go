package main

import (
	"fmt"
	"reflect"
)

//反射

func reflectExample(a interface{}){
	//判断变量的类型信息
	t:=reflect.TypeOf(a)
	fmt.Printf("t type of a is :%v\n", t)

	k := t.Kind()
	switch k {
	case reflect.Int64, reflect.Int:
		fmt.Println("a is int")
	case reflect.Float32, reflect.Float64:
		fmt.Println("a is float")
	}
}

func reflectValue(a interface{}){
	//判断变量值的信息
	v:=reflect.ValueOf(a)
	fmt.Printf("v 的值是 :%v\n", v)
	t := v.Type()
	fmt.Printf("v 的类型是:%v\n", t)
	fmt.Println()
	k := v.Kind()
	switch k {
	case reflect.Int64, reflect.Int:
		fmt.Printf("k a 的类型是int,a 的值是 :%d\n",v.Int())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("k a 的类型是 float a 的值是 :%f\n",v.Float())
	}
}


func reflectSetValve(a interface{}){
	//把传进来的参数重新设置值
	v:=reflect.ValueOf(a)
	fmt.Println(v)
	t := v.Type()
	fmt.Println(t)
	k := v.Kind()
	switch k {
	case reflect.Int, reflect.Int64, reflect.Int32:
		fmt.Printf("Int类型，把值设置为100\n")
		v.SetInt(100)
	case reflect.Float32, reflect.Float64:
		fmt.Printf("float类型，把值设置为6.4\n")
		//v.SetFloat(6.4)
		v.SetFloat(6.4)
	case reflect.Ptr:
		fmt.Printf("指针类型，把值设置为6\n")
		v.Elem().SetFloat(6)
	}
}


type Student struct {
	Name string `json:"name"`
	Sex int	`json:"sex"`
	Age int `json:"age"`
	Score float32 `json:"score"`
}

func (s Student) Setname(name string){
	s.Name = name
	fmt.Printf("s.Name = %s\n", s.Name)
}

func (s Student)Print(){

	fmt.Println(s.Name, s.Age, s.Score)
}

func StrctToReflect(){
	var s Student
	v:= reflect.ValueOf(s)
	fmt.Println(v)
	t := v.Type()
	fmt.Println(t)

	k := v.Kind()
	switch k {
	case reflect.Int32, reflect.Int64, reflect.Int:
		fmt.Println("Int类型")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float类型")
	case reflect.Ptr:
		fmt.Println("指针类型")
	case reflect.Struct:
		fmt.Println("结构体Struct")
	default:
		fmt.Println("不知道的类型")
	}

	num := v.NumField()
	fmt.Println("结构体中的字段是:",num)
	for i :=0; i<num; i++{

		fmt.Println(t.Field(i).Name)
		fmt.Println(v.Field(i))

	}

}


func reflectChangStruct(){
	//通过反射设置结构体内的字段的值
	var s Student
	v:=reflect.ValueOf(&s)
	v.Elem().FieldByName("Name").SetString("Hans")
	v.Elem().FieldByName("Sex").SetInt(1)
	v.Elem().FieldByName("Age").SetInt(32)
	v.Elem().FieldByName("Score").SetFloat(60.5)
	fmt.Println(s)
	fmt.Printf("S:%#v\n", s)
}

func reflectGetStructMethod()  {
	//用反射得到结构体的方法
	//得到方法属于类型，用t
	var s Student
	v := reflect.ValueOf(s)
	t := v.Type()
	fmt.Println(t.NumMethod(),"个方法.")

	for i := 0; i<t.NumMethod(); i++{
		method := t.Method(i)
		fmt.Printf("Struct %d method, name:%s type:%v\n",t.NumMethod(), method.Name, method.Type)
	}
}

func useRefectCallStructMethod()  {
	//用反射调用结构体中的方法。
	//调用方法属于值方面用v
	var s Student
	v := reflect.ValueOf(s)

	fmt.Println("利用反射调用方法,(有参数):")
	m1 := v.MethodByName("Setname")
	fmt.Println(m1)
	var arg1 []reflect.Value
	name := "Stu01"
	nameVal := reflect.ValueOf(name)
	arg1 = append(arg1, nameVal)
	m1.Call(arg1)

	fmt.Println("利用反射调用方法,(无参数):")
	m2 := v.MethodByName("Print")
	fmt.Println(m2)
	var arg2 []reflect.Value
	m2.Call(arg2)
}

func reflectGetTagInfo()  {
	//利用反射获得结构体内的tag
	var s Student
	//传入的是指针，获得tag的方法
	fmt.Println("利用指针获得tag方法")
	v := reflect.ValueOf(&s)
	t:= v.Type()

	for i :=0; i<t.Elem().NumField();i++{
		field := t.Elem().Field(i)
		fmt.Printf("tag json =%s\n", field.Tag.Get("json"))
	}

	//传入的是一个值，获得tag的方法
	fmt.Println("利用值拷贝获得tag方法")
	v2:= reflect.ValueOf(s)
	t1 := v2.Type()

	for i :=0; i<t1.NumField();i++{
		field := t1.Field(i)
		fmt.Printf("tag json =%s\n", field.Tag.Get("json"))
	}


	/*
	两者的是区别：
		利用指针获得: 取字段个数和具体字段是要用t.Elem().NumField()
		利用值拷贝获得则不用这Elem()直接给t.NumField()
	*/
}


func main() {
	var x float64 = 3.1

	fmt.Println("type:")
	reflectExample(x)
	fmt.Println("value:")
	reflectValue(x)
	fmt.Println("Set value:")
	reflectSetValve(&x)
	fmt.Println(x)
	fmt.Println("StructToReflect:")
	StrctToReflect()
	fmt.Println()
	fmt.Println("reflectChangStruct")
	reflectChangStruct()
	fmt.Println()
	fmt.Println("reflectGetStructMethod")
	reflectGetStructMethod()

	fmt.Println()
	fmt.Println("useRefectCallStructMethod")
	useRefectCallStructMethod()

	fmt.Println()
	fmt.Println("reflectGetTagInfo")
	reflectGetTagInfo()
}