//
//  TQNetWokTool.h
//  Ago
//
//  Created by rd on 2019/8/22.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import <Foundation/Foundation.h>
#import "AFNetworking.h"
#import "AccountModel.h"
#import "YYModel.h"
#import "SVProgressHUD.h"
NS_ASSUME_NONNULL_BEGIN

typedef NS_ENUM(NSUInteger, RequestMethod) {
    GET,
    POST,
};

typedef NS_ENUM(NSUInteger, SuccessCode){
    SuccessCode_Success = 0,        //操作正确返回
    SuccessCode_UnknownError = 3,   //未知错误
    SuccessCode_TimeOut = 4,        //读取超时
    SuccessCode_LogOut = 5,         //未登陆，比如被踢下线
    SuccessCode_LoginExpire  = 1011,//登陆过期
    SuccessCode_DianZan      = 2003,//已经点过赞
    SuccessCode_GroupNotExist = 3006,//家族不存在
    SuccessCode_ChatisNotOpen = 3102,//群聊未在开放时间不支持
    SuccessCode_MaxTime = 6001,      // 信号弹发送次数已用完
    SuccessCode_isReleasing = 6002,  // 信号弹释放中
    SuccessCode_Other,
};

@interface TQNetWokTool : NSObject

+(instancetype)instance;

- (void)saveLoginInfo:(NSDictionary*)dic;

@property (nonatomic, copy)NSString *tocken;
@property (nonatomic, strong)AccountModel *account;

- (void)requestGet:(NSString *)url params:(id)params showLoading:(BOOL)showLoading success:(void (^)(id responseObject, SuccessCode codeType))success failure:(void (^)(NSError *error))failure;

- (void)requestPost:(NSString *)url params:(id)params showLoading:(BOOL)showLoading success:(void (^)(id responseObject, SuccessCode codeType))success failure:(void (^)(NSError *error))failure;

@end

NS_ASSUME_NONNULL_END
