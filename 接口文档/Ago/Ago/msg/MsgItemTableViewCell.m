//
//  MsgItemTableViewCell.m
//  zego
//
//  Created by rd on 2019/8/13.
//  Copyright © 2019年 WebView. All rights reserved.
//

#import "MsgItemTableViewCell.h"
#import "Masonry.h"
@interface MsgItemTableViewCell()
@property (nonatomic, strong)UIImageView *headImage;
@property (nonatomic, strong)UILabel *nameLabel;
@property (nonatomic, strong)UILabel *tipsLabel;

@end

@implementation MsgItemTableViewCell

- (instancetype)initWithStyle:(UITableViewCellStyle)style reuseIdentifier:(NSString *)reuseIdentifier{
    if (self = [super initWithStyle:style reuseIdentifier:reuseIdentifier]) {
        
        [self.contentView addSubview:self.headImage];
        [self.contentView addSubview:self.nameLabel];
        [self.contentView addSubview:self.tipsLabel];
        
        [self.headImage mas_makeConstraints:^(MASConstraintMaker *make) {
            make.size.mas_offset(CGSizeMake(46, 46));
            make.left.equalTo(self.contentView).mas_offset(15);
            make.top.equalTo(self.contentView).mas_offset(8);
            make.bottom.equalTo(self.contentView).mas_offset(-8);
        }];
        
        [self.nameLabel mas_makeConstraints:^(MASConstraintMaker *make) {
            make.top.equalTo(self.headImage);
            make.left.equalTo(self.headImage.mas_right).mas_offset(15);
        }];
        
        [self.tipsLabel mas_makeConstraints:^(MASConstraintMaker *make) {
            make.left.equalTo(self.nameLabel);
            make.top.equalTo(self.nameLabel.mas_bottom).mas_offset(10);
        }];
        
    }
    return self;
}
- (void)setSelected:(BOOL)selected animated:(BOOL)animated {
    [super setSelected:selected animated:animated];

    // Configure the view for the selected state
}

- (UIImageView*)headImage{
    if (!_headImage) {
        _headImage = [[UIImageView alloc] init];
        [_headImage setBackgroundColor:[UIColor redColor]];
    }
    return _headImage;
}

- (UILabel*)nameLabel{
    if (!_nameLabel) {
        _nameLabel = [[UILabel alloc] init];
        _nameLabel.textColor = [UIColor blackColor];
        _nameLabel.text = @"名字名字名字";
        _nameLabel.font = [UIFont systemFontOfSize:15];
    }
    return _nameLabel;
}

- (UILabel*)tipsLabel{
    if (!_tipsLabel) {
        _tipsLabel = [[UILabel alloc] init];
        _tipsLabel.textColor = [UIColor grayColor];
        _tipsLabel.text = @"标签标签标签表亲";
        _tipsLabel.font = [UIFont systemFontOfSize:13];
    }
    return _tipsLabel;
}
@end
