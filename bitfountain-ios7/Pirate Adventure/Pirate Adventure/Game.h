//
//  Game.h
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "Character.h"
#import "Tile.h"

@class ViewController;


@interface Game : NSObject

@property (strong, nonatomic) ViewController* controller;
@property (strong, nonatomic) Character* character;
@property (strong, nonatomic) NSArray* tiles;
@property (nonatomic) CGPoint currentLocation;

- (instancetype)initWithController:(UIViewController*)controller;
- (NSArray*)createTiles;
- (Tile*)currentTile;
- (void)render;

@end
