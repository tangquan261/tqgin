//
//  AgoraRTMManager.h
//  Ago
//
//  Created by rd on 2019/8/22.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import <Foundation/Foundation.h>
#import <AgoraRtmKit/AgoraRtmKit.h>
#import "MicModel.h"

NS_ASSUME_NONNULL_BEGIN
static NSString *CHANNEL_USERS_UPDATE = @"CHANNEL_USERS_UPDATE";
static NSString *CHANNEL_MSG_UPDATE = @"CHANNEL_MSG_UPDATE";
static NSString *CHANNEL_MICS_UPDATE = @"CHANNEL_MICS_UPDATE";


@interface AgoraRTMManager : NSObject

+(instancetype)instance;

@property (nonatomic, strong)NSMutableArray *currentChannelMember;
@property (nonatomic, strong)NSMutableArray *currentChannelMsg;

@property (nonatomic, strong)NSMutableArray<MicModel*>*arrayMics;


- (void)login:(NSString*)userID trmToken:(NSString*)rtmtoken;
- (void)loginOut;

- (void)sendPersonMsg:(NSInteger)playerID content:(NSString*)content;

- (void)enterChannel:(NSString*)channelID block:(AgoraRtmJoinChannelBlock _Nullable)completionBlock;
- (void)leaveChannel;

- (void)sendChannelMsg:(NSDictionary*)conten;

- (void)getChannetlMember:(AgoraRtmGetMembersBlock _Nullable)completionBlock;

@end

NS_ASSUME_NONNULL_END
