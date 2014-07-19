//
//  Dog.m
//  Class Testing Ground
//
//  Created by bonobo on 7/18/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Dog.h"

@implementation Dog

- (void)setName:(NSString *)name {
    NSLog(@"setter called");
    _name = name;
}

- (NSString*)name {
    NSLog(@"getter called");
    return _name;
}

@end
