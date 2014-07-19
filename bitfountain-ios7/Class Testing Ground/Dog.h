//
//  Dog.h
//  Class Testing Ground
//
//  Created by bonobo on 7/18/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface Dog : NSObject {
    NSString* _name;
}

- (void)setName:(NSString *)name;
- (NSString*)name;

@end
