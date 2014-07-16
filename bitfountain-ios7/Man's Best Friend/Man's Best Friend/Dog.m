//
//  Dog.m
//  Man's Best Friend
//
//  Created by bonobo on 7/12/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Dog.h"


@implementation Dog

- (instancetype)init {
    self = [super init];
    self.name = @"Rando dog!";
    self.breed = @"Mutt";
    self.age = 5;
    return self;
}

- (void)bark {
    NSLog(@"Woof Woof!");
}

- (void)barkTimes:(int)times {
    for (int i = 0; i < times; i++) {
        [self bark];
    }
}

- (void)barkTimes:(int)times loudly:(BOOL)loud {
    if (loud) {
        for (int i = 0; i < times; i++) {
            NSLog(@"Ruff Ruff!");
        }
    } else {
        for (int i = 0; i < times; i++) {
            NSLog(@"yip yip!");
        }
    }
}

- (void)lycanize {
    self.breed = @"Lycan";
}

@end
