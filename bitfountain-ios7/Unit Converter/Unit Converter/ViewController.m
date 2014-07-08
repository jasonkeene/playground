//
//  ViewController.m
//  Unit Converter
//
//  Created by bonobo on 7/7/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"


// convert bills to football fields
float convertIt(float bills) {
    return bills / 91440;
}


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
    self.numberOfBillsLabel.text = [
        NSString stringWithFormat:@"%f",
        convertIt([self.numberOfBillsTextField.text floatValue])
    ];
    // hide keyboard
    [self.numberOfBillsTextField resignFirstResponder];
}

@end