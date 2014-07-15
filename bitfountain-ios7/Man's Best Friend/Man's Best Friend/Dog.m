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

@end
