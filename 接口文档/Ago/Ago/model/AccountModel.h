//
//  AccountModel.h
//  Ago
//
//  Created by rd on 2019/8/22.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN

@interface AccountModel : NSObject

@property (nonatomic, assign)NSTimeInterval BirthDay;
@property (nonatomic, assign)NSInteger Cash;
@property (nonatomic, assign)NSInteger Charm;
@property (nonatomic, assign)NSInteger Diamond;
@property (nonatomic, assign)NSInteger Gold;
@property (nonatomic, assign)NSInteger PlayerID;
@property (nonatomic, assign)NSInteger Rich;


@property (nonatomic, strong)NSString* CityName;
@property (nonatomic, strong)NSString* DisPlayerID;
@property (nonatomic, strong)NSString* PlayerName;
@property (nonatomic, strong)NSString* Profession;
@property (nonatomic, strong)NSString* School;
@property (nonatomic, strong)NSString* Sign;
@property (nonatomic, strong)NSString* StarSign;

@end

NS_ASSUME_NONNULL_END
