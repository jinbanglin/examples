package main

import (
	"bytes"
	"fmt"
	"mime"
	"mime/multipart"
	"strings"

	proto "github.com/jinbanglin/examples/form/api/proto"
	api "github.com/jinbanglin/go-api/proto"
	"github.com/jinbanglin/go-log"
	"github.com/jinbanglin/go-micro"

	"context"
)

type Form struct{}

func (f *Form) Submit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.Body = fmt.Sprintf("got your values %+v", req.Post)
	fmt.Println("---1--",req.Path)
	fmt.Println(req.Get)
	fmt.Println("===2==",req.String())
	fmt.Println("===3==",req.GetBody())
	fmt.Println("===4==",req.Body)
	return nil
}

func (f *Form) Multipart(ctx context.Context, req *api.Request, rsp *api.Response) error {
	ct := strings.Join(req.Header["Content-Type"].Values, ",")
	mt, p, err := mime.ParseMediaType(ct)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(mt, "multipart/") {
		return fmt.Errorf("%v does not contain multipart", mt)
	}
	r := multipart.NewReader(bytes.NewReader([]byte(req.Body)), p["boundary"])
	form, err := r.ReadForm(32 << 20)
	if err != nil {
		return err
	}
	rsp.Body = fmt.Sprintf("got your values %+v", form)
	fmt.Println("-----",req.Url)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.form"),
	)

	service.Init()

	proto.RegisterFormHandler(service.Server(), new(Form))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
