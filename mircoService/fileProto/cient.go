package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"io"
	"net/http"
	proto "netgo/mircoService/fileProto/proto"
)

var c1 client.Client
var fileService1 proto.FileService

func UploadFile1(rsp http.ResponseWriter, req *http.Request) {
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	// 取到文件对象
	files, ok := req.MultipartForm.File["file"]
	if !ok {
		rsp.WriteHeader(400)
		_, _ = rsp.Write([]byte("请选择上传文件"))
		return
	}
	// 将文件通过流式传输到srv
	file, err := files[0].Open()
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	//建立连接
	//这里用的临时文件储存的方式，如果因为
	//
	next, _ := c1.Options().Selector.Select("file.service")
	node, _ := next()
	stream, err := fileService1.File(context.Background(), func(options *client.CallOptions) {
		// 指定节点
		options.Address = []string{node.Address}
	})
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	for {
		buff := make([]byte, 1024*1024)
		sendLen, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {

				err = stream.Send(&proto.FileSlice{
					Byte: nil,
					Len:  -1,
				})
				if err != nil {
					rsp.WriteHeader(500)
					_, _ = rsp.Write([]byte(err.Error()))
					return
				}
				break
			}
			rsp.WriteHeader(500)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
		err = stream.Send(&proto.FileSlice{
			Byte: buff[:sendLen],
			Len:  int64(sendLen),
		})
		if err != nil {
			rsp.WriteHeader(500)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
	}

	fileMsg := &proto.FileSliceMsg{}
	if err := stream.RecvMsg(fileMsg); err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	_ = stream.Close()
	println(fileMsg.FileName)
}
func main() {
	service := micro.NewService(micro.Name("file.client"))
	service.Init()
	c1 = service.Client()
	fileService1 = proto.NewFileService("file.service", c1)
	http.HandleFunc("/upload", UploadFile1)
	_ = http.ListenAndServe(":8085", nil)
}
