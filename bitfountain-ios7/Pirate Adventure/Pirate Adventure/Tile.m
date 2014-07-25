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
    return [self initWithStory:story
                    background:background
                    actionName:actionName
                        action:action
                          lock:NO];
}

- (instancetype)initWithStory:(NSString*)story
                   background:(UIImage*)background
                   actionName:(NSString*)actionName
                       action:(void (^)(void))action
                         lock:(BOOL)lock {
    self = [super init];
    self.story = story;
    self.backgroundImage = background;
    self.actionButtonName = actionName;
    self.action = action;
    self.lock = lock;
    self.actionState = YES;
    return self;
}

@end
