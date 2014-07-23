//
//  Game.m
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Game.h"
#import "ViewController.h"


@implementation Game

- (instancetype)initWithController:(ViewController*)controller {
    self.controller = controller;
    self.character = [[Character alloc] init];
    return self;
}

- (void)render {
    self.controller.healthLabel.text = [@(self.character.health) stringValue];
    self.controller.damageLabel.text = [@(self.character.damage) stringValue];
    self.controller.weaponLabel.text = self.character.weapon.name;
    self.controller.armorLabel.text = self.character.armor.name;
}

@end
