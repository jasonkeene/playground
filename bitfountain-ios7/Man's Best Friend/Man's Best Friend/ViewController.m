//
//  ViewController.m
//  Man's Best Friend
//
//  Created by bonobo on 7/11/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"
#import "Dog.h"


@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.

    // dog stuff
    NSLog(@"\n"
           "\n"
           "dog stuff\n"
           "=========");
    self.dogs = [@[
        [[Dog alloc] initWithName:@"Nana"
                            breed:@"St. Bernard"
                              age:1
                            image:[UIImage imageNamed:@"St.Bernard.JPG"]],
        [[Dog alloc] initWithName:@"Wishbone"
                            breed:@"Jack Russell Terrier"
                              age:1
                            image:[UIImage imageNamed:@"JRT.jpg"]],
        [[Dog alloc] initWithName:@"Lassie"
                            breed:@"Collie"
                              age:1
                            image:[UIImage imageNamed:@"BorderCollie.jpg"]],
        [[Dog alloc] initWithName:@"Angel"
                            breed:@"Greyhound"
                              age:1
                            image:[UIImage imageNamed:@"ItalianGreyhound.jpg"]],
    ] mutableCopy];
    [self displayDog:self.dogs[3]];

    // Challenge 4: Methods - Problem 1
    NSLog(@"\n"
           "\n"
           "Challenge 4: Methods - Problem 1\n"
           "================================");
    [self printToOneFrom:7];

    // Challenge 4: Methods - Problem 2
    NSLog(@"\n"
           "\n"
           "Challenge 4: Methods - Problem 2\n"
           "================================");
    [self printTo:20 from:11];

    // Challenge 4: Methods - Problem 3
    NSLog(@"\n"
           "\n"
           "Challenge 4: Methods - Problem 3\n"
           "================================");
    NSLog(@"15! == %lli", [self factorial:15]);
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (void)printHelloWorld {
    NSLog(@"Hello world!");
}

- (void)printTo:(int)x from:(int)y {
    if (x > y) {
        [self swap:&x with:&y];
    }
    for (; y > x - 1; y--) {
        NSLog(@"%i", y);
    }
}

- (void)swap:(int*)x with:(int*)y {
    int tmp = *x;
    *x = *y;
    *y = tmp;
}

- (void)printToOneFrom:(int)x {
    [self printTo:1 from:x];
}

- (long long)factorial:(int)x {
    long long result = 1;
    for (; x > 1; x--) {
        result *= x;
    }
    return result;
}

- (void)displayDog:(Dog*)dog {
    self.imageView.image = dog.image;
    self.nameLabel.text = dog.name;
    self.breedLabel.text = dog.breed;
}

@end
