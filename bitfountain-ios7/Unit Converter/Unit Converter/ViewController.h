//
//  ViewController.h
//  Unit Converter
//
//  Created by bonobo on 7/7/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>

@interface ViewController : UIViewController

@property (strong, nonatomic) IBOutlet UILabel *numberOfBillsLabel;

@property (strong, nonatomic) IBOutlet UITextField *numberOfBillsTextField;

- (IBAction)convertUnits:(UIButton *)sender;

@end

