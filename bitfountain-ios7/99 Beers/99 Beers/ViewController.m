//
//  ViewController.m
//  99 Beers
//
//  Created by bonobo on 7/11/14.
//  Copyright (c) 2014 bonobo. All rights reserved.
//

#import "ViewController.h"

@interface ViewController ()
            

@end

@implementation ViewController
            
- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
    NSString *song_fmt = @"%d bottles of beer on the wall, %d bottles of beer. You take one down, pass it around %d bottles of beer on the wall.";
    for (int i = 99; i > 0; i--) {
        NSLog(song_fmt, i, i, i - 1);
    }
    NSLog(@ "No bottles of beer on the wall, no bottles of beer.  I think you had quite enough at this point!");
}

- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}

@end
