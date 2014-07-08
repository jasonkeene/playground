//
//  ViewController.h
//  Age of Laika
//
//  Created by bonobo on 7/7/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>


@interface ViewController : UIViewController

@property (strong, nonatomic) IBOutlet UITextField *input;

@property (strong, nonatomic) IBOutlet UILabel *output;

- (IBAction)calculateAge:(UIButton *)sender;

@end