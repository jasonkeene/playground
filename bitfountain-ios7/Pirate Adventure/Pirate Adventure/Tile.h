//
//  Tile.h
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>


@interface Tile : NSObject

@property (strong, nonatomic) NSString* story;
@property (strong, nonatomic) UIImage *backgroundImage;
@property (strong, nonatomic) NSString *actionButtonName;
@property (strong, nonatomic) void (^action)(void);

- (instancetype)initWithStory:(NSString*)story
                   background:(UIImage*)background
                   actionName:(NSString*)actionName
                       action:(void (^)(void))action;

@end
