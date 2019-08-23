//
//  MicoView.m
//  zego
//
//  Created by rd on 2019/8/12.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "MicoView.h"
#import "MicModel.h"

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
        
        
        _btnHeadMic = [[UIButton alloc] initWithFrame:CGRectMake(100, 0, 100, 40)];
        [self addSubview:_btnHeadMic];
        [_btnHeadMic setTitle:@"上麦" forState:UIControlStateNormal];
        [_btnHeadMic addTarget:self action:@selector(doActionHead:) forControlEvents:UIControlEventTouchUpInside];
        
        for (int i = 0; i < 8; i++) {
            UIButton *btn = nil;
            
            if (i <4) {
                btn = [[UIButton alloc] initWithFrame:CGRectMake(20+i*95, 50, 70, 40)];
            }
            else{
                btn = [[UIButton alloc] initWithFrame:CGRectMake(20+(i-4)*95, 100, 70, 40)];
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

- (void)setArrayMic:(NSArray *)arrayMic{
    
    _arrayMic = arrayMic;
    [self.btnHeadMic setTitle:@"上" forState:UIControlStateNormal];
    
    for (id obj in self.dicMic) {
        UIButton *btn  = [self.dicMic objectForKey:obj];
        [btn setTitle:@"上" forState:UIControlStateNormal];
    }
    for (MicModel* obj in _arrayMic) {
        
        UIButton *btn = [self.dicMic objectForKey:@(obj.MicIndex)];
        if (btn != nil) {
            [btn setTitle:[NSString stringWithFormat:@"%lu",obj.PlayerID] forState:UIControlStateNormal];
        }
    }
}

@end
