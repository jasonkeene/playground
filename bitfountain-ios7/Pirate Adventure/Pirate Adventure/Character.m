//
//  Character.m
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Character.h"


@implementation Character

- (instancetype)init {
    self = [super init];
    self.health = 0;
    self.damage = 0;
    self.armor = [[Armor alloc] init];
    self.weapon = [[Weapon alloc] init];
    return self;
}

@end
