//
//  ViewController.m
//  Class Testing Ground
//
//  Created by bonobo on 7/18/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"
#import "Dog.h"

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    
    NSString* sentence = @"The NewFoundland dog breed has webbed feet which aids in its swimming prowess";
    NSArray* words = [sentence componentsSeparatedByString:@" "];
    NSMutableArray* cappedWords = [[NSMutableArray alloc] init];
//    for (int i = 0; i < [words count]; i++) {
//        cappedWords[i] = [words[i] capitalizedString];
//    }
    for (NSString* word in words) {
        [cappedWords addObject:[word capitalizedString]];
    }
//    NSLog(@"%@", cappedWords);
    Dog* dog = [[Dog alloc] init];
    [dog setName:@"test"];
    NSLog(@"%@",[dog name]);
    dog.name = @"test2";
    NSLog(@"%@", dog.name);
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
