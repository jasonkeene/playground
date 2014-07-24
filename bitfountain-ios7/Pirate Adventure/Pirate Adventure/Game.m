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
    self.currentLocation = CGPointMake(0, 0);
    return self;
}

- (NSArray*)createTiles {
    return @[
             [[Tile alloc] initWithStory:@"story0"
                              background:nil
                              actionName:@"action0"
                                  action:^{ NSLog(@"action0"); }],
             [[Tile alloc] initWithStory:@"story1"
                              background:nil
                              actionName:@"action1"
                                  action:^{ NSLog(@"action1"); }],
             [[Tile alloc] initWithStory:@"story2"
                              background:nil
                              actionName:@"action2"
                                  action:^{ NSLog(@"action2"); }],
             [[Tile alloc] initWithStory:@"story3"
                              background:nil
                              actionName:@"action3"
                                  action:^{ NSLog(@"action3"); }],
             [[Tile alloc] initWithStory:@"story4"
                              background:nil
                              actionName:@"action4"
                                  action:^{ NSLog(@"action4"); }],
             [[Tile alloc] initWithStory:@"story5"
                              background:nil
                              actionName:@"action5"
                                  action:^{ NSLog(@"action5"); }],
             [[Tile alloc] initWithStory:@"story6"
                              background:nil
                              actionName:@"action6"
                                  action:^{ NSLog(@"action6"); }],
             [[Tile alloc] initWithStory:@"story7"
                              background:nil
                              actionName:@"action7"
                                  action:^{ NSLog(@"action7"); }],
             [[Tile alloc] initWithStory:@"story8"
                              background:nil
                              actionName:@"action8"
                                  action:^{ NSLog(@"action8"); }],
             [[Tile alloc] initWithStory:@"story9"
                              background:nil
                              actionName:@"action9"
                                  action:^{ NSLog(@"action9"); }],
             [[Tile alloc] initWithStory:@"story10"
                              background:nil
                              actionName:@"action10"
                                  action:^{ NSLog(@"action10"); }],
             [[Tile alloc] initWithStory:@"story11"
                              background:nil
                              actionName:@"action11"
                                  action:^{ NSLog(@"action11"); }],
    ];
}

- (Tile*)currentTile {
    return self.tiles[(int)self.currentLocation.x * 3 + (int)self.currentLocation.y];
}

- (void)render {
    self.controller.healthLabel.text = [@(self.character.health) stringValue];
    self.controller.damageLabel.text = [@(self.character.damage) stringValue];
    self.controller.weaponLabel.text = self.character.weapon.name;
    self.controller.armorLabel.text = self.character.armor.name;
    [self.controller.actionButton setTitle:[self currentTile].actionButtonName forState:UIControlStateNormal];
    self.controller.storyLabel.text = [self currentTile].story;
    self.controller.northButton.hidden = self.currentLocation.y == 2;
    self.controller.southButton.hidden = self.currentLocation.y == 0;
    self.controller.eastButton.hidden = self.currentLocation.x == 3;
    self.controller.westButton.hidden = self.currentLocation.x == 0;
}

@end
