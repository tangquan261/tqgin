// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Squart.proto

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

@class BannerInfo;
@class HotRoomInfo;
@class TagInfo;

NS_ASSUME_NONNULL_BEGIN

#pragma mark - SquartRoot

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
@interface SquartRoot : GPBRootObject
@end

#pragma mark - TagInfo

typedef GPB_ENUM(TagInfo_FieldNumber) {
  TagInfo_FieldNumber_Id_p = 1,
  TagInfo_FieldNumber_TagName = 2,
};

@interface TagInfo : GPBMessage

@property(nonatomic, readwrite) int64_t id_p;

/** 房间分类tag */
@property(nonatomic, readwrite, copy, null_resettable) NSString *tagName;

@end

#pragma mark - TagsInfo

typedef GPB_ENUM(TagsInfo_FieldNumber) {
  TagsInfo_FieldNumber_TagInfoArray = 1,
};

@interface TagsInfo : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<TagInfo*> *tagInfoArray;
/** The number of items in @c tagInfoArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger tagInfoArray_Count;

@end

#pragma mark - HotRoomInfo

typedef GPB_ENUM(HotRoomInfo_FieldNumber) {
  HotRoomInfo_FieldNumber_RoomId = 1,
  HotRoomInfo_FieldNumber_RoomType = 2,
  HotRoomInfo_FieldNumber_RoomName = 3,
  HotRoomInfo_FieldNumber_Pic = 4,
  HotRoomInfo_FieldNumber_Intro = 5,
  HotRoomInfo_FieldNumber_Password = 6,
  HotRoomInfo_FieldNumber_RoomTagName = 7,
  HotRoomInfo_FieldNumber_MemCount = 8,
  HotRoomInfo_FieldNumber_RoomHot = 9,
};

/**
 * 热门房间
 **/
@interface HotRoomInfo : GPBMessage

@property(nonatomic, readwrite) int64_t roomId;

@property(nonatomic, readwrite) int32_t roomType;

@property(nonatomic, readwrite, copy, null_resettable) NSString *roomName;

@property(nonatomic, readwrite, copy, null_resettable) NSString *pic;

@property(nonatomic, readwrite, copy, null_resettable) NSString *intro;

@property(nonatomic, readwrite, copy, null_resettable) NSString *password;

@property(nonatomic, readwrite, copy, null_resettable) NSString *roomTagName;

@property(nonatomic, readwrite) int32_t memCount;

@property(nonatomic, readwrite) int64_t roomHot;

@end

#pragma mark - HotRooms

typedef GPB_ENUM(HotRooms_FieldNumber) {
  HotRooms_FieldNumber_HotRoomInfoArray = 1,
};

@interface HotRooms : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<HotRoomInfo*> *hotRoomInfoArray;
/** The number of items in @c hotRoomInfoArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger hotRoomInfoArray_Count;

@end

#pragma mark - BannerInfo

typedef GPB_ENUM(BannerInfo_FieldNumber) {
  BannerInfo_FieldNumber_BannerId = 1,
  BannerInfo_FieldNumber_TargetType = 2,
  BannerInfo_FieldNumber_StartTime = 3,
  BannerInfo_FieldNumber_EndTime = 4,
  BannerInfo_FieldNumber_BgImg = 5,
  BannerInfo_FieldNumber_ClickURL = 6,
};

@interface BannerInfo : GPBMessage

@property(nonatomic, readwrite) int32_t bannerId;

@property(nonatomic, readwrite) int32_t targetType;

@property(nonatomic, readwrite) int64_t startTime;

@property(nonatomic, readwrite) int64_t endTime;

@property(nonatomic, readwrite, copy, null_resettable) NSString *bgImg;

@property(nonatomic, readwrite, copy, null_resettable) NSString *clickURL;

@end

#pragma mark - Banners

typedef GPB_ENUM(Banners_FieldNumber) {
  Banners_FieldNumber_BannerInfoArray = 1,
};

@interface Banners : GPBMessage

@property(nonatomic, readwrite, strong, null_resettable) NSMutableArray<BannerInfo*> *bannerInfoArray;
/** The number of items in @c bannerInfoArray without causing the array to be created. */
@property(nonatomic, readonly) NSUInteger bannerInfoArray_Count;

@end

NS_ASSUME_NONNULL_END

CF_EXTERN_C_END

#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
