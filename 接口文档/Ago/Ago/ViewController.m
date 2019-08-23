//
//  ViewController.m
//  zego
//
//  Created by rd on 2019/8/9.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "ViewController.h"
#import "SVProgressHUD.h"
#import "Masonry.h"
#import "MicoView.h"
#import "MsgView.h"
#import "MsgViewController.h"
#import "TQNetWokTool.h"
#import "AgoraRtcEngineKit/AgoraRtcEngineKit.h"
#import "AgoraRTMManager.h"

#import "MicModel.h"

@interface ViewController ()<UITableViewDelegate, UITableViewDataSource,AgoraRtcEngineDelegate>

@property (nonatomic, strong)AgoraRtcEngineKit *AgoraRoom;
@property (nonatomic, strong)UILabel *label;
@property (nonatomic, strong)UIButton* btnSendMsg;
@property (nonatomic, strong)UIButton* btnExit;
@property (nonatomic, strong)UIButton* btnNew;

@property (nonatomic, strong)UITableView *tableView;

@property (nonatomic, strong)MicoView *micView;
@property (nonatomic, strong)MsgView *msgView;
@property (nonatomic, assign)NSInteger userid;
@property (nonatomic, copy)NSString *userName;

@property (nonatomic, strong)UITextField *textAccount;
@property (nonatomic, strong)UITextField *textPassword;

@property (nonatomic, strong)NSMutableDictionary *dicMic;

@property (nonatomic, strong)NSMutableArray *arrayMsg;

@property (nonatomic, strong)NSDictionary*roominfo;

@property (nonatomic, strong)NSArray *arrrayMicQueue;
@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    
    [TQNetWokTool instance];
    // Do any additional setup after loading the view, typically from a nib.
    _dicMic = [NSMutableDictionary dictionaryWithCapacity:10];
    _arrayMsg = [NSMutableArray arrayWithCapacity:10];
    [self.view addSubview:self.btnSendMsg];
    
    [self.view addSubview:self.btnExit];
    [self.view addSubview:self.micView];
    [self.view addSubview:self.msgView];
    [self.view addSubview:self.btnNew];
    [self.view addSubview:self.textAccount];
    [self.view addSubview:self.textPassword];
    
    [self.textAccount mas_makeConstraints:^(MASConstraintMaker *make) {
        make.left.equalTo(self.view);
        make.top.equalTo(self.btnSendMsg.mas_bottom);
        make.size.mas_offset(CGSizeMake(150, 50));
    }];
    
    [self.textPassword mas_makeConstraints:^(MASConstraintMaker *make) {
        make.right.equalTo(self.view);
        make.top.equalTo(self.btnSendMsg.mas_bottom);
        make.size.mas_offset(CGSizeMake(150, 50));
    }];
    
    [self.btnSendMsg mas_makeConstraints:^(MASConstraintMaker *make) {
        make.left.equalTo(self.view).mas_offset(0);
        make.top.equalTo(self.view).mas_offset(0);
        make.size.mas_offset(CGSizeMake(100, 50));
    }];
    
    [self.btnExit mas_makeConstraints:^(MASConstraintMaker *make) {
        make.centerX.equalTo(self.view);
        make.top.equalTo(self.view).mas_offset(0);
        make.size.mas_offset(CGSizeMake(100, 50));
    }];
    
    [self.btnNew mas_makeConstraints:^(MASConstraintMaker *make) {
        make.right.equalTo(self.view);
        make.top.equalTo(self.view).mas_offset(0);
        make.size.mas_offset(CGSizeMake(100, 50));
    }];
    
    [self.view addSubview:self.tableView];
    
    [self.tableView registerClass:[UITableViewCell class] forCellReuseIdentifier:@"UITableViewCell"];
    
    
    [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(updateUsersList:) name:CHANNEL_USERS_UPDATE object:nil];
    
    
    [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(updateMsgList:) name:CHANNEL_MSG_UPDATE object:nil];
    
    [[NSNotificationCenter defaultCenter] addObserver:self selector:@selector(updateMicsList:) name:CHANNEL_MICS_UPDATE object:nil];
    
    
    [self.AgoraRoom setChannelProfile:AgoraChannelProfileCommunication];

}

- (void)updateMicsList:(NSNotification*)notify{
    self.micView.arrayMic = [AgoraRTMManager instance].arrayMics;
}

