//
//  ViewController.m
//  Unit Converter
//
//  Created by bonobo on 7/7/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"

@interface ViewController ()
            

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    NSLog(@"Testing out breakpoints.");
    int x = 12;
    float y = 12.1212;
    NSLog(@"I can haz int? %d", x);
    NSLog(@"I can haz float? %f", y);
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (IBAction)convertUnits:(UIButton *)sender {
}
@end
