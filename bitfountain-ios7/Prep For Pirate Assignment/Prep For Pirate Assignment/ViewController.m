//
//  ViewController.m
//  Prep For Pirate Assignment
//
//  Created by bonobo on 7/19/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"

@interface ViewController ()
            

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    NSString* firstString = @"First String";
    NSString* secondString = @"Second String";

    NSMutableArray* myMutableArray = [[NSMutableArray alloc] init];
    [myMutableArray addObject:firstString];
    [myMutableArray addObject:secondString];
    
    NSArray* myArray = [[NSArray alloc] initWithObjects:firstString, myMutableArray, secondString, nil];

    self.currentPoint = CGPointMake(3, 4);

    int x = 10;
    int y = 20;

    if (x == 10) {
        NSLog(@"x equals 10");
        if (y == 20) {
            NSLog(@"both are true");
        }
    }

    [self.button setTitle:@"Button Pressed" forState:UIControlStateNormal];
    self.button.hidden = NO;

    UIAlertView* alertView = [[UIAlertView alloc] initWithTitle:@"Alert" message:@"You triggered the alert!" delegate:nil cancelButtonTitle:@"Cancel" otherButtonTitles:nil];
    [alertView show];

    self.awesome = [[AwesomeClass alloc] init];

    NSString* testString = @"Has a value";
    if (testString != nil) {
        NSLog(@"The test string has a value!");
    }
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
