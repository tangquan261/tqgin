//
//  MsgView.m
//  zego
//
//  Created by rd on 2019/8/12.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "MsgView.h"
#import "Masonry.h"
#import <AgoraRtmKit/AgoraRtmKit.h>
#import "NSStringUtil.h"
@interface MsgView()<UITableViewDelegate, UITableViewDataSource>

@property (nonatomic, strong)UITableView *tableView;

@end

@implementation MsgView

- (instancetype)initWithFrame:(CGRect)frame
{
    self = [super initWithFrame:frame];
    if (self) {
        
        [self addSubview:self.tableView];
        
        [self.tableView mas_makeConstraints:^(MASConstraintMaker *make) {
            make.edges.equalTo(self);
        }];
         
        [self.tableView registerClass:[UITableViewCell class] forCellReuseIdentifier:@"UITableViewCell"];
        
    }
    return self;
}

- (void)setArray:(NSMutableArray *)array{
    _array = array;
    [self.tableView reloadData];
}
- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section{
 
    return [_array count];
}

- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath{
    UITableViewCell *cell= [tableView dequeueReusableCellWithIdentifier:@"UITableViewCell" forIndexPath:indexPath];
    
    AgoraRtmMessage *msg = _array[indexPath.row];
    
    cell.textLabel.text = [NSString stringWithFormat:@"%@_%lu",msg.text, msg.type];
    
    NSDictionary *dic = [NSStringUtil dictionaryWithJsonString:msg.text];
    NSLog(@"%@",dic);
    
    return cell;
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
@end
