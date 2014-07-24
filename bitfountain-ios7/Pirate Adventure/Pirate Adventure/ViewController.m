//
//  ViewController.m
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"
#import "Game.h"


@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    self.game = [[Game alloc] initWithController:self];
    [self.game render];
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (IBAction)actionButtonPressed:(UIButton *)sender {
    [self.game currentTile].action();
}

- (IBAction)northButtonPressed:(UIButton *)sender {
    self.game.currentLocation = CGPointMake(self.game.currentLocation.x,
                                            self.game.currentLocation.y + 1);
    [self.game render];
}

- (IBAction)southButtonPressed:(UIButton *)sender {
    self.game.currentLocation = CGPointMake(self.game.currentLocation.x,
                                            self.game.currentLocation.y - 1);
    [self.game render];
}

- (IBAction)eastButtonPressed:(UIButton *)sender {
    self.game.currentLocation = CGPointMake(self.game.currentLocation.x + 1,
                                            self.game.currentLocation.y);
    [self.game render];
}

- (IBAction)westButtonPressed:(UIButton *)sender {
    self.game.currentLocation = CGPointMake(self.game.currentLocation.x - 1,
                                            self.game.currentLocation.y);
    [self.game render];
}
@end