- (void)runLoop{

    NSInteger index = [self getMyMicQueue];
    
    if (index >= 0) {
        
        NSString *url = @"/api/v1/mico/mic_heattime";
        
        NSMutableDictionary *dic = [NSMutableDictionary dictionary];
        [dic setObject:[self.roominfo objectForKey:@"RoomID"]  forKey:@"roomid"];
        [dic setObject:@(index) forKey:@"micindex"];
        
        [[TQNetWokTool instance] requestGet:url params:dic showLoading:NO success:^(id  _Nonnull responseObject, SuccessCode codeType) {
           
            if (codeType == SuccessCode_Success) {
                [AgoraRTMManager instance].arrayMics = [NSArray yy_modelArrayWithClass:[MicModel class] json:[responseObject objectForKey:@"data"]];
                [self.micView setArrayMic:[AgoraRTMManager instance].arrayMics];
            }
        } failure:^(NSError * _Nonnull error) {
            
        }];
    }
   
    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, (int64_t)(10 * NSEC_PER_SEC)), dispatch_get_main_queue(), ^{
        [self runLoop];
    });
}

- (NSInteger)getMyMicQueue{
    for (MicModel* obj in [AgoraRTMManager instance].arrayMics) {
        if (obj.PlayerID == [TQNetWokTool instance].account.PlayerID){
            return obj.MicIndex;
        }
    }
    return -1;
}


- (void)updateMsgList:(NSNotification*)notify{
    [self.msgView setArray:[AgoraRTMManager instance].currentChannelMsg];
}

- (void)updateUsersList:(NSNotification*)notify{
    [self.tableView reloadData];
}

- (void)doActionIn:(UIButton*)sender{
    
    if ([sender.titleLabel.text isEqualToString:@"进入"]){
        [sender setTitle:@"离开" forState:UIControlStateNormal];
        
        NSInteger roomID =20000;
        NSMutableDictionary *dic = [NSMutableDictionary dictionaryWithCapacity:10];
        [dic setObject:@(roomID) forKey:@"roomid"];
        
        NSString *url = @"/api/v1/room_manager/apply_enter_Room";
        [self.AgoraRoom setClientRole:AgoraClientRoleAudience];
        
        [[TQNetWokTool instance] requestPost:url params:dic showLoading:YES success:^(id  _Nonnull responseObject, SuccessCode codeType) {
            
            if (codeType != SuccessCode_Success) {
                [SVProgressHUD dismissWithCompletion:^{
                    [SVProgressHUD showInfoWithStatus:[responseObject objectForKey:@"msg"]];
                }];
                return;
            }
            
            NSString *agoratoken = [[responseObject objectForKey:@"data"] objectForKey:@"token"];
            self.roominfo=  [[responseObject objectForKey:@"data"] objectForKey:@"room"];
            
            NSArray *mics = [[responseObject objectForKey:@"data"] objectForKey:@"mics"];
            
            [AgoraRTMManager instance].arrayMics = [NSArray yy_modelArrayWithClass:[MicModel class] json:mics];
            NSArray*micqueue = [[responseObject objectForKey:@"data"] objectForKey:@"micqueue"];
           
            [self.micView setArrayMic:[AgoraRTMManager instance].arrayMics];
            
            [self.AgoraRoom joinChannelByToken:agoratoken channelId:[NSString stringWithFormat:@"channel%ld",(long)roomID]  info:nil uid:[TQNetWokTool instance].account.PlayerID joinSuccess:^(NSString * _Nonnull channel, NSUInteger uid, NSInteger elapsed) {
                
            }];
            
            [self runLoop];
            
        } failure:^(NSError * _Nonnull error) {
            
        }];
        
        [[AgoraRTMManager instance] enterChannel:[NSString stringWithFormat:@"channel%ld",(long)roomID] block:^(AgoraRtmJoinChannelErrorCode errorCode) {
            
            [[AgoraRTMManager instance] getChannetlMember:^(NSArray<AgoraRtmMember *> * _Nullable members, AgoraRtmGetMembersErrorCode errorCode) {
                
                [self.tableView reloadData];
            }];
        }];
        
    }
    else{
        [sender setTitle:@"进入" forState:UIControlStateNormal];
        
        [self.AgoraRoom leaveChannel:^(AgoraChannelStats * _Nonnull stat) {
           
        }];
        
        [[AgoraRTMManager instance] leaveChannel];
        
        [self.tableView reloadData];
        [self.dicMic removeAllObjects];
        [self.arrayMsg removeAllObjects];
        self.msgView.array = @[];
        [AgoraRTMManager instance].arrayMics = @[];
        [self.micView setArrayMic:[AgoraRTMManager instance].arrayMics];
    }
}

- (void)doActionSend:(UIButton*)sender{
    
    NSDictionary *dic = @{@"content":@"大家好",
                          @"type":@(1)
                          };
    [[AgoraRTMManager instance] sendChannelMsg:dic];
}

