//
//  MicModel.h
//  Ago
//
//  Created by rd on 2019/8/23.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import <Foundation/Foundation.h>

NS_ASSUME_NONNULL_BEGIN
/*
 CreatedAt = "2019-08-23T19:16:19+08:00";
 DeletedAt = "<null>";
 ID = 1;
 MicIndex = 1;
 MicStatus = 0;
 PlayerID = 20000;
 RoomID = 20000;
 Status = 0;
 UpdatedAt = "2019-08-23T19:16:19+08:00"
 */
@interface MicModel : NSObject

@property (nonatomic, assign)NSInteger PlayerID;
@property (nonatomic, assign)NSInteger MicIndex;
@property (nonatomic, assign)NSInteger MicStatus;
@property (nonatomic, assign)NSInteger RoomID;
@end

NS_ASSUME_NONNULL_END
