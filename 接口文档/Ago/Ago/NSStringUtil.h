//
//  NSStringUtil.h
//  Ago
//
//  Created by rd on 2019/8/22.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

@interface NSStringUtil : NSString

+(NSString *)convertToJsonData:(NSDictionary *)dict;

+ (NSDictionary *)dictionaryWithJsonString:(NSString *)jsonString;
@end

NS_ASSUME_NONNULL_END
