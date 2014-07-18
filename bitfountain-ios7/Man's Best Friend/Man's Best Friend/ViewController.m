//
//  ViewController.m
//  Man's Best Friend
//
//  Created by bonobo on 7/11/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"
#import "Dog.h"
#import "Puppy.h"


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
    NSLog(@"%@", self.dogs);
    [self setRandomDog];

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

    // puppy stuff
    NSLog(@"\n"
          "\n"
          "puppy stuff\n"
          "===========");
    Puppy* puppy = [[Puppy alloc] initWithName:@"Bo"
                                         breed:@"Portuguese Water Dog"
                                           age:1
                                         image:[UIImage imageNamed:@"Bo.jpg"]];
    [self displayDog:puppy];
    [puppy givePuppyEyes];
    [puppy bark];
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

- (IBAction)newDogButton:(UIBarButtonItem *)sender {
    [self setRandomDogWithCallback:^ {
        sender.title = @"And Another";
    }];
}

- (void)setRandomDog {
    [self displayDog:[self getRandomDog]];
}

- (void)setRandomDogWithCallback:(void(^)())complete {
    [UIView transitionWithView:self.view duration:1.5 options:UIViewAnimationOptionTransitionCrossDissolve animations:^{
        [self displayDog:[self getRandomDog]];
    } completion:^(BOOL finished) {
        complete();
    }];
}

- (Dog*)getRandomDog {
    Dog* randomDog = 0;
    while (randomDog == 0 || randomDog == self.currentDog) {
        randomDog = self.dogs[arc4random() % [self.dogs count]];
    }
    return randomDog;
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
    self.currentDog = dog;
}

@end
