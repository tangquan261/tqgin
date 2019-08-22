//
//  MsgViewController.m
//  zego
//
//  Created by rd on 2019/8/13.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "MsgViewController.h"
#import "SVProgressHUD.h"
#import "Masonry.h"
#import "MsgTableView.h"

@interface MsgViewController ()
@property (nonatomic, strong)UIScrollView *msgScrollview;
@property (nonatomic, strong)NSMutableArray *arrayTabelView;
@end

@implementation MsgViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view.
    [self.view setBackgroundColor:[UIColor whiteColor]];
    
    self.automaticallyAdjustsScrollViewInsets = NO;
    self.edgesForExtendedLayout = UIRectEdgeNone;
    NSArray *arraryTitle = @[@"消息",@"好友",@"关注",@"粉丝"];
    
    for (int i = 0; i < 4; i++) {
        UIButton *btn = [[UIButton alloc] initWithFrame:CGRectMake(50+ i*80, 20, 40, 80)];
        [self.view addSubview:btn];
        [btn setTitle:arraryTitle[i] forState:UIControlStateNormal];
        btn.tag = i;
        [btn setTitleColor:[UIColor blackColor] forState:UIControlStateNormal];
        [btn addTarget:self action:@selector(doActionSelect:) forControlEvents:UIControlEventTouchUpInside];
    }
    
    [self.view addSubview:self.msgScrollview];
    
    UIButton*btn = [[UIButton alloc] init];
    [btn setTitle:@"返回" forState:UIControlStateNormal];
    [btn setTitleColor:[UIColor blackColor] forState:UIControlStateNormal];
    btn.titleLabel.font = [UIFont systemFontOfSize:15];
    [btn addTarget:self action:@selector(doActionBack:) forControlEvents:UIControlEventTouchUpInside];
    [self.view addSubview:btn];
    [btn setBackgroundColor:[UIColor redColor]];
    
    [btn mas_makeConstraints:^(MASConstraintMaker *make) {
        make.left.top.equalTo(self.view);
        make.size.mas_offset(CGSizeMake(100, 100));
    }];
    
    
    
    [self.msgScrollview mas_makeConstraints:^(MASConstraintMaker *make) {
        make.left.right.bottom.equalTo(self.view);
        make.top.equalTo(self.view).mas_offset(120);
    }];
    
    _arrayTabelView = [NSMutableArray arrayWithCapacity:4];
   
    for (int i = 0; i < 4; i++) {
         MsgTableView *tableview = [[MsgTableView alloc] initWithFrame:CGRectMake(i*375, 0, 375, 667)];
        [self.msgScrollview addSubview:tableview];
    }
}
- (void)doActionBack:(UIButton*)sender{
    [self dismissViewControllerAnimated:YES completion:^{
        
    }];
}

- (void)scrollViewDidScroll:(UIScrollView *)scrollView {
    
   // NSInteger currentPage = roundf(scrollView.contentOffset.x/scrollView.size.width);
    
   
}

- (void)doActionSelect:(UIButton*)sender{
    [SVProgressHUD showInfoWithStatus:[NSString stringWithFormat:@"%u",sender.tag]];

    [self.msgScrollview setContentOffset:CGPointMake(sender.tag * 375, 0) animated:YES];
}

- (UIScrollView*)msgScrollview{
    if (!_msgScrollview) {
        _msgScrollview = [[UIScrollView alloc] init];
        _msgScrollview.pagingEnabled = YES;
        [_msgScrollview setShowsHorizontalScrollIndicator:NO];
        [_msgScrollview setShowsVerticalScrollIndicator:NO];
        _msgScrollview.bounces = NO;
        _msgScrollview.directionalLockEnabled=YES;
        [_msgScrollview setContentSize:CGSizeMake(375*4, 667)];
        
        if (@available(iOS 11.0, *)) {
            _msgScrollview.contentInsetAdjustmentBehavior = UIScrollViewContentInsetAdjustmentNever;
        }
    }
    return _msgScrollview;
}

@end
