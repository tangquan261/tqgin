//
//  MicoView.m
//  zego
//
//  Created by rd on 2019/8/12.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "MicoView.h"

@interface MicoView()

@property (nonatomic, strong)UIButton* btnHeadMic;

@property (nonatomic, strong)NSMutableDictionary *dicMic;

@end

@implementation MicoView

- (instancetype)initWithFrame:(CGRect)frame
{
    self = [super initWithFrame:frame];
    if (self) {
        _dicMic = [NSMutableDictionary dictionaryWithCapacity:8];
        
        
        _btnHeadMic = [[UIButton alloc] initWithFrame:CGRectMake(100, 0, 100, 30)];
        [self addSubview:_btnHeadMic];
        [_btnHeadMic setTitle:@"上麦" forState:UIControlStateNormal];
        [_btnHeadMic addTarget:self action:@selector(doActionHead:) forControlEvents:UIControlEventTouchUpInside];
        
        for (int i = 0; i < 8; i++) {
            UIButton *btn = nil;
            
            if (i <4) {
                btn = [[UIButton alloc] initWithFrame:CGRectMake(20+i*75, 50, 50, 30)];
            }
            else{
                btn = [[UIButton alloc] initWithFrame:CGRectMake(20+(i-4)*75, 100, 50, 30)];
            }
            [btn setBackgroundColor:[UIColor yellowColor]];
            [btn setTitleColor:[UIColor blackColor] forState:UIControlStateNormal];
            [self addSubview:btn];
            
            [_dicMic setObject:btn forKey:@(i)];
            
            [btn setTag:i];
            [btn setTitle:[NSString stringWithFormat:@"%u",i] forState:UIControlStateNormal];
            [btn addTarget:self action:@selector(doActionMic:) forControlEvents:UIControlEventTouchUpInside];
        }
    
    }
    return self;
}

- (void)doActionHead:(UIButton*)sender{
    if (self.blockAction) {
        self.blockAction(100);
    }
}

- (void)doActionMic:(UIButton*)sender{
    if (self.blockAction) {
        self.blockAction(sender.tag);
    }
}

- (void)setDicmic:(NSMutableDictionary *)dicmic{
    [self.btnHeadMic setTitle:@"上" forState:UIControlStateNormal];
    
    for (id btnnindex in _dicMic) {
        UIButton *btn  = [_dicMic objectForKey:btnnindex];
        [btn setTitle:@"上" forState:UIControlStateNormal];
    }
    
    for (id obj in dicmic) {
//        ZegoAudioStream *stream = [dicmic objectForKey:obj];
//        if ([stream.extraInfo integerValue] == 100) {
//            [self.btnHeadMic setTitle:stream.userName forState:UIControlStateNormal];
//        }
//        else{
//            UIButton *btn = [self.dicMic objectForKey:@([stream.extraInfo integerValue])];
//            if (btn != nil) {
//                [btn setTitle:stream.userName forState:UIControlStateNormal];
//            }
//        }
    }
}


@end
