// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: login.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers.h>
#else
 #import "GPBProtocolBuffers.h"
#endif

#if GOOGLE_PROTOBUF_OBJC_VERSION < 30002
#error This file was generated by a newer version of protoc which is incompatible with your Protocol Buffer library sources.
#endif
#if 30002 < GOOGLE_PROTOBUF_OBJC_MIN_SUPPORTED_VERSION
#error This file was generated by an older version of protoc which is incompatible with your Protocol Buffer library sources.
#endif

// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

CF_EXTERN_C_BEGIN

NS_ASSUME_NONNULL_BEGIN

#pragma mark - Enum ApplyLogin_LoginType

typedef GPB_ENUM(ApplyLogin_LoginType) {
  /**
   * Value used if any message's field encounters a value that is not defined
   * by this enum. The message will also have C functions to get/set the rawValue
   * of the field.
   **/
  ApplyLogin_LoginType_GPBUnrecognizedEnumeratorValue = kGPBUnrecognizedEnumeratorValue,
  ApplyLogin_LoginType_Mobile = 0,
  ApplyLogin_LoginType_Home = 1,
  ApplyLogin_LoginType_Work = 2,
};

GPBEnumDescriptor *ApplyLogin_LoginType_EnumDescriptor(void);

/**
 * Checks to see if the given value is defined by the enum or was not known at
 * the time this source was generated.
 **/
BOOL ApplyLogin_LoginType_IsValidValue(int32_t value);

#pragma mark - LoginRoot

/**
 * Exposes the extension registry for this file.
 *
 * The base class provides:
 * @code
 *   + (GPBExtensionRegistry *)extensionRegistry;
 * @endcode
 * which is a @c GPBExtensionRegistry that includes all the extensions defined by
 * this file and all files that it depends on.
 **/
@interface LoginRoot : GPBRootObject
@end

#pragma mark - ApplyLogin

typedef GPB_ENUM(ApplyLogin_FieldNumber) {
  ApplyLogin_FieldNumber_Account = 1,
  ApplyLogin_FieldNumber_Password = 2,
  ApplyLogin_FieldNumber_Type = 3,
};

/**
 * 登陆请求
 **/
@interface ApplyLogin : GPBMessage

/** 登陆账户 */
@property(nonatomic, readwrite, copy, null_resettable) NSString *account;

/** 登陆密码 */
@property(nonatomic, readwrite, copy, null_resettable) NSString *password;

/** 登陆类型 */
@property(nonatomic, readwrite) ApplyLogin_LoginType type;

@end

/**
 * Fetches the raw value of a @c ApplyLogin's @c type property, even
 * if the value was not defined by the enum at the time the code was generated.
 **/
int32_t ApplyLogin_Type_RawValue(ApplyLogin *message);
/**
 * Sets the raw value of an @c ApplyLogin's @c type property, allowing
 * it to be set to a value that was not defined by the enum at the time the code
 * was generated.
 **/
void SetApplyLogin_Type_RawValue(ApplyLogin *message, int32_t value);

#pragma mark - ReplyLogin

typedef GPB_ENUM(ReplyLogin_FieldNumber) {
  ReplyLogin_FieldNumber_Errinfo = 1,
  ReplyLogin_FieldNumber_Code = 2,
};

/**
 * 登陆请求返回
 **/
@interface ReplyLogin : GPBMessage

@property(nonatomic, readwrite, copy, null_resettable) NSString *errinfo;

@property(nonatomic, readwrite) int32_t code;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
