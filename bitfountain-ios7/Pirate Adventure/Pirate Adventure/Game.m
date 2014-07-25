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
    NSMutableArray* tiles = [@[
             [[Tile alloc] initWithStory:@"Captain, we need a fearless leader such as yourself to undertake a voyage. You must stop the evil pirate Roberts. Would you like a sword to get started?"
                              background:[UIImage imageNamed:@"PirateStart.jpg"]
                              actionName:@"Take the sword"
                                  action:^(Tile* tile) {
                                      self.character.weapon = [[Weapon alloc] initWithName:@"Blunted Sword" damage:9];
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"You have come across an armory. Would you like to upgrade your armor to steel armor?"
                              background:[UIImage imageNamed:@"PirateBlacksmith.jpeg"]
                              actionName:@"Take the armor"
                                  action:^(Tile* tile) {
                                      self.character.armor = [[Armor alloc] initWithName:@"Steel armor" value:9];
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"Your final faceoff with the fearsome pirate bos"
                              background:[UIImage imageNamed:@"PirateBoss.jpeg"]
                              actionName:@"Fight"
                                  action:^(Tile* tile) {}
                                    lock:YES],
             [[Tile alloc] initWithStory:@"A mysterious dock appears on the horizon. Should we stop and ask for directions?"
                              background:[UIImage imageNamed:@"PirateFriendlyDock.jpg"]
                              actionName:@"Ask for directions."
                                  action:^(Tile* tile) {
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"A group of pirates attempts to board your ship"
                              background:[UIImage imageNamed:@"PirateAttack.jpg"]
                              actionName:@"Don't let them"
                                  action:^(Tile* tile) {
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"You have found a parrot can be used in your armor slot! Parrots are a great defender and are fiercly loyal to their captain."
                              background:[UIImage imageNamed:@"PirateParrot.jpg"]
                              actionName:@"Take the parrot"
                                  action:^(Tile* tile) {
                                      self.character.armor = [[Armor alloc] initWithName:@"Parrot" value:15];
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"You have been captured by pirates and are ordered to walk the plank"
                              background:[UIImage imageNamed:@"PiratePlank.jpg"]
                              actionName:@"Walk the plank"
                                  action:^(Tile* tile) {
                                      self.character.health = self.character.health - 10;
                                      tile.actionState = NO;
                                      tile.lock = NO;
                                  }
                                    lock:YES],
             [[Tile alloc] initWithStory:@"You sight a pirate battle off the coast"
                              background:[UIImage imageNamed:@"PirateShipBattle.jpeg"]
                              actionName:@"Do a little jig"
                                  action:^(Tile* tile) {
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"The legend of the deep, the mighty kraken appears"
                              background:[UIImage imageNamed:@"PirateOctopusAttack.jpg"]
                              actionName:@"OMG!"
                                  action:^(Tile* tile) {
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"You stumble upon a hidden cave of pirate treasure."
                              background:[UIImage imageNamed:@"PirateTreasure.jpeg"]
                              actionName:@"Awesome!"
                                  action:^(Tile* tile) {
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"In the deep of the sea many things live and many things can be found. Will the fortune bring help or ruin?"
                              background:[UIImage imageNamed:@"PirateTreasurer2.jpeg"]
                              actionName:@"Jawsome!"
                                  action:^(Tile* tile) {
                                      tile.actionState = NO;
                                  }],
             [[Tile alloc] initWithStory:@"You have stumbled upon a cache of pirate weapons. Would you like to upgrade to a pistol?"
                              background:[UIImage imageNamed:@"PirateWeapons.jpeg"]
                              actionName:@"Take the pistol"
                                  action:^(Tile* tile) {
                                      self.character.weapon = [[Weapon alloc] initWithName:@"Pistol" damage:15];
                                      tile.actionState = NO;
                                  }],
    ] mutableCopy];
    for (int i = 1; i < [tiles count]; i++) {
        [tiles exchangeObjectAtIndex:i withObjectAtIndex:1 + (arc4random() % ([tiles count] - 1))];
    }
    return [tiles copy];
}

- (Tile*)currentTile {
    return self.tiles[(int)self.currentLocation.x * 3 + (int)self.currentLocation.y];
}

- (void)render {
    // character information
    self.controller.healthLabel.text = [@(self.character.health) stringValue];
    self.controller.damageLabel.text = [@(self.character.damage) stringValue];
    self.controller.weaponLabel.text = [self.character.weapon displayString];
    self.controller.armorLabel.text = [self.character.armor displayString];

    // tile information
    [self.controller.actionButton setTitle:[self currentTile].actionButtonName forState:UIControlStateNormal];
    self.controller.storyLabel.text = [self currentTile].story;
    self.controller.backgroundImageView.image = [self currentTile].backgroundImage;
    self.controller.northButton.hidden = [self currentTile].lock || self.currentLocation.y == 2;
    self.controller.southButton.hidden = [self currentTile].lock || self.currentLocation.y == 0;
    self.controller.eastButton.hidden = [self currentTile].lock || self.currentLocation.x == 3;
    self.controller.westButton.hidden = [self currentTile].lock || self.currentLocation.x == 0;
    self.controller.actionButton.hidden = ![self currentTile].actionState;
}

@end
