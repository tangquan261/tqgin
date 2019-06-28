// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: login.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

#import <stdatomic.h>

#import "Login.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - LoginRoot

@implementation LoginRoot

// No extensions in the file and no imports, so no need to generate
// +extensionRegistry.

@end

#pragma mark - LoginRoot_FileDescriptor

static GPBFileDescriptor *LoginRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"login"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - Enum LoginType

GPBEnumDescriptor *LoginType_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "MobileLogin\000QqLogin\000WeixinLogin\000";
    static const int32_t values[] = {
        LoginType_MobileLogin,
        LoginType_QqLogin,
        LoginType_WeixinLogin,
    };
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(LoginType)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:LoginType_IsValidValue];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL LoginType_IsValidValue(int32_t value__) {
  switch (value__) {
    case LoginType_MobileLogin:
    case LoginType_QqLogin:
    case LoginType_WeixinLogin:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - Enum SexType

GPBEnumDescriptor *SexType_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "SexMale\000SexFemale\000";
    static const int32_t values[] = {
        SexType_SexMale,
        SexType_SexFemale,
    };
    static const char *extraTextFormatInfo = "\002\000\003\244\000\001\003\246\000";
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(SexType)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:SexType_IsValidValue
                              extraTextFormatInfo:extraTextFormatInfo];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL SexType_IsValidValue(int32_t value__) {
  switch (value__) {
    case SexType_SexMale:
    case SexType_SexFemale:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - ApplyLogin

@implementation ApplyLogin

@dynamic account;
@dynamic password;
@dynamic type;

typedef struct ApplyLogin__storage_ {
  uint32_t _has_storage_[1];
  LoginType type;
  NSString *account;
  NSString *password;
} ApplyLogin__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "account",
        .dataTypeSpecific.className = NULL,
        .number = ApplyLogin_FieldNumber_Account,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(ApplyLogin__storage_, account),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "password",
        .dataTypeSpecific.className = NULL,
        .number = ApplyLogin_FieldNumber_Password,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(ApplyLogin__storage_, password),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "type",
        .dataTypeSpecific.enumDescFunc = LoginType_EnumDescriptor,
        .number = ApplyLogin_FieldNumber_Type,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(ApplyLogin__storage_, type),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[ApplyLogin class]
                                     rootClass:[LoginRoot class]
                                          file:LoginRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(ApplyLogin__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t ApplyLogin_Type_RawValue(ApplyLogin *message) {
  GPBDescriptor *descriptor = [ApplyLogin descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:ApplyLogin_FieldNumber_Type];
  return GPBGetMessageInt32Field(message, field);
}

void SetApplyLogin_Type_RawValue(ApplyLogin *message, int32_t value) {
  GPBDescriptor *descriptor = [ApplyLogin descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:ApplyLogin_FieldNumber_Type];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

#pragma mark - ApplyLoginInfo

@implementation ApplyLoginInfo


typedef struct ApplyLoginInfo__storage_ {
  uint32_t _has_storage_[1];
} ApplyLoginInfo__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[ApplyLoginInfo class]
                                     rootClass:[LoginRoot class]
                                          file:LoginRoot_FileDescriptor()
                                        fields:NULL
                                    fieldCount:0
                                   storageSize:sizeof(ApplyLoginInfo__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

#pragma mark - ReplyLogin

@implementation ReplyLogin

@dynamic playerId;
@dynamic playerName;
@dynamic diamond;
@dynamic gold;
@dynamic cash;
@dynamic roomId;
@dynamic sex;
@dynamic token;

typedef struct ReplyLogin__storage_ {
  uint32_t _has_storage_[1];
  SexType sex;
  NSString *playerName;
  NSString *token;
  int64_t playerId;
  int64_t diamond;
  int64_t gold;
  int64_t cash;
  int64_t roomId;
} ReplyLogin__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "playerId",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_PlayerId,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, playerId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "playerName",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_PlayerName,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, playerName),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "diamond",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_Diamond,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, diamond),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "gold",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_Gold,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, gold),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "cash",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_Cash,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, cash),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "roomId",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_RoomId,
        .hasIndex = 5,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, roomId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "sex",
        .dataTypeSpecific.enumDescFunc = SexType_EnumDescriptor,
        .number = ReplyLogin_FieldNumber_Sex,
        .hasIndex = 6,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, sex),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "token",
        .dataTypeSpecific.className = NULL,
        .number = ReplyLogin_FieldNumber_Token,
        .hasIndex = 7,
        .offset = (uint32_t)offsetof(ReplyLogin__storage_, token),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[ReplyLogin class]
                                     rootClass:[LoginRoot class]
                                          file:LoginRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(ReplyLogin__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\007\001\007A\000\002\n\000\003G\000\004D\000\005D\000\006EA\000\007C\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t ReplyLogin_Sex_RawValue(ReplyLogin *message) {
  GPBDescriptor *descriptor = [ReplyLogin descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:ReplyLogin_FieldNumber_Sex];
  return GPBGetMessageInt32Field(message, field);
}

void SetReplyLogin_Sex_RawValue(ReplyLogin *message, int32_t value) {
  GPBDescriptor *descriptor = [ReplyLogin descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:ReplyLogin_FieldNumber_Sex];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

#pragma mark - RegisterInfo

@implementation RegisterInfo

@dynamic account;
@dynamic password;
@dynamic type;
@dynamic nickNmae;
@dynamic sexType;

typedef struct RegisterInfo__storage_ {
  uint32_t _has_storage_[1];
  LoginType type;
  SexType sexType;
  NSString *account;
  NSString *password;
  NSString *nickNmae;
} RegisterInfo__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "account",
        .dataTypeSpecific.className = NULL,
        .number = RegisterInfo_FieldNumber_Account,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(RegisterInfo__storage_, account),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "password",
        .dataTypeSpecific.className = NULL,
        .number = RegisterInfo_FieldNumber_Password,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(RegisterInfo__storage_, password),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "type",
        .dataTypeSpecific.enumDescFunc = LoginType_EnumDescriptor,
        .number = RegisterInfo_FieldNumber_Type,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(RegisterInfo__storage_, type),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
      {
        .name = "nickNmae",
        .dataTypeSpecific.className = NULL,
        .number = RegisterInfo_FieldNumber_NickNmae,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(RegisterInfo__storage_, nickNmae),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "sexType",
        .dataTypeSpecific.enumDescFunc = SexType_EnumDescriptor,
        .number = RegisterInfo_FieldNumber_SexType,
        .hasIndex = 4,
        .offset = (uint32_t)offsetof(RegisterInfo__storage_, sexType),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[RegisterInfo class]
                                     rootClass:[LoginRoot class]
                                          file:LoginRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(RegisterInfo__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\002\004\010\000\005\007\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t RegisterInfo_Type_RawValue(RegisterInfo *message) {
  GPBDescriptor *descriptor = [RegisterInfo descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:RegisterInfo_FieldNumber_Type];
  return GPBGetMessageInt32Field(message, field);
}

void SetRegisterInfo_Type_RawValue(RegisterInfo *message, int32_t value) {
  GPBDescriptor *descriptor = [RegisterInfo descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:RegisterInfo_FieldNumber_Type];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

int32_t RegisterInfo_SexType_RawValue(RegisterInfo *message) {
  GPBDescriptor *descriptor = [RegisterInfo descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:RegisterInfo_FieldNumber_SexType];
  return GPBGetMessageInt32Field(message, field);
}

void SetRegisterInfo_SexType_RawValue(RegisterInfo *message, int32_t value) {
  GPBDescriptor *descriptor = [RegisterInfo descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:RegisterInfo_FieldNumber_SexType];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
