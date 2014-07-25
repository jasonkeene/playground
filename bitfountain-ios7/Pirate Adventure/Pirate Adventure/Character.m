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
    self.health = 100;
    self.damage = 5;
    self.armor = [[Armor alloc] init];
    self.weapon = [[Weapon alloc] init];
    return self;
}

- (int)damage {
    return _damage + self.weapon.damage;
}

- (void)setDamage:(int)value {
    _damage = value;
}

@end
