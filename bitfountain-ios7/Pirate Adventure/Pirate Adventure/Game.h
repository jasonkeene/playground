//
//  Game.h
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>


@interface Game : NSObject

@property (strong, nonatomic) UIViewController* controller;
@property (nonatomic) CGPoint currentLocation;

- (instancetype)initWithController:(UIViewController*)controller;

@end
