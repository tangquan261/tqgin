//
//  AgoraRTMManager.m
//  Ago
//
//  Created by rd on 2019/8/22.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "AgoraRTMManager.h"

#import "TQNetWokTool.h"
#import "NSStringUtil.h"
@interface AgoraRTMManager()<AgoraRtmChannelDelegate, AgoraRtmDelegate>

//RTM私聊
@property (nonatomic, strong)AgoraRtmKit *kit;
@property (nonatomic, strong)AgoraRtmChannel *rtmChannel;
@end

@implementation AgoraRTMManager

- (instancetype)init
{
    self = [super init];
    
    if (self) {
        _kit = [[AgoraRtmKit alloc] initWithAppId:@"1f836f0e094446d2858f156ca366313d" delegate:self];
        _currentChannelMember = [NSMutableArray arrayWithCapacity:10];
        _currentChannelMsg = [NSMutableArray arrayWithCapacity:100];
        _arrayMics = [NSMutableArray arrayWithCapacity:8];
    }
    
    return self;
}

-  (void)setArrayMics:(NSMutableArray<MicModel *> *)arrayMics{
    [_arrayMics  removeAllObjects];
    [_arrayMics addObjectsFromArray:arrayMics];
}


- (void)login:(NSString*)userID trmToken:(NSString*)rtmtoken{
    [_kit loginByToken:rtmtoken user:userID completion:^(AgoraRtmLoginErrorCode errorCode) {
        
    }];
}

- (void)loginOut{
    [self.kit logoutWithCompletion:^(AgoraRtmLogoutErrorCode errorCode) {
        
    }];
}

- (void)sendPersonMsg:(NSInteger)playerID content:(NSString*)content{
    
    AgoraRtmSendMessageOptions *option = [[AgoraRtmSendMessageOptions alloc] init];
    option.enableOfflineMessaging = YES;
    
    [self.kit sendMessage:[[AgoraRtmMessage alloc] initWithText:content] toPeer:[NSString stringWithFormat:@"%lu",playerID] sendMessageOptions:option completion:^(AgoraRtmSendPeerMessageErrorCode errorCode) {
        
    }];
}


- (void)enterChannel:(NSString*)channelID block:(AgoraRtmJoinChannelBlock _Nullable)completionBlock{
    
    if (_rtmChannel) {
        [_rtmChannel leaveWithCompletion:^(AgoraRtmLeaveChannelErrorCode errorCode) {
            
        }];
        _rtmChannel = nil;
    }
    
    _rtmChannel = [_kit createChannelWithId:channelID delegate:self];
    
    [_rtmChannel joinWithCompletion:^(AgoraRtmJoinChannelErrorCode errorCode) {
        if(errorCode == AgoraRtmJoinChannelErrorOk) {
            NSLog(@"join success");
        } else {
            NSLog(@"join failed: %@", @(errorCode));
        }
        
        if (completionBlock) {
            completionBlock(errorCode);
        }
    }];
}

- (void)leaveChannel{
    
    if (_rtmChannel) {
        [_rtmChannel leaveWithCompletion:^(AgoraRtmLeaveChannelErrorCode errorCode) {
            
        }];
        _rtmChannel = nil;
    }
    
    [_currentChannelMember removeAllObjects];
    [_currentChannelMsg removeAllObjects];
}

- (void)sendChannelMsg:(NSDictionary*)dic{
    if (_rtmChannel == nil) {
        return;
    }
    
    NSString *json = [NSStringUtil convertToJsonData:dic];
    AgoraRtmMessage *msg = [[AgoraRtmMessage alloc] initWithText:json];
    
    [_rtmChannel sendMessage:msg completion:^(AgoraRtmSendChannelMessageErrorCode errorCode) {
        [_currentChannelMsg addObject:msg];
        
         [[NSNotificationCenter defaultCenter] postNotificationName:CHANNEL_MSG_UPDATE object:nil];
    }];
}

- (void)getChannetlMember:(AgoraRtmGetMembersBlock _Nullable)completionBlock{
    
    if (_rtmChannel == nil) {
        return;
    }
    
    [_rtmChannel getMembersWithCompletion:^(NSArray<AgoraRtmMember *> * _Nullable members, AgoraRtmGetMembersErrorCode errorCode) {
       
        if ([members count]) {
            [_currentChannelMember addObjectsFromArray:members];
        }
        
        if (completionBlock) {
            completionBlock(members,errorCode);
        }
    }];
}

+(instancetype)instance{
    
    static id _instance = nil;
    
    static dispatch_once_t onceToken;
    dispatch_once(&onceToken, ^{
        _instance = [[self alloc] init];
    });
    
    return _instance;
}

- (void)rtmKit:(AgoraRtmKit * _Nonnull)kit connectionStateChanged:(AgoraRtmConnectionState)state reason:(AgoraRtmConnectionChangeReason)reason
{
    NSLog(@"Connection state changed to %@", @(reason));
}

//群聊房间进入
- (void)channel:(AgoraRtmChannel * _Nonnull)channel memberJoined:(AgoraRtmMember * _Nonnull)member{
    NSLog(@"%@ joined channel %@", member.userId, member.channelId);
    [_currentChannelMember addObject:member];
    [[NSNotificationCenter defaultCenter] postNotificationName:CHANNEL_USERS_UPDATE object:nil];
}

//群聊房间退出
- (void)channel:(AgoraRtmChannel * _Nonnull)channel memberLeft:(AgoraRtmMember * _Nonnull)member{
    NSLog(@"%@ left channel %@", member.userId, member.channelId);
    
    [self removeByID:member.userId];
    
    [[NSNotificationCenter defaultCenter] postNotificationName:CHANNEL_USERS_UPDATE object:nil];
}

//获取到群聊消息
- (void)channel:(AgoraRtmChannel * _Nonnull)channel messageReceived:(AgoraRtmMessage * _Nonnull)message fromMember:(AgoraRtmMember * _Nonnull)member{
    NSLog(@"Message received from %@: %@", message.text, channel);
    
    NSDictionary *dic = [message.text yy_modelToJSONObject];
    
    NSInteger type = [[dic objectForKey:@"type"] integerValue];
    
    if(type == 100){
        
        NSInteger subType = [[dic objectForKey:@"subtype"] integerValue];
        
        if (subType == 1){
            
            [self.arrayMics removeAllObjects];
            
            NSArray* mics = [dic objectForKey:@"mics"];
            if (mics) {
                self.arrayMics =  [NSArray yy_modelArrayWithClass:[MicModel class] json:mics];
                
                [[NSNotificationCenter defaultCenter] postNotificationName:CHANNEL_MICS_UPDATE object:nil];
            }
            
        }
        
        
    }
    else{
        [_currentChannelMsg addObject:message];
        [[NSNotificationCenter defaultCenter] postNotificationName:CHANNEL_MSG_UPDATE object:nil];
    }
}

//获取到点对点消息
- (void)rtmKit:(AgoraRtmKit * _Nonnull)kit messageReceived:(AgoraRtmMessage * _Nonnull)message fromPeer:(NSString * _Nonnull)peerId{
    NSLog(@"Message received from %@: %@", message.text, peerId);
}

- (void)removeByID:(NSString*)uid{
    NSMutableArray *arrayDel = [NSMutableArray arrayWithCapacity:2];
    for (AgoraRtmMember *obj in _currentChannelMember) {
        if ([obj.userId isEqualToString:uid]) {
            [arrayDel addObject:obj];
        }
    }
    [_currentChannelMember removeObjectsInArray:arrayDel];
}
@end
