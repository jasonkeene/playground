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


@interface Character : NSObject

@property (nonatomic) int health;
@property (nonatomic) int damage;
@property (strong, nonatomic) Armor* armor;
@property (strong, nonatomic) Weapon* weapon;

@end
