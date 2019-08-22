//
//  MsgTableView.m
//  zego
//
//  Created by rd on 2019/8/13.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "MsgTableView.h"
#import "SVProgressHUD.h"
#import "Masonry.h"
#import "MsgItemTableViewCell.h"

@interface MsgTableView()<UITableViewDelegate, UITableViewDataSource>

@property (nonatomic, strong)UITableView *tableView;
@end

@implementation MsgTableView

- (instancetype)initWithFrame:(CGRect)frame
{
    self = [super initWithFrame:frame];
    if (self) {
        
        [self addSubview:self.tableView];
        
        [self.tableView mas_makeConstraints:^(MASConstraintMaker *make) {
            make.edges.equalTo(self);
        }];
        
        [self.tableView registerClass:[MsgItemTableViewCell class] forCellReuseIdentifier:@"MsgItemTableViewCell"];
    }
    return self;
}

- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section{
    return 10;
}
- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath{
    
    MsgItemTableViewCell *cell = [tableView dequeueReusableCellWithIdentifier:@"MsgItemTableViewCell" forIndexPath:indexPath];
    
    return cell;
}

- (UITableView*)tableView{
    if (!_tableView) {
        _tableView = [[UITableView alloc] initWithFrame:CGRectZero style:UITableViewStyleGrouped];
        _tableView.dataSource = self;
        _tableView.delegate = self;
        [_tableView setBackgroundColor:[UIColor whiteColor]];
        [_tableView setShowsVerticalScrollIndicator:NO];
        [_tableView setSeparatorStyle:UITableViewCellSeparatorStyleNone];
        [_tableView setKeyboardDismissMode:UIScrollViewKeyboardDismissModeOnDrag];
        _tableView.rowHeight = UITableViewAutomaticDimension;
        [_tableView setAllowsSelection:NO];
        if(@available(iOS 11.0, *)){
            _tableView.contentInsetAdjustmentBehavior = UIScrollViewContentInsetAdjustmentNever;
        }
       
    }
    return _tableView;
}

@end
