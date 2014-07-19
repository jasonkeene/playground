//
//  ViewController.m
//  Class Testing Ground
//
//  Created by bonobo on 7/18/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"

@interface ViewController ()
            

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    
    NSString* sentence = @"The NewFoundland dog breed has webbed feet which aids in its swimming prowess";
    NSArray* words = [sentence componentsSeparatedByString:@" "];
    NSLog(@"%@", words);
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
