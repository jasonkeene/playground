//
//  Dog.h
//  Man's Best Friend
//
//  Created by bonobo on 7/12/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <UIKit/UIKit.h>


@interface Dog : NSObject

@property (nonatomic, strong) NSString* name;
@property (nonatomic, strong) NSString* breed;
@property (nonatomic) int age;
@property (nonatomic, strong) UIImage* image;

- (instancetype)initWithName:(NSString*)name breed:(NSString*)breed age:(int)age image:(UIImage*)image;
- (void)bark;
- (void)barkTimes:(int)times;
- (void)barkTimes:(int)times loudly:(BOOL)loud;
- (void)lycanize;
- (int)ageInDogYearsFromAge:(int)years;

@end
