// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/config/enforcer/jwt_generator.proto

package enforcer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// JWT Generator model
type JWTGenerator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable               bool       `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Encoding             string     `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
	ClaimDialect         string     `protobuf:"bytes,3,opt,name=claim_dialect,json=claimDialect,proto3" json:"claim_dialect,omitempty"`
	ConvertDialect       bool       `protobuf:"varint,4,opt,name=convert_dialect,json=convertDialect,proto3" json:"convert_dialect,omitempty"`
	Header               string     `protobuf:"bytes,5,opt,name=header,proto3" json:"header,omitempty"`
	SigningAlgorithm     string     `protobuf:"bytes,6,opt,name=signing_algorithm,json=signingAlgorithm,proto3" json:"signing_algorithm,omitempty"`
	EnableUserClaims     bool       `protobuf:"varint,7,opt,name=enable_user_claims,json=enableUserClaims,proto3" json:"enable_user_claims,omitempty"`
	GatewayGeneratorImpl string     `protobuf:"bytes,8,opt,name=gateway_generator_impl,json=gatewayGeneratorImpl,proto3" json:"gateway_generator_impl,omitempty"`
	ClaimsExtractorImpl  string     `protobuf:"bytes,9,opt,name=claims_extractor_impl,json=claimsExtractorImpl,proto3" json:"claims_extractor_impl,omitempty"`
	TokenTtl             int32      `protobuf:"varint,10,opt,name=token_ttl,json=tokenTtl,proto3" json:"token_ttl,omitempty"`
	Keypairs             []*Keypair `protobuf:"bytes,11,rep,name=keypairs,proto3" json:"keypairs,omitempty"`
}

func (x *JWTGenerator) Reset() {
	*x = JWTGenerator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_config_enforcer_jwt_generator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTGenerator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTGenerator) ProtoMessage() {}

func (x *JWTGenerator) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_config_enforcer_jwt_generator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTGenerator.ProtoReflect.Descriptor instead.
func (*JWTGenerator) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescGZIP(), []int{0}
}

func (x *JWTGenerator) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *JWTGenerator) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *JWTGenerator) GetClaimDialect() string {
	if x != nil {
		return x.ClaimDialect
	}
	return ""
}

func (x *JWTGenerator) GetConvertDialect() bool {
	if x != nil {
		return x.ConvertDialect
	}
	return false
}

func (x *JWTGenerator) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *JWTGenerator) GetSigningAlgorithm() string {
	if x != nil {
		return x.SigningAlgorithm
	}
	return ""
}

func (x *JWTGenerator) GetEnableUserClaims() bool {
	if x != nil {
		return x.EnableUserClaims
	}
	return false
}

func (x *JWTGenerator) GetGatewayGeneratorImpl() string {
	if x != nil {
		return x.GatewayGeneratorImpl
	}
	return ""
}

func (x *JWTGenerator) GetClaimsExtractorImpl() string {
	if x != nil {
		return x.ClaimsExtractorImpl
	}
	return ""
}

func (x *JWTGenerator) GetTokenTtl() int32 {
	if x != nil {
		return x.TokenTtl
	}
	return 0
}

func (x *JWTGenerator) GetKeypairs() []*Keypair {
	if x != nil {
		return x.Keypairs
	}
	return nil
}

var File_wso2_discovery_config_enforcer_jwt_generator_proto protoreflect.FileDescriptor

var file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDesc = []byte{
	0x0a, 0x32, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72,
	0x2f, 0x6a, 0x77, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x1a, 0x2c, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x2f, 0x6b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcf, 0x03, 0x0a, 0x0c, 0x4a, 0x57, 0x54, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x5f, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x44, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x44, 0x69,
	0x61, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a,
	0x11, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e,
	0x67, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6d,
	0x70, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6d, 0x70, 0x6c, 0x12, 0x32,
	0x0a, 0x15, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6d,
	0x70, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x74, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x74, 0x6c, 0x12,
	0x43, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x70, 0x61, 0x69, 0x72, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x70,
	0x61, 0x69, 0x72, 0x73, 0x42, 0x98, 0x01, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x42, 0x11, 0x4a, 0x57, 0x54, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x3b, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescOnce sync.Once
	file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescData = file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDesc
)

func file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescGZIP() []byte {
	file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescData)
	})
	return file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDescData
}

var file_wso2_discovery_config_enforcer_jwt_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_wso2_discovery_config_enforcer_jwt_generator_proto_goTypes = []interface{}{
	(*JWTGenerator)(nil), // 0: wso2.discovery.config.enforcer.JWTGenerator
	(*Keypair)(nil),      // 1: wso2.discovery.config.enforcer.Keypair
}
var file_wso2_discovery_config_enforcer_jwt_generator_proto_depIdxs = []int32{
	1, // 0: wso2.discovery.config.enforcer.JWTGenerator.keypairs:type_name -> wso2.discovery.config.enforcer.Keypair
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wso2_discovery_config_enforcer_jwt_generator_proto_init() }
func file_wso2_discovery_config_enforcer_jwt_generator_proto_init() {
	if File_wso2_discovery_config_enforcer_jwt_generator_proto != nil {
		return
	}
	file_wso2_discovery_config_enforcer_keypair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_config_enforcer_jwt_generator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTGenerator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_config_enforcer_jwt_generator_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_config_enforcer_jwt_generator_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_config_enforcer_jwt_generator_proto_msgTypes,
	}.Build()
	File_wso2_discovery_config_enforcer_jwt_generator_proto = out.File
	file_wso2_discovery_config_enforcer_jwt_generator_proto_rawDesc = nil
	file_wso2_discovery_config_enforcer_jwt_generator_proto_goTypes = nil
	file_wso2_discovery_config_enforcer_jwt_generator_proto_depIdxs = nil
}
