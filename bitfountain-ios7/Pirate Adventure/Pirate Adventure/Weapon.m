//
//  Weapon.m
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Weapon.h"


@implementation Weapon

- (instancetype)init {
    self = [super init];
    self.name = @"Starter Weapon";
    self.damage = 1;
    return self;
}

@end
