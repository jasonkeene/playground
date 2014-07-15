//
//  ViewController.m
//  Man's Best Friend
//
//  Created by bonobo on 7/11/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"
#import "Dog.h"


@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    Dog* dog = [[Dog alloc] init];
    [dog barkTimes:2];
    NSLog(@"%@", dog.breed);
    [dog lycanize];
    NSLog(@"%@", dog.breed);
    [self printHelloWorld];
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (void)printHelloWorld {
    NSLog(@"Hello world!");
}

@end
