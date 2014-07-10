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

float calculate_real_age(float years) {
    float dog_years = 0;
    if (years <= 2) {
        dog_years += years * 10.5;
    } else {
        dog_years += 2 * 10.5;
        dog_years += (years - 2) * 4;
    }
    return dog_years;
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

- (IBAction)calculateRealAge:(UIButton *)sender {
    float result = [self.input.text floatValue];
    self.output.text = [
        NSString stringWithFormat:@"%.2f years",
        calculate_real_age(result)
    ];
    // hide keyboard
    [self.input resignFirstResponder];
}

@end
