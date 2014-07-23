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
    self.tiles = [self createTiles];
    return self;
}

- (NSArray*)createTiles {
    return @[
        [[Tile alloc] initWithStory:@"story0"
                         background:nil
                         actionName:@"action0"
                             action:^{ NSLog(@"action0"); }],
    ];
}

- (Tile*)currentTile {
    return self.tiles[0];
}

- (void)render {
    self.controller.healthLabel.text = [@(self.character.health) stringValue];
    self.controller.damageLabel.text = [@(self.character.damage) stringValue];
    self.controller.weaponLabel.text = self.character.weapon.name;
    self.controller.armorLabel.text = self.character.armor.name;
    [self.controller.actionButton setTitle:[self currentTile].actionButtonName forState:UIControlStateNormal];
    self.controller.storyLabel.text = [self currentTile].story;
}

@end
