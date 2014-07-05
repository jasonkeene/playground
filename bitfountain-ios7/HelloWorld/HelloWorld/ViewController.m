//
//  ViewController.m
//  HelloWorld
//
//  Created by bonobo on 7/5/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"

@interface ViewController ()
            

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (IBAction)buttonPressed:(UIButton *)sender {
    self.titleLabel.text = self.textField.text.length > 0 ? self.textField.text : @"Default";
    [self.textField resignFirstResponder];
}
@end
