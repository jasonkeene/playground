//
//  Armor.h
//  Pirate Adventure
//
//  Created by bonobo on 7/20/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <Foundation/Foundation.h>


@interface Armor : NSObject

@property (strong, nonatomic) NSString* name;
@property (nonatomic) int value;

- (instancetype)initWithName:(NSString*)name value:(int)value;
- (NSString*)displayString;

@end
