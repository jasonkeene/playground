//
//  Armor.m
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "Armor.h"


@implementation Armor

- (instancetype)init {
    return [self initWithName:@"Shirt" value:5];
}

- (instancetype)initWithName:(NSString*)name value:(int)value {
    self = [super init];
    self.name = name;
    self.value = value;
    return self;
}

- (NSString*)displayString {
    if (self.value > 0) {
        return [NSString stringWithFormat:@"%@ (+%d)", self.name,
                self.value, nil];
    } else {
        return [NSString stringWithFormat:@"%@", self.name, nil];
    }
}

@end
