//
//  TQNetWokTool.m
//  Ago
//
//  Created by rd on 2019/8/22.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "TQNetWokTool.h"

@interface TQNetWokTool()

@property (nonatomic, strong)AFHTTPSessionManager *manager;
@end

@implementation TQNetWokTool

- (void)saveLoginInfo:(NSDictionary*)dic{
    
    NSData * data  = [NSKeyedArchiver archivedDataWithRootObject:dic];

    NSUserDefaults *ud = [NSUserDefaults standardUserDefaults];
    [ud setObject:data forKey:@"userInfo"];
    [ud synchronize];
    
    [self setDic:dic];
}

- (void)setDic:(NSDictionary*)dic{
    NSString * token = [[dic objectForKey:@"data"] objectForKey:@"token"];
   
    _tocken = token;
    NSDictionary *user = [[dic objectForKey:@"data"] objectForKey:@"user"];
    _account = [AccountModel yy_modelWithJSON:user];
    NSLog(@"%@",_account);
}

- (void)setAccount:(AccountModel *)account{
    _account = account;
}

- (instancetype)init
{
    self = [super init];
    if (self) {
        
        NSUserDefaults * defaults = [NSUserDefaults standardUserDefaults];
        id temp = [defaults objectForKey:@"userInfo"];
        
        if (temp) {
            [SVProgressHUD showInfoWithStatus:@"已经登录"];
            [SVProgressHUD dismissWithDelay:1];
            [self setDic:[NSKeyedUnarchiver unarchiveObjectWithData:temp]];
        }
        else{
            [SVProgressHUD showInfoWithStatus:@"请登录"];
            [SVProgressHUD dismissWithDelay:1];
        }
        
        _manager = [[AFHTTPSessionManager alloc] initWithBaseURL:[NSURL URLWithString:@"http://127.0.0.1:8080"]];
        //_manager.requestSerializer = [AFHTTPRequestSerializer serializer];
        _manager.requestSerializer = [AFJSONRequestSerializer serializer];
        _manager.responseSerializer = [AFHTTPResponseSerializer serializer];
        _manager.requestSerializer.timeoutInterval = 10;
        [_manager.requestSerializer setValue:@"application/json; charset=utf-8" forHTTPHeaderField:@"Content-Type"];
        _manager.responseSerializer.acceptableContentTypes = [NSSet setWithObjects:@"application/json", @"text/json", @"text/javascript", @"text/html", @"text/plain", nil];
    }
    return self;
}

-(NSString*)getCookie{
    
    NSString *uuid = [UIDevice currentDevice].identifierForVendor.UUIDString;
    NSString *token = self.tocken != nil? self.tocken:@"";
    NSInteger playerID = self.account.PlayerID;
    NSRegularExpression *regular = [NSRegularExpression regularExpressionWithPattern:@"[a-zA-Z.-]" options:0 error:NULL];
    NSString *versionString = [[[NSBundle mainBundle] infoDictionary] objectForKey:@"CFBundleShortVersionString"];
    NSString *result = [regular stringByReplacingMatchesInString:versionString options:0 range:NSMakeRange(0, [versionString length]) withTemplate:@""];
    NSDictionary *infoDict = [NSBundle mainBundle].infoDictionary;
    NSString *buildVersion = infoDict[@"CFBundleVersion"];
    //NSString *adId = [[[ASIdentifierManager sharedManager] advertisingIdentifier] UUIDString];
    NSString *cookie = [NSString stringWithFormat:@"playerid=%lu;device_id=D01%@;system=iOS;channel=iOS;token=%@;app_version=%@;app_version_code=%@",playerID,uuid,token,result,buildVersion];
    
    return cookie;
}

+(instancetype)instance{
    
    static id _instance = nil;
    
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        _instance = [[self alloc] init];
    });
    
    return _instance;
}

- (void)requestGet:(NSString *)url params:(id)params showLoading:(BOOL)showLoading success:(void (^)(id responseObject, SuccessCode codeType))success failure:(void (^)(NSError *error))failure{
    
    NSString *cookie = [self getCookie];
    if (cookie == nil) {
        return;
    }
    [self.manager.requestSerializer setValue:cookie forHTTPHeaderField:@"Cookie"];
    
    [self.manager GET:url parameters:params progress:^(NSProgress * _Nonnull downloadProgress) {
        
    } success:^(NSURLSessionDataTask * _Nonnull task, id  _Nullable responseObject) {
        
        /// 是否取消请求状态
        NSError *error;
        NSDictionary *dict = [NSJSONSerialization JSONObjectWithData:responseObject options:NSJSONReadingAllowFragments error:&error];
        if (dict == nil || [dict class] == [NSNull class]) {
            
            return;
        }
        SuccessCode result = [[dict objectForKey:@"code"] integerValue];
        
        success(dict, result);
        
    } failure:^(NSURLSessionDataTask * _Nullable task, NSError * _Nonnull error) {
        /// 是否取消请求状态
        failure(error);
    }];
}

- (void)requestPost:(NSString *)url params:(id)params showLoading:(BOOL)showLoading success:(void (^)(id responseObject, SuccessCode codeType))success failure:(void (^)(NSError *error))failure{
    
    NSString *cookie = [self getCookie];
    if (showLoading) {
    }
    
    [self.manager.requestSerializer setValue:cookie forHTTPHeaderField:@"Cookie"];
    
    [self.manager POST:url parameters:params progress:^(NSProgress * _Nonnull uploadProgress) {
        
        
    } success:^(NSURLSessionDataTask * _Nonnull task, id  _Nullable responseObject) {
        /// 是否取消请求状态
        NSError *error;
        NSDictionary *dict = [NSJSONSerialization JSONObjectWithData:responseObject options:NSJSONReadingAllowFragments error:&error];
        if (dict == nil || [dict class] == [NSNull class]) {
            
            return;
        }
        SuccessCode result = [[dict objectForKey:@"code"] integerValue];
        
        success(dict, result);
    } failure:^(NSURLSessionDataTask * _Nullable task, NSError * _Nonnull error) {
        /// 是否取消请求状态
        failure(error);
    }];
}

@end