- (void)doActionNewMsg:(UIButton*)sender{
//    MsgViewController *pvc = [[MsgViewController alloc] init];
//    pvc.hidesBottomBarWhenPushed = YES;
//
//    [self presentModalViewController:pvc animated:YES];

    NSMutableDictionary *dic = [NSMutableDictionary dictionaryWithCapacity:10];
    
    [dic setObject:self.textAccount.text forKey:@"account"];
    [dic setObject:self.textPassword.text forKey:@"passowrd"];
    [SVProgressHUD showInfoWithStatus:@"登录中"];
    
    NSString *url = @"auth/login";
   
    
    [[TQNetWokTool instance] requestPost:url params:dic showLoading:YES success:^(id  _Nonnull responseObject, SuccessCode codeType) {
        
        if (codeType != SuccessCode_Success) {
            [SVProgressHUD dismissWithCompletion:^{
                [SVProgressHUD showInfoWithStatus:[responseObject objectForKey:@"msg"]];
            }];
            return;
        }
        [[TQNetWokTool instance] saveLoginInfo:responseObject];
        
        [SVProgressHUD dismissWithCompletion:^{
            [SVProgressHUD showInfoWithStatus:@"登录成功"];
        }];
        
    } failure:^(NSError * _Nonnull error) {
        [SVProgressHUD dismissWithCompletion:^{
            [SVProgressHUD showInfoWithStatus:@"登录失败"];
        }];
    }];

}

- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section{
    
    return [[AgoraRTMManager instance].currentChannelMember count];
    
}
- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath{
    UITableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:@"UITableViewCell" forIndexPath:indexPath];
    
    AgoraRtmMember *model = [AgoraRTMManager instance].currentChannelMember[indexPath.row];
    
    cell.textLabel.text = [NSString stringWithFormat:@"%@_%@",model.userId,model.channelId];
    return cell;
}


-(UILabel*)label{
    if (!_label) {
        _label = [[UILabel alloc] initWithFrame:CGRectMake(0, 0, 100, 20)];
        [_label setText:@"你好"];
    }
    return _label;
}

- (AgoraRtcEngineKit*)AgoraRoom{
    if (!_AgoraRoom) {
        _AgoraRoom = [AgoraRtcEngineKit sharedEngineWithAppId:@"1f836f0e094446d2858f156ca366313d" delegate:self];
    }
    return _AgoraRoom;
}

- (UIButton*)btnSendMsg{
    if (!_btnSendMsg) {
        _btnSendMsg = [[UIButton alloc] init];
        [_btnSendMsg setBackgroundColor:[UIColor grayColor]];
        [_btnSendMsg setTitle:@"发送" forState:UIControlStateNormal];
        [_btnSendMsg addTarget:self action:@selector(doActionSend:) forControlEvents:UIControlEventTouchUpInside];
        
    }
    return _btnSendMsg;
}

- (UIButton*)btnExit{
    if (!_btnExit) {
        _btnExit = [[UIButton alloc] init];
        [_btnExit setBackgroundColor:[UIColor grayColor]];
        [_btnExit setTitle:@"进入" forState:UIControlStateNormal];
        [_btnExit addTarget:self action:@selector(doActionIn:) forControlEvents:UIControlEventTouchUpInside];
    }
    return _btnExit;
}

- (UIButton*)btnNew{
    if (!_btnNew) {
        _btnNew = [[UIButton alloc] init];
        [_btnNew setBackgroundColor:[UIColor grayColor]];
        [_btnNew setTitle:@"登录" forState:UIControlStateNormal];
        [_btnNew addTarget:self action:@selector(doActionNewMsg:) forControlEvents:UIControlEventTouchUpInside];
    }
    return _btnNew;
}


- (UITableView*)tableView{
    if (!_tableView) {
        _tableView = [[UITableView alloc] initWithFrame:CGRectMake(0, 400, 200, 200) style:UITableViewStyleGrouped];
        _tableView.dataSource = self;
        _tableView.delegate= self;
        [_tableView setShowsVerticalScrollIndicator:NO];
        [_tableView setSeparatorStyle:UITableViewCellSeparatorStyleSingleLine];
        _tableView.tableFooterView = [[UIView alloc] initWithFrame:CGRectZero];
        [_tableView setKeyboardDismissMode:UIScrollViewKeyboardDismissModeOnDrag];
        _tableView.rowHeight = UITableViewAutomaticDimension;
        [_tableView setAllowsSelection:NO];
        if(@available(iOS 11.0, *)){
            _tableView.contentInsetAdjustmentBehavior = UIScrollViewContentInsetAdjustmentNever;
        }
    }
    return  _tableView;
}

