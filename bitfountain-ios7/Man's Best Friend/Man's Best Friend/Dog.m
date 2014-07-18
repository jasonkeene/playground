//
//  Dog.m
//  Man's Best Friend
//
//  Created by bonobo on 7/12/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Dog.h"


@implementation Dog


- (instancetype)initWithName:(NSString*)name breed:(NSString*)breed age:(int)age image:(UIImage*)image {
    self = [super init];
    self.name = name;
    self.breed = breed;
    self.age = age;
    self.image = image;
    return self;
}

- (instancetype)init {
    return [self initWithName:@"Rando dog!" breed:@"Mutt" age:5 image:nil];
}

- (void)bark {
    NSLog(@"Woof Woof!");
}

- (void)barkTimes:(int)times {
    for (int i = 0; i < times; i++) {
        [self bark];
    }
}

- (void)barkTimes:(int)times loudly:(BOOL)loud {
    if (loud) {
        for (int i = 0; i < times; i++) {
            NSLog(@"Ruff Ruff!");
        }
    } else {
        for (int i = 0; i < times; i++) {
            NSLog(@"yip yip!");
        }
    }
}

- (void)lycanize {
    self.breed = @"Lycan";
}

- (int)ageInDogYearsFromAge:(int)years {
    return years * 7;
}

- (NSString*)description {
    return [NSString stringWithFormat:@"<Dog: %@>", self.name];
}

@end
