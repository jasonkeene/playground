//
//  ViewController.h
//  Man's Best Friend
//
//  Created by bonobo on 7/11/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "Dog.h"

@interface ViewController : UIViewController

@property (strong, nonatomic) IBOutlet UIImageView *imageView;
@property (strong, nonatomic) IBOutlet UILabel *nameLabel;
@property (strong, nonatomic) IBOutlet UILabel *breedLabel;
@property (strong, nonatomic) NSMutableArray* dogs;
@property (strong, nonatomic) Dog* currentDog;

- (IBAction)newDogButton:(UIBarButtonItem *)sender;

- (void)printHelloWorld;
- (void)printTo:(int)x from:(int)y;
- (void)printToOneFrom:(int)x;
- (long long)factorial:(int)x;
- (void)displayDog:(Dog*)dog;

@end
