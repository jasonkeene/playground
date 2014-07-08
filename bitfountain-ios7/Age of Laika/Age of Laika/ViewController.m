//
//  ViewController.m
//  Age of Laika
//
//  Created by bonobo on 7/7/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"


@interface ViewController ()
@end


float calculate_age(float years) {
    return years * 7;
}


@implementation ViewController

- (IBAction)calculateAge:(UIButton *)sender {
    float result = [self.input.text floatValue];
    self.output.text = [
        NSString stringWithFormat:@"%.2f years",
        calculate_age(result)
    ];
    // hide keyboard
    [self.input resignFirstResponder];
}

@end