- (MicoView*)micView{
    if (!_micView) {
        _micView = [[MicoView alloc] initWithFrame:CGRectMake(0, 130, 375, 200)];
        [_micView setBlockAction:^(NSInteger nindex) {
            
            NSString *url = @"/api/v1/mico/mico_apply_up";
            
            NSMutableDictionary *dic = [NSMutableDictionary dictionary];
            [dic setObject:[self.roominfo objectForKey:@"RoomID"] forKey:@"roomid"];
            [dic setObject:@(nindex) forKey:@"micindex"];
            NSInteger playerID = [self getmicIdByindex:nindex];
            
            if(0 == playerID){
                //没有用户，自己上麦
                url = @"/api/v1/mico/mico_apply_up";
            }
            else{
                if (playerID == [TQNetWokTool instance].account.PlayerID) {
                    //是自己,自己下麦
                     url = @"/api/v1/mico/mico_apply_down";
                }
                else{
                    //是别人，抱下麦
                    url = @"/api/v1/mico/mic_hold_down";
                    [dic setObject:@(playerID) forKey:@"tarPlayerID"];
                }
            }
            
            [[TQNetWokTool instance] requestPost:url params:dic showLoading:NO success:^(id  _Nonnull responseObject, SuccessCode codeType) {
                
                if (codeType == SuccessCode_Success) {
                    
                    id mics =[[responseObject objectForKey:@"data"] objectForKey:@"mics"];
                    [AgoraRTMManager instance].arrayMics = [NSArray yy_modelArrayWithClass:[MicModel class] json:mics];
                    [self.micView setArrayMic:[AgoraRTMManager instance].arrayMics];
                    
                    NSMutableDictionary *dic = [NSMutableDictionary dictionaryWithCapacity:4];
                    [dic setObject:@(100) forKey:@"type"];
                    [dic setObject:@(1) forKey:@"subtype"];
                    [dic setObject:mics forKey:@"mics"];
                    [[AgoraRTMManager instance] sendChannelMsg:dic];
                }
            } failure:^(NSError * _Nonnull error) {
                
            }];
            
        }];
        [_micView setBackgroundColor:[UIColor grayColor]];
        
        
    }
    return _micView;
}

- (UITextField*)textAccount{
    if (!_textAccount) {
        _textAccount = [[UITextField alloc] init];
        _textAccount.placeholder= @"账号";
    }
    return _textAccount;
}

- (UITextField*)textPassword{
    if (!_textPassword) {
        _textPassword = [[UITextField alloc] init];
        _textPassword.placeholder= @"密码";
    }
    return _textPassword;
}


- (NSInteger)getmicIdByindex:(NSInteger)nindex{
    
    for (MicModel* mic in [AgoraRTMManager instance].arrayMics) {
        if(mic.MicIndex == nindex){
            return mic.PlayerID;
        }
    }
    return 0;
}

- (MsgView*)msgView{
    if(!_msgView){
        _msgView = [[MsgView alloc] initWithFrame:CGRectMake(210, 350,375-205, 250)];
        [_msgView setBackgroundColor:[UIColor grayColor]];
    }
    return _msgView;
}


- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine didOccurWarning:(AgoraWarningCode)warningCode{
    
}

- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine didOccurError:(AgoraErrorCode)errorCode
{
    
}

- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine didApiCallExecute:(NSInteger)error api:(NSString * _Nonnull)api result:(NSString * _Nonnull)result{
    
}


- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine didJoinChannel:(NSString * _Nonnull)channel withUid:(NSUInteger)uid elapsed:(NSInteger) elapsed{
    
}


- (void)rtcEngine:(AgoraRtcEngineKit *_Nonnull)engine audioTransportStatsOfUid:(NSUInteger)uid delay:(NSUInteger)delay lost:(NSUInteger)lost rxKBitRate:(NSUInteger)rxKBitRate{
    NSLog(@"%lu_%lu_%lu_%lu",uid, delay, lost, rxKBitRate);
}
- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine reportAudioVolumeIndicationOfSpeakers:(NSArray<AgoraRtcAudioVolumeInfo *> * _Nonnull)speakers totalVolume:(NSInteger)totalVolume{
    
}

- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine activeSpeaker:(NSUInteger)speakerUid{
    
}

- (void)rtcEngine:(AgoraRtcEngineKit * _Nonnull)engine didJoinedOfUid:(NSUInteger)uid elapsed:(NSInteger)elapsed{
    
}
@end
