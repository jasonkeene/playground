//
//  Tile.m
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Tile.h"


@implementation Tile

- (instancetype)initWithStory:(NSString*)story
                   background:(UIImage*)background
                   actionName:(NSString*)actionName
                       action:(void (^)(void))action {
    self = [super init];
    self.story = story;
    self.backgroundImage = background;
    self.actionButtonName = actionName;
    self.action = action;
    self.actionState = YES;
    return self;
}

@end
