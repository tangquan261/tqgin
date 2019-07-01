// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: roominfo.proto

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

#import "Roominfo.pbobjc.h"
#import "Userinfo.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - RoominfoRoot

@implementation RoominfoRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - RoominfoRoot_FileDescriptor

static GPBFileDescriptor *RoominfoRoot_FileDescriptor(void) {
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

#pragma mark - Enum roomType

GPBEnumDescriptor *roomType_EnumDescriptor(void) {
  static _Atomic(GPBEnumDescriptor*) descriptor = nil;
  if (!descriptor) {
    static const char *valueNames =
        "RoomTypeNoaml\000RoomTypeSecond\000";
    static const int32_t values[] = {
        roomType_RoomTypeNoaml,
        roomType_RoomTypeSecond,
    };
    static const char *extraTextFormatInfo = "\002\000(\345\000\001(\346\000";
    GPBEnumDescriptor *worker =
        [GPBEnumDescriptor allocDescriptorForName:GPBNSStringifySymbol(roomType)
                                       valueNames:valueNames
                                           values:values
                                            count:(uint32_t)(sizeof(values) / sizeof(int32_t))
                                     enumVerifier:roomType_IsValidValue
                              extraTextFormatInfo:extraTextFormatInfo];
    GPBEnumDescriptor *expected = nil;
    if (!atomic_compare_exchange_strong(&descriptor, &expected, worker)) {
      [worker release];
    }
  }
  return descriptor;
}

BOOL roomType_IsValidValue(int32_t value__) {
  switch (value__) {
    case roomType_RoomTypeNoaml:
    case roomType_RoomTypeSecond:
      return YES;
    default:
      return NO;
  }
}

#pragma mark - ApplyCreateRoom

@implementation ApplyCreateRoom

@dynamic roomName;
@dynamic roomTags;
@dynamic roomtype;

typedef struct ApplyCreateRoom__storage_ {
  uint32_t _has_storage_[1];
  roomType roomtype;
  NSString *roomName;
  NSString *roomTags;
} ApplyCreateRoom__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "roomName",
        .dataTypeSpecific.className = NULL,
        .number = ApplyCreateRoom_FieldNumber_RoomName,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(ApplyCreateRoom__storage_, roomName),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "roomTags",
        .dataTypeSpecific.className = NULL,
        .number = ApplyCreateRoom_FieldNumber_RoomTags,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(ApplyCreateRoom__storage_, roomTags),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "roomtype",
        .dataTypeSpecific.enumDescFunc = roomType_EnumDescriptor,
        .number = ApplyCreateRoom_FieldNumber_Roomtype,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(ApplyCreateRoom__storage_, roomtype),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldHasEnumDescriptor),
        .dataType = GPBDataTypeEnum,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[ApplyCreateRoom class]
                                     rootClass:[RoominfoRoot class]
                                          file:RoominfoRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(ApplyCreateRoom__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\002\001\010\000\002\010\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end

int32_t ApplyCreateRoom_Roomtype_RawValue(ApplyCreateRoom *message) {
  GPBDescriptor *descriptor = [ApplyCreateRoom descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:ApplyCreateRoom_FieldNumber_Roomtype];
  return GPBGetMessageInt32Field(message, field);
}

void SetApplyCreateRoom_Roomtype_RawValue(ApplyCreateRoom *message, int32_t value) {
  GPBDescriptor *descriptor = [ApplyCreateRoom descriptor];
  GPBFieldDescriptor *field = [descriptor fieldWithNumber:ApplyCreateRoom_FieldNumber_Roomtype];
  GPBSetInt32IvarWithFieldInternal(message, field, value, descriptor.file.syntax);
}

#pragma mark - applyEnterRoom

@implementation applyEnterRoom

@dynamic roomId;
@dynamic roomName;
@dynamic urmIcon;
@dynamic masterId;
@dynamic roomUserInfoArray, roomUserInfoArray_Count;

typedef struct applyEnterRoom__storage_ {
  uint32_t _has_storage_[1];
  NSString *roomName;
  NSString *urmIcon;
  NSMutableArray *roomUserInfoArray;
  int64_t roomId;
  int64_t masterId;
} applyEnterRoom__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "roomId",
        .dataTypeSpecific.className = NULL,
        .number = applyEnterRoom_FieldNumber_RoomId,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(applyEnterRoom__storage_, roomId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "roomName",
        .dataTypeSpecific.className = NULL,
        .number = applyEnterRoom_FieldNumber_RoomName,
        .hasIndex = 1,
        .offset = (uint32_t)offsetof(applyEnterRoom__storage_, roomName),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "urmIcon",
        .dataTypeSpecific.className = NULL,
        .number = applyEnterRoom_FieldNumber_UrmIcon,
        .hasIndex = 2,
        .offset = (uint32_t)offsetof(applyEnterRoom__storage_, urmIcon),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeString,
      },
      {
        .name = "masterId",
        .dataTypeSpecific.className = NULL,
        .number = applyEnterRoom_FieldNumber_MasterId,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(applyEnterRoom__storage_, masterId),
        .flags = (GPBFieldFlags)(GPBFieldOptional | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeInt64,
      },
      {
        .name = "roomUserInfoArray",
        .dataTypeSpecific.className = GPBStringifySymbol(RoomUserInfo),
        .number = applyEnterRoom_FieldNumber_RoomUserInfoArray,
        .hasIndex = GPBNoHasBit,
        .offset = (uint32_t)offsetof(applyEnterRoom__storage_, roomUserInfoArray),
        .flags = (GPBFieldFlags)(GPBFieldRepeated | GPBFieldTextFormatNameCustom),
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[applyEnterRoom class]
                                     rootClass:[RoominfoRoot class]
                                          file:RoominfoRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(applyEnterRoom__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
#if !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    static const char *extraTextFormatInfo =
        "\005\001\005A\000\002\010\000\003\007\000\004\007A\000\005\000RoomUserInfo\000";
    [localDescriptor setupExtraTextInfo:extraTextFormatInfo];
#endif  // !GPBOBJC_SKIP_MESSAGE_TEXTFORMAT_EXTRAS
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
