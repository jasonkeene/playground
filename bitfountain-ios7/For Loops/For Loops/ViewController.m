//
//  ViewController.m
//  For Loops
//
//  Created by bonobo on 7/10/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"
#include <math.h>

@interface ViewController ()
            

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    unsigned long long count = 1;
    for (int i = 2; i <= 64; i++) {
        NSLog(@"%llu", count);
        count *= 2;
    }
    NSLog(@"%llu", count);
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
