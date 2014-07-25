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
    return [self initWithName:@"Fists" damage:5];
}

- (instancetype)initWithName:(NSString*)name damage:(int)damage {
    self = [super init];
    self.name = name;
    self.damage = damage;
    return self;
}

- (NSString*)displayString {
    if (self.damage > 0) {
        return [NSString stringWithFormat:@"%@ (+%d)", self.name,
                self.damage, nil];
    } else {
        return [NSString stringWithFormat:@"%@", self.name, nil];
    }
}

@end
