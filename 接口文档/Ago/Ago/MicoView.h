//
//  MicoView.h
//  zego
//
//  Created by rd on 2019/8/12.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import <UIKit/UIKit.h>

NS_ASSUME_NONNULL_BEGIN

@interface MicoView : UIView

@property(nonatomic,copy)void(^blockAction)(NSInteger    nindex);

- (void)setDicmic:(NSMutableDictionary *)dicmic;

@end

NS_ASSUME_NONNULL_END
