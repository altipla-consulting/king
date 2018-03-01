package main

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type templateData struct {
	proto    *descriptor.FileDescriptorProto
	typesMap map[string]string

	Version        string
	Package        string
	SourceFilename string
	Imports        []string
}

func (p *templateData) Services() []*templateService {
	services := make([]*templateService, len(p.proto.GetService()))

	for i, svc := range p.proto.GetService() {
		services[i] = &templateService{
			proto:    svc,
			pkg:      p.proto.GetPackage(),
			typesMap: p.typesMap,
		}
	}

	return services
}

func (p *templateData) Quote() string {
	return "`"
}

type templateService struct {
	proto    *descriptor.ServiceDescriptorProto
	pkg      string
	typesMap map[string]string
}

func (svc *templateService) ClientName() string {
	return svc.proto.GetName() + "Client"
}

func (svc *templateService) ServerName() string {
	return svc.proto.GetName() + "Server"
}

func (svc *templateService) Name() string {
	return svc.proto.GetName()
}

func (svc *templateService) ClientImplName() string {
	return "clientImpl" + svc.proto.GetName()
}

func (svc *templateService) ServiceName() string {
	if svc.pkg == "" {
		return svc.proto.GetName()
	}

	return svc.pkg + "." + svc.proto.GetName()
}

func (svc *templateService) Methods() []*templateMethod {
	methods := make([]*templateMethod, len(svc.proto.GetMethod()))

	for i, proto := range svc.proto.GetMethod() {
		methods[i] = &templateMethod{
			proto:    proto,
			typesMap: svc.typesMap,
		}
	}

	return methods
}

type templateMethod struct {
	proto    *descriptor.MethodDescriptorProto
	typesMap map[string]string
}

func (m *templateMethod) MethodName() string {
	return m.proto.GetName()
}

func (m *templateMethod) InType() string {
	return m.typesMap[m.proto.GetInputType()]
}

func (m *templateMethod) OutType() string {
	return m.typesMap[m.proto.GetOutputType()]
}

const browserTemplate = `
package {{.Package}}

import (
	{{range .Imports}}{{.}}
	{{end}}
)

// Code generated by protoc-gen-king_go {{.Version}}, DO NOT EDIT.
// Source: {{.SourceFilename}}

{{range .Services}}
{{$serviceName := .ServiceName}}
{{$implName := .ClientImplName}}

type {{.ServerName}} interface {
	{{range .Methods}}
	{{.MethodName}}(ctx context.Context, in *{{.InType}}) (out *{{.OutType}}, err error){{end}}
}

func Register{{.Name}}(server {{.ServerName}}) {
	serviceDef := &runtime.Service{
		Name: "{{.ServiceName}}",
		Methods: []*runtime.Method{
			{{range .Methods}}
				{
					Name: "{{.MethodName}}",
					Input: func() proto.Message { return new({{.InType}}) },
					Handler: func(ctx context.Context, in proto.Message) (proto.Message, error) {
						return server.{{.MethodName}}(ctx, in.(*{{.InType}}))
					},
				},
			{{end}}
		},
	}
	runtime.Services = append(runtime.Services, serviceDef)
}

type {{.ClientName}} interface {
	{{range .Methods}}
	{{.MethodName}}(ctx context.Context, in *{{.InType}}) (out *{{.OutType}}, err error){{end}}
}

 type {{$implName}} struct {
 	caller *runtime.ClientCaller
 }

func New{{.ClientName}}(server string, opts ...runtime.ClientOption) {{.ClientName}} {
	impClient := &{{$implName}}{
		caller: &runtime.ClientCaller{
			Server: server,
		},
	}

	for _, opt := range opts {
		opt(impClient.caller)
	}

	return impClient
}

{{range .Methods}}
func (impl *{{$implName}}) {{.MethodName}}(ctx context.Context, in *{{.InType}}) (out *{{.OutType}}, err error) {
	out = new({{.OutType}})
	if err := impl.caller.Call(ctx, "{{$serviceName}}", "{{.MethodName}}", in, out); err != nil {
		return nil, err
	}

	return out, nil
}
{{end}}

{{end}}
`
