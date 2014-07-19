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

    NSLog(@"%@", myArray);

    self.currentPoint = CGPointMake(3, 4);
    NSLog(@"%f %f", self.currentPoint.x, self.currentPoint.y);
    NSLog(@"%@", NSStringFromCGPoint(self.currentPoint));
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
