//
//  ViewController.h
//  Prep For Pirate Assignment
//
//  Created by bonobo on 7/19/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "AwesomeClass.h"

@interface ViewController : UIViewController

@property (nonatomic) CGPoint currentPoint;
@property (strong, nonatomic) IBOutlet UIButton *button;
@property (strong, nonatomic) AwesomeClass* awesome;

@end

