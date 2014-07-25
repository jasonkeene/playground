//
//  Character.h
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <Foundation/Foundation.h>

#import "Armor.h"
#import "Weapon.h"


@interface Character : NSObject {
    int _damage;
}

@property (nonatomic) int health;
@property (strong, nonatomic) Armor* armor;
@property (strong, nonatomic) Weapon* weapon;

- (int)damage;
- (void)setDamage:(int)value;

@end
