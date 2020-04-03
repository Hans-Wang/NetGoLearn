package main
import(
	"context"
	"fmt"
	//     "io"
	"io/ioutil"
	//     "os"
	proto "netgo/mircoService/fileProto/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/errors"
)

type File struct{}

func (g *File) File(ctx context.Context, file proto.File_FileStream) error {
	//将接收到的文件存储在临时文件夹中
	// File_FileStream类型有Recv函数，源源不断从客户端收到slice的消息
	temp,err:=ioutil.TempFile("","micro")
	if err!=nil{
		return errors.InternalServerError("file.service", err.Error())
	}
	for {
		b,err:=file.Recv()//b被自动unmarshall成FileSlice类型
		if err!=nil{
			return errors.InternalServerError("file.service", err.Error())
		}
		if b.Len==-1{//预先定义好如果Len==-1，就说明到头了,EOF包
			break
		}
		if _,err:=temp.Write(b.Byte);err!=nil{//流式的Appen到文件中
			return errors.InternalServerError("file.service", err.Error())
		}
	}
	println(temp.Name())//生成的临时文件名打印
	return file.SendMsg(&proto.FileSliceMsg{
		FileName:temp.Name(),
	})
}

func main(){
	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	service := micro.NewService(
		micro.Name("file.service"),
		micro.Version("latest"),
	)
	service.Init()

	// 注册服务
	_ = proto.RegisterFileHandler(service.Server(), new(File))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}