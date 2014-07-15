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

@end
